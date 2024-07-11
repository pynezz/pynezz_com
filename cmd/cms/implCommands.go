package cms

import (
	"fmt"

	"github.com/pynezz/pynezzentials/ansi"
)

func (c *ListPages) Run(args ...interface{}) interface{} {
	var requiredArgs = 2
	ansi.PrintDebug("listpages function called!")

	if len(args) < requiredArgs {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of %d", len(args), requiredArgs))
		return nil
	}

	min, _ := args[0].(int)
	max, _ := args[1].(int)

	return listPages(min, max)
}

func (c *CreatePage) Run(args ...interface{}) interface{} {
	ansi.PrintDebug("createpage function called!")

	if len(args) < 1 {
		ansi.PrintError(fmt.Sprintf("Not enough arguments: %d instead of 1", len(args)))
		return nil
	}

	path, _ := args[0].(string)

	return createPage(path)
}

func (*CreatePage) Help() string {
	return "Create a page"
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
