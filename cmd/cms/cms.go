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

	// fmt.Println("Creating commands...")
	return map[string]ICommand{
		prefix + "list":      c["list"],
		prefix + "page":      c["page"],
		prefix + "create":    c["create"],
		prefix + "edit":      c["edit"],
		prefix + "delete":    c["delete"],
		prefix + "publish":   c["publish"],
		prefix + "unpublish": c["unpublish"],
		prefix + "status":    c["status"],
		prefix + "tags":      c["tags"],
		prefix + "config":    c["config"],
	}
}

func listCommands() string {
	var s string

	// fmt.Println("Listing commands...")
	cmds := commands()

	pad := longestStringMap(cmds)
	// fmt.Println("Commands length: ", len(cmds))
	// fmt.Println("Padding: ", pad)
	for k, v := range cmds {
		if v != nil {
			s += "  " + k + ansi.AddPadding(" ", pad-len(k)) + v.Help() + "\n"
		} else {
			s += k + "\t" + "⚠️ error:  command not initialized!\n"
		}
	}
	// for k, v := range cmds {
	// 	s += k + "\t" + v.Help() + "\n"
	// }
	return s
}

func Help(args ...string) string {
	return listCommands()
}

// func CMS(args []string) string {
// 	return cms.Help(cms{}, args...)
// }

func (c *cms) CMS(args ...string) *cms {
	return &cms{}
}

func run(c ICommand, args ...string) {
	fmt.Println("running command: ", c.Name())
	c.Run(args)
}

func Execute(args ...string) {
	fmt.Printf("Hello from the CMS module!\n")

	if len(args) < 1 {
		fmt.Println(listCommands())
		return
	}

	// If the command exists, run it
	for k, v := range commands() {
		for _, arg := range args {
			if arg == k {
				if v != nil {
					// v.Run("hello")
					run(v, args...)
				} else {
					fmt.Printf("Command %s is not initialized.\n", k)
				}
				return
			}
			// if arg == k {
			// 	v.Run("hello")
			// 	return
			// }
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
