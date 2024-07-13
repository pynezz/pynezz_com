package parser

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"

	fsutil "github.com/pynezz/pynezzentials/fsutil"
)

// MarkdownToHTML converts markdown to HTML
func MarkdownToHTML(mdPath string) []byte {
	metadataStr, contentStr, err := parseMarkdownFile(mdPath)
	if err != nil {
		return []byte{}
	}

	metadata, err := ParseMetadata([]byte(metadataStr))
	if err != nil {
		return []byte("error: encountered an error while parsing metadata\n")
	}

	document := parseContent(strings.Split(contentStr, "\n"))
	document.Metadata = metadata

	return []byte(document.String())
}

// parseMarkdownFile reads the markdown file and separates metadata from content
func parseMarkdownFile(mdPath string) (string, string, error) {
	mdFile, err := fsutil.GetFile(mdPath)
	if err != nil {
		return "", "", err
	}
	defer mdFile.Close()

	scanner := bufio.NewScanner(mdFile)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return "", "", scanner.Err()
	}

	mdContent := strings.Join(lines, "\n")
	metadata, content := extractMetadata(mdContent)
	return metadata, content, nil
}

// parseContent parses the content into a MarkdownDocument struct
func parseContent(lines []string) *MarkdownDocument {
	md := &MarkdownDocument{}
	var currentTitle Heading
	var currentContent []TextContent

	// Find links in the form [text](url)
	linkPattern := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		switch {
		case strings.HasPrefix(line, "## "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h2{heading{line[3:], "h2"}}
		case strings.HasPrefix(line, "### "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h3{heading{line[4:], "h3"}}
		case strings.HasPrefix(line, "#### "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h4{heading{line[5:], "h4"}}
		case strings.HasPrefix(line, "##### "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h5{heading{line[6:], "h5"}}
		case strings.HasPrefix(line, "###### "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h6{heading{line[7:], "h6"}}
		case strings.HasPrefix(line, "|"):
			currentContent = append(currentContent, parseTable(line))
		// case strings.Contains(line, "[") && strings.Contains(line, "]("):
		// 	currentContent = append(currentContent, parseLink(line))
		default:
			line = linkPattern.ReplaceAllStringFunc(line, func(match string) string {
				return parseLink(match).String()
			})
			currentContent = append(currentContent, textP{textContent{line}})
		}
	}

	if currentTitle != nil {
		md.AddSection(currentTitle, currentContent)
	}

	return md
}

// parseTable parses a markdown table line into HTML
func parseTable(line string) TextContent {
	return textTable{textContent{fmt.Sprintf("<table><tr>%s</tr></table>", line)}}
}

// parseLink parses a markdown link into an HTML link
func parseLink(line string) TextContent {
	s := strings.Split(line, "](")
	if len(s) != 2 {
		return textA{textContent{""}}
	}

	link := s[1]
	link = strings.TrimSuffix(link, ")")

	lText := s[0]
	lText = strings.TrimPrefix(lText, "[")

	return textA{textContent{fmt.Sprintf(`<a href="%s">%s</a>`, link, lText)}}
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

func (md *MarkdownDocument) String() string {
	// Add creation date and title
	document := fmt.Sprintf("<h1>%s</h1>\n<p class=\"date\">%s</p>\n", md.Metadata.Title, md.Metadata.Date.Format("02.01.2006"))
	for _, section := range md.Sections {
		document += section.String()
	}

	// Adding tags at the bottom
	if len(md.Metadata.Tags) > 0 {
		document += "<p>Tags: "
		for _, tag := range md.Metadata.Tags {
			document += fmt.Sprintf(`<a class='tag' href='/tags/%s'>%s</a> `, tag, tag)
		}
		document += "</p>\n"
	}

	// document += css()

	prepend := "<!DOCTYPE html>\n<html>\n<head>\n<meta charset='utf-8'>\n<title>" + md.Metadata.Title + "</title>" + cssRel() + "\n</head>\n<body>\n"
	append := "</body>\n</html>"
	document = prepend + nav() + document + append

	return document
}

func cssRel() string {
	return `<link rel="stylesheet" type="text/css" href="/css/post.css">`
}

func nav() string {
	return fmt.Sprintf(` <ul class="main-nav">
		<li class="nav-item">
			<a href="/" class="nav-link">/</a>
		</li>
    <li class="nav-item">
      <a href="/posts/" class="nav-link">posts</a>
    </li>
</ul>`)
}

func css() string {
	return `
<style>
	.date {
		font-size: 0.8em;
		color: #666;
	}
	.tag {
		font-size: 0.8em;
		color: #666;
	}
	table {
		width: 100%;
		border-collapse: collapse;
	}
	th, td {
		border: 1px solid #ddd;
		padding: 8px;
	}
	th {
		background-color: #f2f2f2;
	}
	body {
		font-family: Inter, sans-serif;
	}
	h1 {
		font-size: 2em;
		color: #333;
	}
	code {
		background-color: #f2f2f2;
		padding: 2px;
	  width: 100%;
		box-sizing: border-box;
		border-radius: 1rem;
	}
	code.inline {
		background-color: #f2f2f2;
		padding: 2px;
		border-radius: .25rem;
	}
	html {
		font-size: 16px;
		background-color: #f9f9f9;
		display: flex;
		justify-content: center;
	}
	article {
		width: 100%;
		max-width: 800px;
		padding: 1rem;
	}
	section {
		margin: 1rem 0;
	}
</style>`
}
