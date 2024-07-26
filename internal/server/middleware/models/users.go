package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
	"gorm.io/gorm"
)

// User defines the structure of the user model
type User struct {
	gorm.Model // Embed the `Model` struct, which contains fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`

	UserID   uint   `gorm:"primaryKey; unique;"`
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`

	Role string `gorm:"not null"`
}

// Separate the user and admin models
type Admin struct {
	gorm.Model

	Name          string
	DisplayName   string
	Authenticated bool
	AdminID       uint `gorm:"primaryKey;unique;"`
	// 	Credentials   []webauthn.Credential `gorm:"type:json;not null"`
	Credentials JSONSlice `gorm:"type:json;not null"` // Custom type due to gorm not supporting webauthn.Credential
}

type Session struct {
	gorm.Model

	// SessionData webauthn.SessionData
	SessionData JSONSessionData `gorm:"type:json;not null"`
	SessionID   string          `gorm:"primaryKey;unique;"`
}

// Custom type - should make gorm happy
type JSONSlice []webauthn.Credential
type JSONSessionData webauthn.SessionData

/**
These methods are used to implement the driver.Valuer and sql.Scanner interfaces for the JSONSlice and JSONSessionData types.
*/

func (j JSONSlice) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONSlice) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), j)
}

func (j JSONSessionData) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONSessionData) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), j)
}

/**
These methods are used to implement the webauthn.User interface for the Admin type.
*/

// WebAuthnCredentials implements webauthn.User.
func (a Admin) WebAuthnCredentials() []webauthn.Credential {
	return a.Credentials
}

// WebAuthnDisplayName implements webauthn.User.
func (a Admin) WebAuthnDisplayName() string {
	return a.DisplayName
}

// WebAuthnID implements webauthn.User.
func (a Admin) WebAuthnID() []byte {
	return []byte(fmt.Sprintf("%d", a.AdminID))
}

// WebAuthnIcon implements webauthn.User.
func (a Admin) WebAuthnIcon() string {
	return ""
}

// WebAuthnName implements webauthn.User.
func (a Admin) WebAuthnName() string {
	return a.Name
}

func (a *Admin) SetCredentials(creds []webauthn.Credential) {
	a.Credentials = creds
}

func (a *Admin) AddCredential(cred *webauthn.Credential) {
	a.Credentials = append(a.Credentials, *cred)
}
