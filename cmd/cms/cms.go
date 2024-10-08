package cms

import (
	"fmt"

	"github.com/pynezz/pynezzentials/ansi"
)

type ICMS interface {
	ICommand // Embedding the ICommand interface to inherit its methods

	CMS(args ...string) cms
}

type cms struct{}

// INFO: This is example commands for now.
func commands() map[string]ICommand {
	prefix := "--"

	// 	typo := regexp.MustCompile(`(.*?)`)

	// fmt.Println("Creating commands...")
	return map[string]ICommand{
		prefix + "list": c["list"],
		// prefix + "build":     c["build"], // Build markdown files to HTML // Might be better suited on the publish command
		prefix + "page":      c["page"],
		prefix + "create":    c["create"],
		prefix + "edit":      c["edit"],
		prefix + "delete":    c["delete"],
		prefix + "publish":   c["publish"], // Publish a page / convert markdown to HTML (which in turns creates a new HTML file in the public directory, updates the database, and the site)
		prefix + "unpublish": c["unpublish"],
		prefix + "status":    c["status"],
		prefix + "tags":      c["tags"],
		prefix + "config":    c["config"],
		prefix + "parse":     c["parse"], // Parse all markdown files in the content directory
		prefix + "build":     c["build"], // Alias for parse
	}
}

func listCommands() string {
	var s string

	cmds := commands()
	pad := longestStringMap(cmds)

	for k, v := range cmds {
		if v != nil {
			s += "  " + k + ansi.AddPadding(" ", pad-len(k)) + v.Help() + "\n"
		} else {
			s += k + "\t" + "⚠️ error:  command not initialized!\n"
		}
	}
	return s
}

func isValidCommand(cmd string) bool {
	for _, v := range validCommands {
		if cmd == v {
			return true
		}
	}
	return false
}

func Help(args ...string) string {
	return listCommands()
}

func (c *cms) CMS(args ...string) *cms {
	return &cms{}
}

func run(c ICommand, args ...string) {
	fmt.Println("running command: ", c.Name())

	c.Run(args)
}

func Execute(args ...string) {
	fmt.Printf("Hello from the CMS module!\n")
	for i, arg := range args {
		fmt.Printf("[%d]:%s\n", i, arg)
		if arg == "help" && len(args) > 1 {
			Help(args...)
			return
		}
	}

	if len(args) < 1 {
		fmt.Println(listCommands())
		return
	}

	if ok := testDbConnection(); !ok {
		ansi.PrintError("❌ Database connection failed.")
		return
	}

	// If the command exists, run it
	for k, v := range commands() {
		for _, arg := range args {
			if arg == k {
				if v != nil {
					run(v, args...)
				} else {
					fmt.Printf("Command %s is not initialized.\n", k)
					return
				}
				return
			}
		}
	}

	// If no valid command is found, check for typos
	for _, arg := range args {
		if !isValidCommand(arg) {
			fmt.Printf("Invalid command: %s\n", arg)
			fmt.Println("Did you mean one of these?")
			fmt.Println(listCommands())
			return
		}
	}
}

// Get the longest string in a slice of strings and return its length
func longestString(s []string) int {
	var max int
	for _, v := range s {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

func longestStringMap(m map[string]ICommand) int {
	var max int
	for k := range m {
		if len(k) > max {
			max = len(k)
		}
	}
	return max
}
