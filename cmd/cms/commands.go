package cms

import (
	"fmt"

	ansi "github.com/pynezz/pynezzentials/ansi"
)

var c = map[string]Command{
	"list":      Command{HelpString: "List all the pages", Run: listPages},
	"page":      Command{HelpString: "Show a specific page", Run: showPage},
	"create":    Command{HelpString: "Create a new page. Takes a markdown filepath as input", Run: createPage},
	"edit":      Command{HelpString: "Edit an existing page", Run: editPage},
	"delete":    Command{HelpString: "Delete a page", Run: deletePage},
	"publish":   Command{HelpString: "Publish a page", Run: publishPage},
	"unpublish": Command{HelpString: "Unpublish a page", Run: unpublishPage},
	"status":    Command{HelpString: "Get the status of a page", Run: statusPage},
	"tags":      Command{HelpString: "Get the tags of a page", Run: tagsPage},
	"stats":     Command{HelpString: "Get the statistics of the CMS", Run: statsCMS},
	"config":    Command{HelpString: "Show the configuration", Run: configCMS},
	"push":      Command{HelpString: "Push the CMS to the server", Run: pushCMS},
	"sync":      Command{HelpString: "Sync the changes to the server", Run: syncCMS},
}

func createPage(path string) any {
	var str string
	var err error
	if str, err = ansi.SprintHexf("#ff0000", "Hello from the createPage function!"); err != nil {
		return err
	}

	fmt.Println(str)

	return "Hello from the createPage function!"
}
