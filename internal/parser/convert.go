package parser

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

/*
DECLARATION OF HEADING ELEMENTS
We want to enforce heading to be of type h2...h6
*/
func (h h2) isHeading() {}
func (h h2) String() string {
	return string(h)
}

func (h h3) isHeading() {}
func (h h3) String() string {
	return string(h)
}

func (h h4) isHeading() {}
func (h h4) String() string {
	return string(h)
}

func (h h5) isHeading() {}
func (h h5) String() string {
	return string(h)
}

func (h h6) isHeading() {}
func (h h6) String() string {
	return string(h)
}

// The text content interface is used to enforce that the content of a section
// is of type TextContent, such as p, a, em, b, i
type TextContent interface {
	isTextContent()
	String() string
}

func (t p) isTextContent() {}
func (t p) String() string {
	return string(t)
}

func (t a) isTextContent() {}
func (t a) String() string {
	return string(t)
}

func (t img) isTextContent() {}
func (t img) String() string {
	return string(t)
}

func (t em) isTextContent() {}
func (t em) String() string {
	return string(t)
}

func (t b) isTextContent() {}
func (t b) String() string {
	return string(t)
}

// A section contains at least one subtitle and can contain multiple subtitles,
// section with a h2 will continue until the next h2 or the end of the document
type section struct {
	Title[h2, h3, h4, h5, h6]
}
