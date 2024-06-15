package parser

import (
	"bufio"
	"fmt"
	"path"

	ansi "github.com/pynezz/pynezzentials/ansi"
	fsutil "github.com/pynezz/pynezzentials/fsutil"
)

// Markdown to HTML
// NB! Might need to account for encoding issues (CRLF vs LF etc.)
func MarkdownToHTML(mdPath string) []byte {
	metadata := &Metadata{}
	content := &[]byte{}

	mdFile, err := fsutil.GetFile(mdPath)
	if err != nil {
		return []byte{}
	}

	o, _ := ansi.SprintHexf("#7BD4F1", "Parsing markdown file "+path.Base(mdPath)+" to HTML")
	fmt.Println(o)

	r := bufio.NewReader(mdFile)
	md, _ := r.ReadBytes('\n')

	if mdatabytes, err := readMetadata(md); err == nil {
		*metadata, err = parseMetadata(mdatabytes)
		if err != nil {
			return []byte{}
		}
		ansi.PrintSuccess("Metadata parsed successfully")
	} else {
		return []byte("error: encountered an error while parsing metadata\n")
	}

	*content = []byte(parseContent(mdPath))

	return md
}

// TODO: Use some library here? Or write my own...?
func parseContent(file string) string {
	return ""
}
