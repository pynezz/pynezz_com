package parser

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"

	"github.com/pynezz/pynezz_com/templates/layout"
	"github.com/pynezz/pynezzentials/ansi"
	fsutil "github.com/pynezz/pynezzentials/fsutil"
)

// Constants for splitting the markdown content
// const splLst = "## |### |#### |##### |###### |\\|"
// const splNewLine = "\n\n"
// const splitCond = splLst + "|" + splNewLine

// MarkdownToHTML converts markdown to HTML
func MarkdownToHTML(mdPath string) ([]byte, MarkdownDocument) {
	metadataStr, contentStr, err := parseMarkdownFile(mdPath)
	if err != nil {
		return []byte{}, MarkdownDocument{}
	}

	metadata, err := ParseMetadata([]byte(metadataStr))
	if err != nil {
		return []byte("error: encountered an error while parsing metadata\n"), MarkdownDocument{}
	}

	// re := regexp.MustCompile(`|## |### |#### |##### |###### |\\||\n\n|\n-|---\n`)
	re := regexp.MustCompile("\n\n|\n-|---\n")

	// contentParts := splitMarkdownContent(contentStr)
	// for _, part := range contentParts {
	// 	ansi.PrintInfo("part: " + part)
	// }
	contentParts := re.Split(contentStr, -1)

	ansi.PrintInfo("parsing content...")
	document := parseContent(contentParts)
	document.Metadata = metadata

	return []byte(document.String()), *document
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

// splitMarkdownContent splits markdown content by multiple delimiters
func splitMarkdownContent(content string) []string {
	re := regexp.MustCompile("```(\\w*)\n([\\s\\S]*?)\n```")
	for _, match := range re.FindAllString(content, -1) {
		ansi.PrintColorUnderline(ansi.Cyan, "match: "+match)
	}

	return re.FindAllString(content, -1)
}

func extractCodeBlocks(content string) []string {
	ansi.PrintDebug("extracting code blocks from: " + content)
	var backticks int
	codeblocks := []string{}
	startLoc := 0

	for i, r := range content {
		if content[r] == '`' {
			backticks++
		}
		if backticks == 3 {
			// code block found
			// find the next 3 backticks
			startLoc = i
			// add the content between the backticks to the code blocks slice
			// reset backticks to 0
		}
		if backticks == 6 {
			// code block found
			// find the next 3 backticks
			codeblocks = append(codeblocks, content[startLoc:i-3])
		}
		return codeblocks
	}

	return codeblocks
}

// parseContent parses the content into a MarkdownDocument struct
func parseContent(lines []string) *MarkdownDocument {
	md := &MarkdownDocument{}
	var currentTitle Heading
	var currentContent []TextContent

	ansi.PrintDebug("checking content: " + strings.Join(lines, "\n"))

	// Find links in the form [text](url)
	linkPattern := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	codes := regexp.MustCompile("```(\\w*)\n([\\s\\S]*?)\n```")
	// Find code blocks in the form ```lang\ncontent\n```
	codePattern := regexp.MustCompile("```(\\w*)\n")
	if codePattern == nil {
		ansi.PrintError("error: could not compile code pattern")
	}

	// Find inline code in the form `content`
	inlineCode := regexp.MustCompile("`([^`]+)`")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		switch {
		case codePattern.MatchString(line):
			ansi.PrintColor(ansi.Cyan, "code block!")
			codeBlocks := extractCodeBlocks(strings.Join(lines, "\n"))
			for _, block := range codeBlocks {
				currentContent = append(currentContent, textCodeblock{textContent{parseCodeBlock(block, "lang")}})
			}			
			// matches := codePattern.FindStringSubmatch(line)
			// currentContent = append(currentContent, textCodeblock{textContent{parseCodeBlock(matches[2], matches[1])}})
		case strings.HasPrefix(line, "## "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h2{heading{line[3:], "h2", H2Style}}
		case strings.HasPrefix(line, "### "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h3{heading{line[4:], "h3", H3Style}}
		case strings.HasPrefix(line, "#### "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h4{heading{line[5:], "h4", H4Style}}
		case strings.HasPrefix(line, "##### "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h5{heading{line[6:], "h5", H5Style}}
		case strings.HasPrefix(line, "###### "):
			if currentTitle != nil {
				md.AddSection(currentTitle, currentContent)
				currentContent = nil
			}
			currentTitle = h6{heading{line[7:], "h6", H6Style}}
		case strings.HasPrefix(line, "|"):
			currentContent = append(currentContent, parseTable(line))
			// case strings.Contains(line, "[") && strings.Contains(line, "]("):
			// 	currentContent = append(currentContent, parseLink(line))
		case strings.HasPrefix(line, "|"):
			currentContent = append(currentContent, parseTable(line))
		case strings.HasPrefix(line, "[") && strings.Contains(line, "]("):
			currentContent = append(currentContent, parseLink(line))
		case strings.HasPrefix(line, "- "):
			currentContent = append(currentContent, parseList(line, "ul"))
		case strings.HasPrefix(line, "1. "):
			currentContent = append(currentContent, parseList(line, "ol"))

		case inlineCode.MatchString(line):
			currentContent = append(currentContent, textCode{textContent{parseInlineCode(line)}})

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

// parseList parses a markdown list into an HTML list
func parseList(line string, listType string) TextContent {
	var itemTag string
	if listType == "ul" {
		itemTag = "ul"
	} else {
		itemTag = "ol"
	}

	listItems := strings.Split(line, "\n")
	listHTML := fmt.Sprintf("<%s>", itemTag)
	for _, item := range listItems {
		listHTML += fmt.Sprintf("<li>%s</li>", strings.TrimSpace(item))
	}
	listHTML += fmt.Sprintf("</%s>", itemTag)

	return textContent{content: listHTML}
}

func parseInlineCode(line string) string {
	ansi.PrintBold("inline code: " + line)
	re := regexp.MustCompile("`([^`]+)`")

	return re.ReplaceAllString(line, "<code class=\"bg-gray-200 text-gray-800 px-1 py-0.5 rounded\">$1</code>")
}

func parseCodeBlock(content, lang string) string {
	ansi.PrintColor(ansi.Cyan, "parseCodeBlock!")
	// Escape HTML special characters
	content = strings.ReplaceAll(content, "&", "&amp;")
	content = strings.ReplaceAll(content, "<", "&lt;")
	content = strings.ReplaceAll(content, ">", "&gt;")
	content = strings.ReplaceAll(content, "\n", "<br>\n")

	ansi.PrintColor(ansi.Cyan, "lang: "+lang)
	ansi.PrintColor(ansi.Yellow, "content: "+content)
	return fmt.Sprintf(`<pre class="bg-gray-900 text-text p-4 rounded-lg overflow-x-auto"><code class="language-%s">%s</code></pre>`, lang, content)
}

// func parseCodeBlock(block string) string {
// 	re := regexp.MustCompile("(?s)```(\\w*)\\n(.*?)\\n```")
// 	match := re.FindStringSubmatch(block)
// 	ansi.PrintBold("code block: " + block)
// 	if len(match) != 3 {
// 		ansi.PrintError("error: could not parse code block")
// 		return ""
// 	}

// 	lang := match[1]
// 	content := match[2]

// 	// Escape HTML special characters
// 	content = strings.ReplaceAll(content, "&", "&amp;")
// 	content = strings.ReplaceAll(content, "<", "&lt;")
// 	content = strings.ReplaceAll(content, ">", "&gt;")
// 	content = strings.ReplaceAll(content, "\n", "<br>\n")

// 	ansi.PrintColor(ansi.Cyan, "lang: "+lang)
// 	ansi.PrintColor(ansi.Yellow, "content: "+content)
// 	return fmt.Sprintf(`<pre class="bg-gray-900 text-text p-4 rounded-lg overflow-x-auto"><code class="language-%s">%s</code></pre>`, lang, content)
// }

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

	return textA{textContent{fmt.Sprintf(`<a href="%s" class="%s">%s</a>`, link, layout.Link, lText)}}
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
	document := fmt.Sprintf("<article class=\"content\"><h1 class=\"font-sans\">%s</h1>\n<p class=\"date\">%s</p>\n", md.Metadata.Title, md.Metadata.Date.Format("02.01.2006"))
	for _, section := range md.Sections {
		document += section.String()
	}

	// Adding tags at the bottom
	if len(md.Metadata.Tags) > 0 {
		document += "<p class=\"p-2 m-2 rounded bg-mantle text-green\">Tags: "
		for _, tag := range md.Metadata.Tags {
			document += fmt.Sprintf(`<a class='tag' href='/tags/%s'>%s</a> `, tag, tag)
		}
		document += "</p>\n"
	}

	// document += css()

	prepend := "<!DOCTYPE html>\n<html>\n<head>\n<meta charset='utf-8'>\n<title>" + md.Metadata.Title + "</title>" + cssRel() + "\n</head>\n<body>\n"
	append := "\n</article>\n</body>\n</html>"
	document = prepend + nav() + document + append

	return document
}

func genCode() {
	// codeBlock := regexp.MustCompile("```")
	// codeInline := regexp.MustCompile("`")

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
