package models

import (
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
	User // Embed the `User` struct, which contains fields `UserID`, `Username`, `Password`, `Role`

	AdminID uint `gorm:"primaryKey; unique;"`
	webauthn.Credential
}
