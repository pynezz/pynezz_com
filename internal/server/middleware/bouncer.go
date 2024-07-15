package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	"github.com/pynezz/pynezz_com/templates"
	"github.com/pynezz/pynezzentials/ansi"
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
			return render(c, http.StatusInternalServerError, templates.Login())
		}

		token, err := VerifyJWTToken(cookie.Value)
		if err != nil {
			return render(c, http.StatusUnauthorized, templates.Login())
		}

		fmt.Println("Token send by client: ", token)
		// Check if the token is valid
		if !token.Valid {
			return echo.ErrUnauthorized
		}

		username, err := token.Claims.GetSubject()
		if err != nil {
			return render(c, http.StatusInternalServerError, templates.Login())
			// return c.Redirect(http.StatusFound, "/login")
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

		return next(c)
	}
}

func Login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Got a POST request to /login")

		fmt.Println("Logging in user")

		if c.FormValue("username") == "" || c.FormValue("password") == "" {
			return echo.ErrBadRequest
		}

		fmt.Println("Username: ", c.FormValue("username"))
		// fmt.Println("Password: ", c.FormValue("password"))
		password := c.FormValue("password")

		// Check if the user exists
		if !userExists(&models.User{Username: c.FormValue("username")}) {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message:": "Failed to login user",
			})
		}

		// Get the user from the database
		user := &models.User{Username: c.FormValue("username")}
		if res, httpCode := getUserHash(user.Username); res != "" && (httpCode == http.StatusOK || httpCode == http.StatusAccepted) {
			params, salt, hash, err := DecodeHash(res)
			encodedhash := HashToEncodedHash(params, hash, salt)
			encodedSuppliedPassword := NewArgon2().InitArgonWithSalt(password, string(salt))
			if encodedhash != encodedSuppliedPassword.GetEncodedHash() {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Failed to login user",
				})
			} else {
				ansi.PrintSuccess("Password and hash match")
			}

			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Failed to login user",
				})
			}
			// if ok, err := HashesMatch(password, salt); err != nil || !ok {
			// 	return c.JSON(http.StatusUnauthorized, echo.Map{
			// 		"message": "Failed to login user",
			// 	})
			// }
		} else {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to login user",
			})
		}

		// Generate a JWT token
		token := GenerateJWTToken(*user)
		c.Response().Header().Set("Authorization: Bearer ", token)
		c.SetCookie(&http.Cookie{
			Name:     "Authorization",
			Value:    token,
			HttpOnly: true, // OWASP: https://owasp.org/www-community/HttpOnly
			SameSite: http.SameSiteDefaultMode,
		})

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
			SameSite: http.SameSiteStrictMode,
		})

		return next(c)
	}
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func SecurityHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Frame-Options", "DENY")
		c.Response().Header().Set("X-Content-Type-Options", "nosniff")
		c.Response().Header().Set("X-XSS-Protection", "1; mode=block")

		// this will need to be implemented, but it breaks the site as of now - need to fix it
		// c.Response().Header().Set("Content-Security-Policy", "default-src 'self'")
		return next(c)
	}
}	
