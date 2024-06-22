package parser

import (
	"crypto/sha256"
	"errors"
	"time"
)

var (
	ErrorNoClosingDelimiter = errors.New("error: error reading metadata, no closing delimiter found.\nFormat:\n\t---\n\tmetadata\n\t---\n\tcontent\n")
)

type Metadata struct {
	Title        string
	Description  string
	Date         time.Time
	LastModified time.Time
	Tags         []string
}

type Post struct {
	Metadata Metadata
	Content  []byte

	sha256sum string

	CalculateSha256Sum func() string
}

type PostsPage struct {
	*Page    // Embedding the Page struct to inherit its fields
	Metadata Metadata
	Posts    []Post

	sha256sum string
}

type Page struct {
	Metadata Metadata

	sha256sum string
}

type Site struct {
	Pages []Page

	sha256sum string
}

func NewPost(content string) *Post {
	p := &Post{
		Content: []byte(content),
	}

	return &Post{
		CalculateSha256Sum: func() string {
			if p.sha256sum == "" && p.Content != nil {
				hash := sha256.New()
				hash.Write(p.Content)
				p.sha256sum = string(hash.Sum(nil))
			}
			return p.sha256sum
		},
	}
}

func NewPostsPage() *PostsPage {
	return &PostsPage{}
}

func NewPage() *Page {
	return &Page{}
}

func NewSite() *Site {
	return &Site{}
}
