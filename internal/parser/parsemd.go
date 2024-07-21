package parser

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	// codeRe := regexp.MustCompile("```(\\w*)\n([\\s\\S]*?)\n```")
	// contentParts := splitMarkdownContent(contentStr)
	// for _, part := range contentParts {
	// 	ansi.PrintInfo("part: " + part)
	// }
	re := regexp.MustCompile("\n\n|\n-|---\n")
	contentParts := re.Split(contentStr, -1)
	// for _, part := range contentParts {
	// 	contentParts = append(contentParts, codeRe.FindAllString(part, -1)...)
	// }

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

func log(content string) {
	fileName := "parser_log.log"
	if !fsutil.FileExists(fileName) {
		fsutil.CreateFile(fileName)
	}
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		ansi.PrintError("error opening log file")
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		ansi.PrintError("error writing to log file")
	}
}

// parseContent parses the content into a MarkdownDocument struct
func parseContent(lines []string) *MarkdownDocument {
	md := &MarkdownDocument{}
	var currentTitle Heading
	var currentContent []TextContent

	// Find links in the form [text](url)
	linkPattern := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	// Find inline code in the form `content`
	inlineCode := regexp.MustCompile("`([^`]+)`")

	codeParser := &CodeblockParser{
		inCodeBlock: false,
		blockCount:  0,
	}

	skipLines := 0

	for i, line := range lines {
		// line = strings.TrimSpace(line)

		log(fmt.Sprintf("line %d: %s\n", i, line))

		if len(line) == 0 {
			continue
		}

		//  choose to skip lines if a set amount of lines will be processed elsewhere
		if skipLines > 0 {
			ansi.PrintDebug("skipping lines: " + strconv.FormatInt(int64(skipLines), 10))
			skipLines--
			continue
		}

		switch {
		case strings.HasPrefix(line, "```"):
			firstLine := strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "```")), "\n")
			codeParser.codeBlockLang = firstLine[0] // Capture the language of the code block

			ansi.PrintColor(ansi.Cyan, fmt.Sprintf("code block first line: %s", strings.Join(firstLine[1:], "\n")))
			firstContentLine := strings.Join(firstLine[1:], "\n")
			firstContentLine = strings.TrimSpace(firstContentLine) + "\n"

			codeParser.content = append(codeParser.content, strings.Split(firstContentLine, "\n")...)
			if codeParser.codeBlockLang == "" {
				codeParser.codeBlockLang = "plaintext"
			}

		linesCheck:
			for _, l := range lines[i+1:] {
				skipLines++
				for _, substr := range strings.Split(l, "\n") {
					if len(substr) >= 3 {
						if substr[0:3] == "```" {
							ansi.PrintInfo("end of code block found")
							codeParser.content = append(codeParser.content, strings.Trim(substr, "`"))
							break linesCheck
						}
					}
					codeParser.content = append(codeParser.content, syntaxHighlight(substr))
				}
			}

			currentContent = append(
				currentContent,
				textCodeblock{
					textContent{
						parseCodeBlock(strings.Join(codeParser.content, "\n"),
							codeParser.codeBlockLang) + "\n",
					},
				},
			)
			n := &CodeblockParser{
				blockCount: codeParser.blockCount + 1,
			}
			codeParser = n
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
		ansi.PrintColorBold(ansi.Cyan, "adding section: "+currentTitle.String())
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

	return re.ReplaceAllString(line, "<code class=\"bg-surface0 text-subtext1 p-1 rounded-md\">$1</code>")
}

func syntaxHighlight(code string) string {
	t := GetGoTypes()
	line := strings.Split(code, " ")
	for _, word := range line {
		if t.GetType(word) != "" {
			code = strings.ReplaceAll(code, word, fmt.Sprintf(`<span class="%s">%s</span>`, t.GetType(word), word))
		}
	}

	return code
}

func parseCodeBlock(content, lang string) string {
	ansi.PrintColor(ansi.Cyan, "parseCodeBlock!")
	// Escape HTML special characters
	content = strings.ReplaceAll(content, "&", "&amp;")
	// content = strings.ReplaceAll(content, "<", "&lt;")
	// content = strings.ReplaceAll(content, ">", "&gt;")
	content = strings.ReplaceAll(content, "\n", "<br>")

	ansi.PrintColor(ansi.Cyan, "lang: "+lang)
	ansi.PrintColor(ansi.Yellow, "content: "+content)
	return fmt.Sprintf(`<pre class="bg-crust overflow-x-auto rounded-md"><code id="language-%s" class="text-sm w-max">%s</code></pre>`, lang, content)
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

func cssRel() string {
	return `<link rel="stylesheet" type="text/css" href="/css/post.css">`
}

func nav() string {
	return `<ul class="main-nav">
		<li class="nav-item">
			<a href="/" class="nav-link">/</a>
		</li>
    <li class="nav-item">
      <a href="/posts/" class="nav-link">posts</a>
    </li>
</ul>`
}
