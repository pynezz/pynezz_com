package cms

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pynezz/pynezzentials/ansi"
)

func needHelp(arg string) bool {
	return arg == "help" || arg == "--help" || arg == "-h"
}

var listHelp = func(args ...string) string {
	return fmt.Sprintf(`Usage:
%s %s [min] [max]

Options:
    min	Minimum number of pages to list
    max	Maximum number of pages to list
Example:
    listpages 1 10`, args[0], args[1])
}

func (c *ListPages) Run(args ...interface{}) interface{} {
	var requiredArgs = 2
	ansi.PrintDebug("listpages function called!")

	subArray := args[0].([]string) // Get the arguments from the interface, as it's passing a slice of strings: [["args"]]
	subArray = subArray[1:]        // Remove the command name from the arguments

	if len(subArray) < requiredArgs {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of %d", len(args), requiredArgs))
		ansi.PrintWarning(listHelp(filepath.Base(os.Args[0]), c.Name()))
		return nil
	}

	if len(subArray) > requiredArgs {
		for _, arg := range subArray {
			ansi.PrintDebug("Arg: " + arg)
			if needHelp(arg) {
				return c.Help()
			}
		}
		ansi.PrintError("Too many arguments: " + strconv.Itoa(len(subArray)) + " instead of " + strconv.Itoa(requiredArgs))
		ansi.PrintWarning(listHelp(filepath.Base(os.Args[0]), c.Name()))
		return nil
	}

	min, _ := strconv.Atoi(subArray[0])
	max, _ := strconv.Atoi(subArray[1])

	if max < min {
		ansi.PrintError("Max must be greater than min")
		ansi.PrintWarning(listHelp(filepath.Base(os.Args[0]), c.Name()))
		return nil
	}

	return listPages(min, max)
}

func (c *CreatePage) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("createpage function called!")
	fmt.Println("Args: ", args)
	if len(args) < 1 {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of 1", len(args)))
		return nil
	}

	if len(args) > 1 {
		for _, arg := range args {
			if needHelp(arg.(string)) {
				return c.Help()
			}
		}
	}

	path, _ := args[0].(string)

	return createPage(path)
}

func (c *EditPage) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("editpage function called!")
	if len(args) < 1 {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of 1", len(args)))
		return nil
	}

	id, _ := args[0].(string)

	return editPage(id)
}

func (c *DeletePage) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("deletepage function called!")
	if len(args) < 1 {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of 1", len(args)))
		return nil
	}

	id, _ := args[0].(string)

	return deletePage(id)
}

func (c *PublishPage) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("publishpage function called!")
	if len(args) < 1 {
		ansi.PrintError(c.HelpStr)
		return nil
	}

	id := fmt.Sprintf("%s", args[0]) // args[0].(string)

	return publishPage(id)
}

func (c *UnpublishPage) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("unpublishpage function called!")
	if len(args) < 1 {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of 1", len(args)))
		return nil
	}

	id, _ := args[0].(string)

	return unpublishPage(id)
}

func (c *PageStatus) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("pagestatus function called!")
	if len(args) < 1 {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of 1", len(args)))
		return nil
	}

	id, _ := args[0].(string)

	return pageStatus(id)
}

func (c *PageTags) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("pagetags function called!")
	if len(args) < 1 {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of 1", len(args)))
		return nil
	}

	id, _ := args[0].(string)

	return pageTags(id)
}

func (c *Config) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("config function called!")
	if len(args) < 1 {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of 1", len(args)))
		return nil
	}

	key, _ := args[0].(string)

	return config(key)
}

func (c *ShowPage) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("showpage function called!")
	if len(args) < 1 {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of 1", len(args)))
		return nil
	}

	id, _ := args[0].(string)

	return showPage(id)
}

func (c *ParseAll) Run(args ...interface{}) interface{} {
	var force bool
	force = false

	for _, arg := range args[0].([]string)[1:] { // Remove the command name from the arguments
		if arg == "rebuild" {
			force = confirmForce()
		}
	}
	return parseAll(force)
}

func (c *Nop) Run(args ...interface{}) interface{} {

	// Spell checking algorithm
	// 1. Get the word
	// 2. Check if the word is in the dictionary (commands list)
	// 3. If the word is in the dictionary, return the word
	// 4. If not, calculate the Levenshtein distance between the word and all the words in the dictionary
	// 5. Suggest the word with the smallest distance to the user
	return noop()
}
