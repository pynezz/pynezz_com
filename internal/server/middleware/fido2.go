package middleware

import (
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
)

// func initWebAuthn() {
// 	var webAuthn *webauthn.WebAuthn
// 	var err error

// }

func newWebAuthn(rpid string, origins ...string) *webauthn.WebAuthn {
	var webAuthn *webauthn.WebAuthn
	var err error

	conf := &webauthn.Config{
		RPDisplayName: "Pynezz",
		RPID:          rpid, // should map this to a port, and configure it in openresty
		RPOrigins:     origins,
	}

	if webAuthn, err = webauthn.New(conf); err != nil {
		fmt.Println(err)
	}

	return webAuthn
}

func BeginRegistration() {

}
