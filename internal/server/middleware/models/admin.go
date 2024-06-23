package models

// Separate the user and admin models
type Admin struct {
	User `gorm:"embedded"`

	Role string `gorm:"not null" json:"role"`
}
