package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
)

// CustomContext extends echo.Context with user information
type CustomContext struct {
	echo.Context
	User models.User
}

// Bouncer is a middleware that checks if the user is authenticated
func Bouncer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Checking if user is authenticated")

		cookie, err := c.Request().Cookie("Authorization")
		if err != nil {
			fmt.Println("Error getting cookie: ", err)
			return echo.ErrCookieNotFound
		}

		token, err := VerifyJWTToken(cookie.Value)
		if err != nil {
			return echo.ErrInternalServerError
		}

		fmt.Println("Token send by client: ", token)
		// Check if the token is valid
		if !token.Valid {
			return echo.ErrUnauthorized
		}

		username, err := token.Claims.GetSubject()
		if err != nil {
			return c.Redirect(http.StatusFound, "/login")
		}

		user := models.User{Username: username}
		if res, httpCode := getUser(token.Raw, username); res != "" && (httpCode == http.StatusOK || httpCode == http.StatusAccepted) {
			fmt.Println("User is authenticated")
			err := json.Unmarshal([]byte(res), &user)
			if err != nil {
				return echo.ErrInternalServerError
			}
			cc := &CustomContext{c, user}
			return next(cc)
		}

		// if res, httpCode := getUser(token.Raw, username); res != "" && (httpCode == http.StatusOK || httpCode == http.StatusAccepted) {
		// 	// User is authenticated
		// 	fmt.Println("User is authenticated")

		// 	return next(c), models.User{
		// 		Username: username,
		// 		UserID:   Uuid(username).AsUint(),
		// 	}
		// }

		return next(c)
	}
}

func Register(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Got a POST request to /register")

		fmt.Println("Registering user")

		if c.FormValue("username") == "" || c.FormValue("password") == "" {
			return echo.ErrBadRequest
		}

		fmt.Println("Username: ", c.FormValue("username"))
		fmt.Println("Password: ", c.FormValue("password"))

		// Argon2 hash the password
		argon2 := NewArgon2().InitArgon(c.FormValue("password"))
		encodedHash := argon2.GetEncodedHash()

		// Check if the user already exists
		if userExists(&models.User{Username: c.FormValue("username")}) {
			return c.JSON(http.StatusConflict, echo.Map{
				"message:": "Failed to register user",
			})
		}

		username := c.FormValue("username")
		userId := Uuid(c.FormValue(username)).AsUint()
		fmt.Println("User ID: ", userId)

		// Register the user
		newUser := &models.User{
			Username: username,
			Password: encodedHash,
			UserID:   Uuid(username).AsUint(),
		}

		// Save the user to the database
		if err := writeUser(newUser); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to register user",
			})
		}

		// Generate a JWT token
		token := GenerateJWTToken(*newUser)
		c.Response().Header().Set("Authorization: Bearer ", token)
		c.SetCookie(&http.Cookie{
			Name:     "Authorization",
			Value:    token,
			HttpOnly: true, // OWASP: https://owasp.org/www-community/HttpOnly
		})

		return next(c)
	}
}
