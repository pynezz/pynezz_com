package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
)

// Bouncer is a middleware that checks if the user is authenticated
func Bouncer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Checking if user is authenticated")

		cookie, err := c.Request().Cookie("Authorization")
		if err != nil {
			return err
		}

		token, err := VerifyJWTToken(cookie.Value)
		if err != nil {
			return err
		}

		fmt.Println("Token send by client: ", token)
		// Check if the token is valid
		if !token.Valid {
			return echo.ErrUnauthorized
		}

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

		// Check if the user already exists
		if userExists(&models.User{Username: c.FormValue("username")}) {
			return c.JSON(http.StatusConflict, echo.Map{
				"message:": "Failed to register user",
			})
		}

		// Register the user
		newUser := &models.User{
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
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

		return c.JSON(http.StatusOK, echo.Map{
			"message": "User registered",
			"token:	": token,
		})
	}
}
