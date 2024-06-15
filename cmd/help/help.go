package help

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pynezz/pynezz_com/cmd/cms"
	"github.com/pynezz/pynezz_com/cmd/serve"
)

var usage = func() string {

	// theoretical usage function - for testing
	return `Usage:
	%s cms [options]
Options:
	--help		Print this help message
	--version	Print the version of the CMS
	--list		List all the pages
	--page		Show a specific page
	--create	Create a new page
	--edit		Edit an existing page
	--delete	Delete a page
	--publish	Publish a page
	--unpublish	Unpublish a page
	--status	Get the status of a page
	--tags		Get the tags of a page
	--tag		Get the pages with a specific tag
	--search	Search for a page
	--config	Show the configuration
	--set		Set a configuration value
	--unset		Unset a configuration value`
}

func Help(args ...string) string {

	return displayHelp(args...)
}

var displayHelp func(args ...string) string = func(args ...string) string {
	if len(args) < 1 {
		return ""
	}

	fmt.Println("displayHelp: ", args)

	h := map[string]func(...string) string{
		"cms":   cms.Help,
		"serve": serve.Help,
	}

	return fmt.Sprintf(`Usage:
	%s %s [options]
Options:
	%s`, filepath.Base(os.Args[0]), args[1], h[args[1]](args...))
}
