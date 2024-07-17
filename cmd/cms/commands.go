package cms

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pynezz/pynezz_com/internal/helpers"
	"github.com/pynezz/pynezz_com/internal/parser"
	"github.com/pynezz/pynezz_com/internal/server/middleware"
	"github.com/pynezz/pynezz_com/internal/server/middleware/models"
	ansi "github.com/pynezz/pynezzentials/ansi"
	"github.com/pynezz/pynezzentials/fsutil"
	"gorm.io/datatypes"
)

// Declaration of the Commands goes here
// REMEMBER: When adding a new command:
// 1. Add the command to the map in the commands() function
// 2. Implement the Run method for the command (in implCommands.go)
// 3. Implement the Help method for the command (in implCommands.go)
// 4. Implement the Run method for the command (in this file) - this is the actual implementation of the command
// - example: func listPages(min, max int) []parser.Page { ... }
// 5. Implement the Help method for the command (in this file) - this is the help message of the command
// - example: func (c *ListPages) Help() string { ... }
// 6. Implement the Run method for the command (in this file) - this is the actual implementation of the command
// - example: func (c *ListPages) Run(args ...interface{}) interface{} { ... }
// 7. Implement the Help method for the command (in this file) - this is the help message of the command
// - example: func (c *ListPages) Help() string { ... }
// 8. Add the command to the map in the commands() function
// - example: prefix + "list": c["list"],
// 9. Add the command to the switch statement in the Run method of the Command struct (in implCommands.go)
// - example: case "list": run(c["list"], args...)
var c = map[string]ICommand{
	"edit":      &EditPage{&Command{HelpStr: "Edit a page", NameStr: "edit"}},
	"parse":     &ParseAll{&Command{HelpStr: "Parse and build all pages", NameStr: "parse"}},
	"build":     &ParseAll{&Command{HelpStr: "Parse and build all pages", NameStr: "parse"}}, // alias for parse
	"list":      &ListPages{&Command{HelpStr: "List all pages", NameStr: "list"}},
	"create":    &CreatePage{&Command{HelpStr: "Create a page", NameStr: "create"}},
	"delete":    &DeletePage{&Command{HelpStr: "Delete a page", NameStr: "delete"}},
	"publish":   &PublishPage{&Command{HelpStr: "Publish a page", NameStr: "publish"}},
	"unpublish": &UnpublishPage{&Command{HelpStr: "Unpublish a page", NameStr: "unpublish"}},
	"status":    &PageStatus{&Command{HelpStr: "Show the status of a page", NameStr: "status"}},
	"tags":      &PageTags{&Command{HelpStr: "Show the tags of a page", NameStr: "tags"}},
	"page":      &ShowPage{&Command{HelpStr: "Show a page", NameStr: "page"}},
	"config":    &Config{&Command{HelpStr: "Show the config of a page", NameStr: "config"}},

	"nop": &Nop{&Command{HelpStr: "Noop", NameStr: "nop"}}, // for typo checking
}

var validCommands = []string{"list", "edit", "parse", "build", "create", "delete", "publish", "unpublish", "status", "tags", "config", "page"}

func noop() bool {
	fmt.Println("Noop called")
	return false
}

func parseAll() bool {
	fmt.Println("ParseAll called")

	// Read "content/*.md" files
	// Parse the content
	// Create the pages
	// Write the pages to "public/*.html" files
	// Return true if successful, false otherwise
	// Use the parser package for this

	files, err := fsutil.GetFiles("content")
	if err != nil {
		ansi.PrintError("error reading contents of 'content' directory")
	}

	for _, file := range files {
		// check if the file is already parsed:
		// if it is, skip it
		// if it is not, parse it

		if isParsed(file) {
			ansi.PrintInfo("file already parsed: " + file)
			continue
		}

		ansi.PrintDebug("file is not yet parsed: " + file)
		bytes, doc := parser.MarkdownToHTML(file)
		if bytes == nil {
			ansi.PrintError("error parsing file: " + file)
			return false
		}

		// write the parsed content to a file
		// the file should be in the "public" directory
		newName := filenameConvert(file)
		ansi.PrintDebug("newName: " + newName)
		f, err := fsutil.CreateFile("pynezz/public/" + newName)
		if err != nil {
			ansi.PrintError("error writing parsed content to file: " + newName)
			return false
		}
		written, err := f.Write(bytes)
		if err != nil {
			ansi.PrintError("error writing parsed content to file: " + newName)
			return false
		}

		ansi.PrintInfo(fmt.Sprintf("parsed content written to file: %s (%d bytes)", newName, written))

		// write to database
		// the database should be in the "db" directory
		// post := middleware.ContentsDB.GenerateMetadata(bytes)
		post := parser.Post{
			Metadata: doc.Metadata,
			// Content:  []byte(doc.String()),
		}
		slug := middleware.ContentsDB.GenerateSlug(post.Metadata.Title)
		ansi.PrintDebug("generated slug: " + slug)
		postMetadata := models.PostMetadata{
			Title: post.Metadata.Title,
			Path:  newName,
			Slug:  slug,

			// PostID: int(crc32.ChecksumIEEE([]byte(slug))), // always unique and reproducible.
			// Changed to Adler32 - check the hash_bench.go in root directory for explanation.
			PostID: int(helpers.Adler32(slug)), // always unique and reproducible
			Tags:   datatypes.JSON(strings.Join(post.Metadata.Tags, ",")),
		}

		// write to database
		if err := middleware.ContentsDB.NewPost(postMetadata); err != nil {
			ansi.PrintError("error writing to database")
			return false
		}

		contentBytes := []byte{}
		for _, b := range doc.Sections {
			contentBytes = append(contentBytes, []byte(b.String())...)
		}

		if err := middleware.ContentsDB.WriteContentsToDatabase(slug, contentBytes); err != nil {
			ansi.PrintError("error writing to database")
			return false
		}
		ansi.PrintInfo("written to database")

		// ! IMPORTANT: parser and models have different ways of handling the same data

	}

	return true
}

func showPage(id string) bool {
	return false
}

func createPage(path string) bool {
	var str string
	var err error
	if str, err = ansi.SprintHexf("#7BD4F1", "Hello from the createPage function!"); err != nil {
		return false
	}

	fmt.Println(str)

	return true
}

func listPages(min, max int) []parser.Page {
	var str string
	var err error
	if str, err = ansi.SprintHexf("#7BD4F1", "Hello from the listPages function!"); err != nil {
		return nil
	}

	fmt.Println(str)

	return nil
}

func editPage(id string) bool {
	ansi.PrintInfo("edit page called with param: " + id)

	return false
}

func deletePage(id string) bool {
	ansi.PrintInfo("delete page called with param: " + id)

	return false
}

func publishPage(id string) bool {
	ansi.PrintInfo("publish page called with param: " + id)

	return false
}

func unpublishPage(id string) bool {
	ansi.PrintInfo("unpublish page called with param: " + id)

	return false
}

func showPageStatus(id string) bool {
	ansi.PrintInfo("show page status called with param: " + id)

	return false
}

func showPageTags(id string) bool {
	ansi.PrintInfo("show page tags called with param: " + id)

	return false
}

func showPageConfig(id string) bool {
	ansi.PrintInfo("show page config called with param: " + id)

	return false
}

func pageStatus(id string) bool {
	ansi.PrintInfo("page status called with param: " + id)

	return false
}

func pageTags(id string) bool {
	ansi.PrintInfo("page tags called with param: " + id)

	return false
}

func config(id string) bool {
	ansi.PrintInfo("page config called with param: " + id)

	return false
}

func isParsed(file string) bool {
	// check the "public" directory for the file
	return fsutil.FileExists("pynezz/public/" + filenameConvert(file)) // if the file exists, it is parsed
}

// filenameConvert converts a markdown filename to an html filename
// and vice versa
// example: filenameConvert("file.md") -> "file.html"
// example: filenameConvert("file.html") -> "file.md"
func filenameConvert(file string) string {
	fileName := strings.Split(filepath.Base(file), ".")[0]
	fileType := filepath.Ext(file)

	ansi.PrintDebug("converting filename: " + fileName + fileType)
	if fileType == ".md" {
		fmt.Println("Converting " + fileName + " to " + fileName + ".html")
		return strings.Split(fileName, ".")[0] + ".html"
	}

	fmt.Println("Converting " + fileName + " to " + strings.TrimSuffix(fileName, fileType) + ".html")
	return strings.TrimSuffix(fileName, fileType) + ".md"
}
