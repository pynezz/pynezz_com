package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Metadata struct {
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Date         time.Time      `json:"date"`
	LastModified time.Time      `json:"last_modified"`
	Tags         datatypes.JSON `json:"tags" gorm:"type:json"`
}

// Post is a struct that represents a post.
type Post struct {
	gorm.Model
	Title    string   `json:"title"`
	Metadata Metadata `json:"metadata" gorm:"embedded;embeddedPrefix:metadata_"`
	Content  string   `json:"content"`
	Path     string   `json:"path"` // Path to the markdown file - e.g. "/posts/2021-01-01-post.md"
}

// This should make it the easiest to search and sort the posts
type PostMetadata struct {
	gorm.Model
	Path         string         `json:"path"` // Location of the post, either relative to the executable or an absolute path
	Title        string         `json:"title"`
	Slug         string         `json:"slug"`    // A short title for the post, used in the URL
	PostID       int            `json:"post_id"` // Todo: Need to add format 'slug (shorturl_ddmmyy)' to the post ID
	LastModified time.Time      `json:"last_modified"`
	Tags         datatypes.JSON `json:"tags" gorm:"type:json"`
	Summary      string         `json:"summary"` // A short summary of the post
}

type Posts []Post

/*
PostsMetadata
*/
