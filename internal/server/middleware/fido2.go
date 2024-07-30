package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	"github.com/pynezz/pynezzentials/ansi"
)

var (
	webAuthn *webauthn.WebAuthn
	// err      error

	datastore PasskeyStore
	sessions  = make(map[string]models.Session)
)

// UserWrap is a wrapper for user data, used for registration, making sure parameters are correct
type UserWrap struct {
	Username    string
	Displayname string
	Credentials webauthn.Credential
}

func newAdmin(u UserWrap) (models.Admin, error) {
	if u.Username == "" || u.Displayname == "" {
		return models.Admin{}, fmt.Errorf("missing required fields")
	}

	return models.Admin{
		Name:          u.Username,
		DisplayName:   u.Displayname,
		Authenticated: false,
		AdminID:       Uuid(u.Username).AsUint(),
		Credentials:   models.JSONSlice{u.Credentials},
	}, nil
}

type PasskeyUser interface {
	webauthn.User
	AddCredential(*webauthn.Credential)
	UpdateCredential(*webauthn.Credential)
}

type PKParams struct {
	Host    string
	Port    string
	Proto   string
	AppName string
}

type PasskeyStore interface {
	GetUser(userName string) PasskeyUser
	SaveUser(PasskeyUser)
	GetSession(token string) webauthn.SessionData
	SaveSession(token string, data webauthn.SessionData)
	DeleteSession(token string)
}

func DefaultWAuth() *webauthn.WebAuthn {
	if webAuthn != nil {
		return webAuthn
	}

	h := getEnv("HOST")
	p := getEnv("PORT")
	proto := getEnv("PROTO") // should be https, but it's http during development
	appName := getEnv("APP_NAME")
	switch {
	case h == "":
		ansi.PrintError("HOST (your server) is not set.\nPlease set it in the .env file.\nExample:\n\tHOST=example.com")
	case p == "":
		ansi.PrintError("PORT is not set.\nPlease set it in the .env file.\nExample:\n\tPORT=8080")
	case proto == "":
		ansi.PrintError("PROTO (protocol) is not set.\nPlease set it in the .env file.\nExample:\n\tPROTO=https")
	case appName == "":
		ansi.PrintError("APP_NAME is not set.\nPlease set it in the .env file.\nExample:\n\tAPP_NAME=AppName")
	default: // all good
		ansi.PrintSuccess(".env file is set up correctly: " + proto + "://" + h + ":" + p)
	}

	rpid := h
	origins := []string{fmt.Sprintf("%s://%s:%s", proto, h, p)}
	return newWebAuthn("Pynezz", rpid, origins...)
}

func newWebAuthn(displayName, rpid string, origins ...string) *webauthn.WebAuthn {
	var wa *webauthn.WebAuthn
	var err error

	conf := &webauthn.Config{
		RPDisplayName: displayName,
		RPID:          rpid, // should map this to a port, and configure it in openresty
		RPOrigins:     origins,
	}

	if wa, err = webauthn.New(conf); err != nil {
		ansi.PrintError(err.Error())
	}

	return wa
}

func finishFido2Login(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	username := c.FormValue("username")
	t := c.Request().Header.Get("Session-Key")

	gottenKey := datastore.GetSession(t)
	ansi.PrintInfo("Got session key: " + string(gottenKey.UserID))

	a, err := getAdminByUsername(username)
	if err != nil {
		return echo.ErrInternalServerError
	}

	session := sessions[string(gottenKey.UserID)].SessionData

	credential, err := DefaultWAuth().FinishLogin(a, webauthn.SessionData(session), c.Echo().AcquireContext().Request())
	if err != nil {
		return echo.ErrInternalServerError
	}

	a.AddCredential(credential)

	if err = writeAdminToDatabae(&a); err != nil {
		return echo.ErrInternalServerError
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Header().Set("Session-Key", string(session.UserVerification))

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"status":  "ok",
	})
}

func HandleFirstTimeAdminRegister(c echo.Context) error {

	r := json.NewDecoder(c.Request().Body)
	ansi.PrintDebug("Got a POST request to /passkey/registerStart: " + fmt.Sprintf("%+v", r))

	// if adminIsInitialized() {
	// 	return c.JSON(http.StatusBadRequest, echo.Map{
	// 		"message":    "No funny business!",
	// 		"status":     "error",
	// 		"statusText": "Admin already initialized",
	// 	})
	// }
	req := c.Request()
	if req.Method != http.MethodPost {
		return c.JSON(http.StatusMethodNotAllowed, echo.Map{
			"message": "Method not allowed",
			"status":  "error",
		})
	}

	body := req.Body
	defer body.Close()

	var bodyContent []byte
	_, err := body.Read(bodyContent)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message":    "Missing form parameter: " + err.Error(),
			"status":     "error",
			"statusText": "Unable to process form data",
		})
	}

	username, displayname, err := getUsername(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message":    "Missing form parameter: " + err.Error(),
			"status":     "error",
			"statusText": "Unable to process form data",
		})
	}

	if username == "" || displayname == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message":    "Missing form parameter",
			"status":     "error",
			"statusText": "Username and display name are required",
		})
	}

	ansi.PrintBold("Got a POST request to /passkey/registerStart. Payload:\n" + "{" + username + ", " + displayname + "}")

	a, err := newAdmin(UserWrap{
		Username:    username,
		Displayname: displayname,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message":    "Missing form parameter: " + err.Error(),
			"status":     "error",
			"statusText": "Unable to process form data",
		})
	}

	options, sessionData, err := DefaultWAuth().BeginRegistration(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message":    "Failed to begin registration",
			"status":     "error",
			"statusText": "Unable to process form data",
		})
	}

	if err = writeAdminToDatabae(&a); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message":    "Failed to write admin to database",
			"status":     "error",
			"statusText": "Unable to process form data",
		})
	}

	sess := models.Session{
		SessionID:   string(sessionData.UserID),
		SessionData: models.JSONSessionData(*sessionData),
	}

	// Make a session key and store the sessionData values
	t := uuid.New().String()
	datastore.SaveSession(t, *sessionData)

	// TODO: switch out with proper handling in the future
	sessions[t] = sess
						
	writeWASessionData(&sess)

	// return c.JSON(http.StatusOK, echo.Map{
	// 	"options":     options,
	// 	"sessionData": sessionData,
	// })'
	return JSONResponse(c.Response().Writer, t, options, http.StatusOK)
}

func finishFido2Registration(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	// t := c.Request().Header.Get("Session-Key")

	username := c.Request().FormValue("username")

	a, err := getAdminByUsername(username)
	if err != nil {
		return echo.ErrInternalServerError
	}

	// TODO: This will fail. The user id is used as the session key to fetch the session data, not the username
	// Kinda vulnerable due to a fake request with a valid username? I know FIDO2 is supposed to be secure,
	session := sessions[string(a.WebAuthnID())].SessionData

	ansi.PrintSuccess("Set session data: " + string(session.UserID))

	credential, err := DefaultWAuth().FinishRegistration(a, webauthn.SessionData(session), c.Echo().AcquireContext().Request())
	if err != nil {
		return echo.ErrInternalServerError
	}

	a.Credentials = append(a.Credentials, *credential)

	if err = writeAdminToDatabae(&a); err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Registration successful",
		"status":  "ok",
	})
}

func generateRegistrationOptions(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	username := c.FormValue("username")

	a, err := getAdminByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get admin by username",
			"status":  "error",
		})
	}

	options, sessionData, err := DefaultWAuth().BeginRegistration(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to begin registration",
			"status":  "error",
		})
	}

	t := Uuid(username).Identifier

	datastore.SaveSession(t, *sessionData)

	return c.JSON(http.StatusOK, echo.Map{
		"message": sessionData.UserID,
		"status":  "ok",
		"options": options,
	})
}
