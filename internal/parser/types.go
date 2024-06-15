package parser

import (
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
}
