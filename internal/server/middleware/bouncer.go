package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func Sec(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ua := c.Request().UserAgent()
		// if vips[ua] {
		// ansi.PrintBold(ansi.FormatRoundedBox("VIP Detected\n"))
		// return next(c)
		// }
		ansi.PrintBold("user agent: " + ua)
		if h := NewArgon2().InitArgonWithSalt(ua, "mykKhiHGSuBxmY5NAi"); strings.Split(h.encodedHash, "$")[5] == "CDHqziZMzGB0D8S4NOfFV3TbRMY3WruoNhI6QEOwnLc" {
			ansi.PrintBold(ansi.FormatRoundedBox("New VIP Detected\n"))
			// vips[ua] = true
		} else {
			ansi.PrintWarning("not VIP. hash: " + h.encodedHash) //
		}

		return next(c)
	}
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

// GenerateNonce generates a random nonce
func GenerateNonce() (string, error) {
	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(nonce), nil
}

// getFileHash returns the hash of a file for cache busting
func getFileHash(file string) string {
	return ""
}

// CommonHeaders sets common headers for all requests
func CommonHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// set cache control headers and common headers (cache static files for 1 week)
		c.Response().Header().Set("Cache-Control", "public, max-age=604800")

		// pragma and expires are deprecated
		c.Response().Header().Set("Pragma", "cache")

		return next(c)
	}
}

func SecurityHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// set security headers for proxying through openresty
		c.Response().Header().Set("X-Frame-Options", "DENY")
		c.Response().Header().Set("X-Content-Type-Options", "nosniff")
		c.Response().Header().Set("X-XSS-Protection", "1; mode=block")
		return next(c)
	}
}

func FIDO2Auth(next echo.HandlerFunc) echo.HandlerFunc {
	// handleFido2Register( )

	// wAuth.BeginLogin()
	return func(c echo.Context) error {
		return next(c)
	}
}

// Fido and Brutus are both dog names. Brutus is the guard dog.
// Brutus is a middleware that checks if the user is FIDO authenticated.
func Brutus(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := c.Request().Header.Get("Session-ID")
		session, exists := sessions[sessionID]
		if !exists {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		uid := uint(binary.BigEndian.Uint64(session.SessionData.UserID))
		user, err := getAdminByID(uid)
		if err != nil || !user.Authenticated {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}

func handleLogout(c echo.Context) error {
	cookie, _ := c.Cookie("Authorization")
	if cookie != nil {
		cookie.MaxAge = -1
		c.SetCookie(cookie)
	}

	return c.Redirect(http.StatusMovedPermanently, "/login")
}

func HandleFido2LoginFinish(c echo.Context) error {
	return finishFido2Login(c)
}

func HandleFido2Login(c echo.Context) error {

	body := c.Request().Body
	defer body.Close()

	username := c.FormValue("username")

	a, err := getAdminByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to login",
			"status":  "error",
		})
	}

	options, sessionData, err := DefaultWAuth().BeginLogin(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to login",
			"status":  "error",
		})
	}

	t := Uuid(username).Identifier

	datastore.SaveSession(t, *sessionData)

	fmt.Printf("Got a POST request to /fido2/register: %+v\n", body)

	return c.JSON(http.StatusOK, echo.Map{
		"message": sessionData.UserID,
		"status":  "ok",
		"options": options,
	})
}

func HandleFido2RegisterFinish(c echo.Context) error {
	return finishFido2Registration(c)
}

func getUsername(r *http.Request) (username string, displayname string, err error) {
	type user struct {
		Username    string `json:"username"`
		DisplayName string `json:"displayname"`
	}
	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return "", "", err
	}

	return u.Username, u.DisplayName, nil
}

func JSONResponse(w http.ResponseWriter, sessionKey string, data interface{}, status int) error {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Session-Key", sessionKey)
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)

	return err
}

func GenerateRegistrationOptions(c echo.Context) error {
	return generateRegistrationOptions(c)
}
