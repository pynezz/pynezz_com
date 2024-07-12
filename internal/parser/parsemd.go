package parser

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"path"
	"strings"

	ansi "github.com/pynezz/pynezzentials/ansi"
	fsutil "github.com/pynezz/pynezzentials/fsutil"
)

// Markdown to HTML
// NB! Might need to account for encoding issues (CRLF vs LF etc.)
func MarkdownToHTML(mdPath string) []byte {
	metadata := &Metadata{}
	document := &MarkdownDocument{}

	mdFile, err := fsutil.GetFile(mdPath)
	if err != nil {
		return []byte{}
	}
	defer mdFile.Close()

	o, _ := ansi.SprintHexf("#7BD4F1", "Parsing markdown file "+path.Base(mdPath)+" to HTML")
	fmt.Println(o)

	scanner := bufio.NewScanner(mdFile)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return []byte{}
	}

	mdContent := strings.Join(lines, "\n")
	mdatabytes, err := readMetadata([]byte(mdContent))
	if err == nil {
		*metadata, err = ParseMetadata(mdatabytes)
		if err != nil {
			return []byte{}
		}
		ansi.PrintSuccess("Metadata parsed successfully")
	} else {
		return []byte("error: encountered an error while parsing metadata\n")
	}

	document = parseContent(strings.Split(mdContent, "\n"))
	document.Metadata = *metadata

	return []byte(document.String())

}

// // readContents ignores the metadata and reads only the contents
// func readContents(f *os.File) string {
// 	reader := &bufio.Reader{}
// 	reader = bufio.NewReader(f)
// 	contents := ""

// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err != nil {
// 			break
// 		}

// 		contents += line
// 	}

// }

// TODO: Use some library here? Or write my own...?
// I'm writing my own. WIP.
func parseContent(lines []string) *MarkdownDocument {
	md := MarkdownDocument{}
	var currentTitle Heading
	var currentContent []TextContent

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "## ") {
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h2(line[3:])
		} else if strings.HasPrefix(line, "### ") {
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h3(line[4:])
		} else if strings.HasPrefix(line, "#### ") {
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h4(line[5:])
		} else if strings.HasPrefix(line, "##### ") {
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h5(line[6:])
		} else if strings.HasPrefix(line, "###### ") {
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h6(line[7:])
		} else {
			currentContent = append(currentContent, p(line))
		}
	}

	if currentTitle != nil {
		md.AddSection(currentTitle, currentContent)
	}

	return &md
}

// ParseDescription sets the description of a post if it exists, otherwise it returns the first 100 characters of the content
func SetDescription(p *Post) {
	if p.Metadata.Description != "" {
		return
	}

	if len(p.Content) < 100 {
		p.Metadata.Description = string(p.Content)
		return
	}

	p.Metadata.Description = string(p.Content[:100])
}

// NewPost creates a new Post.
func NewPost(content string) *Post {
	p := &Post{
		Content: []byte(content),
	}

	return &Post{
		Content: p.Content,
		CalculateSha256Sum: func() string {
			if p.sha256sum == "" && p.Content != nil {
				hash := sha256.New()
				hash.Write(p.Content)
				p.sha256sum = fmt.Sprintf("%x", hash.Sum(nil))
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

// genToc generates a table of contents for a markdown file
func genToc() {}

func (md *MarkdownDocument) AddSection(title Heading, content []TextContent) {
	section := Section{
		Title:       title,
		TextContent: content,
	}
	md.Sections = append(md.Sections, section)
}

func (md MarkdownDocument) String() string {
	document := ""
	for _, section := range md.Sections {
		document += section.String()
	}
	return document
}
