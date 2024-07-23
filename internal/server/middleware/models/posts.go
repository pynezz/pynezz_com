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
	Title    string         `json:"title"`
	Metadata Metadata       `json:"metadata" gorm:"embedded;embeddedPrefix:metadata_"`
	Content  datatypes.JSON `json:"content"`
	Path     string         `json:"path"` // Path to the markdown file - e.g. "/posts/2021-01-01-post.md"
	Slug     string         `json:"slug"`
}

// This should make it the easiest to search and sort the posts
// It's different from the other two due to it being a "slim" version of the Post struct.
// It's used to display a list of posts, and not the full content of a post.
// Will need to reconsider this in the future (EOM July 2024) though. It just doesn't quite feel right.
// TODO: Consider the use of this struct and how it fits together with the overall structure.
type PostMetadata struct {
	gorm.Model
	Path         string         `json:"path"` // Location of the post, either relative to the executable or an absolute path
	Title        string         `json:"title"`
	Slug         string         `json:"slug"` // A short title for the post, used in the URL
	PostID       int            `json:"post_id"`
	LastModified time.Time      `json:"last_modified"`
	Tags         datatypes.JSON `json:"tags" gorm:"type:json"`
	Summary      string         `json:"summary"` // A short summary of the post
}

type Tag struct {
	gorm.Model
	Tags datatypes.JSON `json:"tags" gorm:"type:json"`
}

type Posts []Post
