package cms

import "fmt"

type ICMS interface {
	ICommand // Embedding the ICommand interface to inherit its methods

	CMS(args ...string) cms
}

type cms struct{}

// INFO: This is example commands for now.
func commands() map[string]ICommand {
	prefix := "--"

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
		prefix + "stats":     c["stats"],
		prefix + "config":    c["config"],
	}
}

func listCommands() string {
	var s string
	for k, v := range commands() {
		s += k + "\t" + v.Help() + "\n"
	}
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

func run(c ICommand) {
	fmt.Println("running command: ", c.Name())
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
				v.Run("hello")
				return
			}
		}
	}
}
