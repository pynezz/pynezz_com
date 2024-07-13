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
// Define common types
type p string
type a string

type img string

type em string
type b string
type strong string
type italic string

type code string
type codeblock string

type h1 string

type pre string

type blockquote string

type hr string
type br string

type table string
type tr string
type th string
type td string

type ul string
type ol string
type li string

func (h heading) isHeading() {}
func (h heading) String() string {
	return h.content
}
func (h heading) HTMLTag() string {
	return h.tag
}

// TextContent interface for content elements.
type TextContent interface {
	isTextContent()
	String() string
}

// General text content struct to embed common text content behavior
type textContent struct {
	content string
}

func (t textContent) isTextContent() {}
func (t textContent) String() string {
	return t.content
}

// Specific text content types embedding textContent
type textP struct{ textContent }
type textA struct{ textContent }
type textImg struct{ textContent }
type textEm struct{ textContent }
type textB struct{ textContent }
type textStrong struct{ textContent }
type textItalic struct{ textContent }
type textCode struct{ textContent }
type textCodeblock struct{ textContent }
type textPre struct{ textContent }
type textBlockquote struct{ textContent }
type textHr struct{ textContent }
type textBr struct{ textContent }
type textTable struct{ textContent }
type textTr struct{ textContent }
type textTh struct{ textContent }
type textTd struct{ textContent }
type textUl struct{ textContent }
type textOl struct{ textContent }
type textLi struct{ textContent }

// General heading struct to embed common heading behavior
type heading struct {
	content string
	tag     string
}

// Specific heading types embedding heading
type h2 struct{ heading }
type h3 struct{ heading }
type h4 struct{ heading }
type h5 struct{ heading }
type h6 struct{ heading }

// interface for title elements -
type Heading interface {
	isHeading()
	String() string
	HTMLTag() string
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
