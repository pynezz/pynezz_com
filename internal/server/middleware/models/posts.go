package models

import (
	"time"

	"gorm.io/gorm"
)

// Post is a struct that represents a post.
type Post struct {
	gorm.Model

	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Metadata Metadata `json:"metadata" gorm:"embedded, embeddedPrefix:metadata_"`
	Content  string   `json:"content"`
	Path     string   `json:"path"` // Path to the markdown file - e.g. "/posts/2021-01-01-post.md"
}

type Metadata struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Date         time.Time
	LastModified time.Time
	Tags         []string
}

type Posts []Post

// This should make it the easiest to search and sort the posts
type PostsMetadata struct {
	Path         string // Location of the post, either relative to the executable or an absolute path
	Title        string
	PostID       int `json:"post_id" gorm:"primaryKey"` // Todo: Need to add format 'shorttitle-DDMMYY' to the post ID
	LastModified time.Time
}
