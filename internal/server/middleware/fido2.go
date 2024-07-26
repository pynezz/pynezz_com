package middleware

import (
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	"github.com/pynezz/pynezzentials/ansi"
)

var (
	webAuthn *webauthn.WebAuthn
	// err      error

	datastore PasskeyStore
	sessions  = make(map[string]models.Session)

	// pkparams = &PKParams{
	// 	Host:    getEnv("HOST"),
	// 	Port:    getEnv("PORT"),
	// 	Proto:   getEnv("PROTO"),
	// 	AppName: getEnv("APP_NAME"),
	// }
)

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
