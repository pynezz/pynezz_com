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

func (h heading) Class() string {
	return h.style
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

// func (c code) String() string {
// 	return c.content
// }

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
	style   string
}

type code struct {
	content string
	tag     string
	style   string
}

const c = ""
const cEnd = ""

// const c = "class =\""
// const cEnd = "\""
const H2Style = c + "text-xl text-subtext1 underline mb-0.5 pt-4 "
const H3Style = c + "text-subtext0 font-bold text-xl font-sans mb-1 pt-4" + cEnd
const H4Style = c + "text-maroon font-lg font-sans mb-1 pt-3" + cEnd
const H5Style = c + "text-sky font-md font-sans mb-1" + cEnd
const H6Style = c + "text-mauve font-sm font-sans mb-1" + cEnd

// const CodeBlockStyle = c + "text-text font-mono font-md bg-mantle p-4 rounded-md" + cEnd

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
	Class() string
}

type Code interface {
	isCode()
	HTMLTag() string
	Class() string
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
	// Style       Class
}

// MarkdownDocument struct represents a complete markdown document.
type MarkdownDocument struct {
	Metadata Metadata
	Sections []Section
}

type Class struct {
	Tailwind string
}

// CodeblockParser struct to parse codeblocks in markdown
type CodeblockParser struct {
	inCodeBlock   bool     // whether the parser is currently in a codeblock
	blockCount    int      // the amount of codeblocks in the document
	content       []string // the content of the codeblock
	codeBlockLang string   // the language of the codeblock
}
