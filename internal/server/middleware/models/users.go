package models

import "gorm.io/gorm"

// User defines the structure of the user model
type User struct {
	gorm.Model // Embed the `Model` struct, which contains fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`

	Username string `gorm:"unique"`
	Password string `gorm:"not null"`

	Session string `gorm:"not null" json:"session"`
}
