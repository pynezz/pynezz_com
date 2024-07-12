package parser

import (
	"errors"
	"time"
)

var (
	ErrorNoClosingDelimiter = errors.New("error: error reading metadata, no closing delimiter found.\nFormat:\n\t---\n\tmetadata\n\t---\n\tcontent\n")
)

// Good resource for rules when writing markdown
// (and the standard adhered to in this project)
// https://xiangxing98.github.io/Markdownlint_Rules.html#md001
// (raw: https://raw.githubusercontent.com/DavidAnson/markdownlint/v0.24.0/doc/Rules.md)

type h1 string
type h2 string
type h3 string
type h4 string
type h5 string
type h6 string
type p string
type a string
type img string
type ul string
type ol string
type li string

// `code`
type code string

// ```language
// codeblock
// ```
type codeblock string

// preformatted text
// ```
// preformatted text
// ```
type pre string

// > blockquote
type blockquote string

// ---
type hr string

// | table | header |
type br string

// | table | row |
type table string

// | table | header |
type tr string

type th string
type td string

// **strong** - for important text
type strong string

// **bold** - for bold text, use sparingly
type bold string
type b string

// *italic* - for
type italic string

// *emphasis* - for emphasis
type em string

// interface for title elements -
type Heading interface {
	isHeading()
	String() string
}

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

// The text content interface is used to enforce that the content of a section
// is of type TextContent, such as p, a, em, b, i
type TextContent interface {
	isTextContent()
	String() string
}

// Section struct to represent a section with a title and content
// A section contains at least one subtitle and can contain multiple subtitles,
// Example: Section with a h2 will continue until the next h2 or the end of the document
type Section struct {
	Title       Heading
	TextContent []TextContent
}

// MarkdownDocument struct represents a complete markdown document.
type MarkdownDocument struct {
	Metadata Metadata
	Sections []Section
}
