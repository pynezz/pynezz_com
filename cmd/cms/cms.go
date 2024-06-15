package cms

import "fmt"

type ICMS interface {
	Help(args ...string) string
	CMS(args ...string) cms
}


type cms struct{}

func commands() map[string]string {
	prefix := "--"

	return map[string]Command{
		prefix + "list": "List all the pages",
		prefix + "page":
		"--create":      "Create a new page",
		"--edit":        "Edit an existing page",
		"--delete":      "Delete a page",
		"--publish":     "Publish a page",
		"--unpublish":   "Unpublish a page",
		"--status":      "Get the status of a page",
		"--tags":        "Get the tags of a page",

		"--stats":  "Get the statistics of the CMS",
		"--config": "Show the configuration",
	}
}

func listCommands() string {
	var s string
	for k, v := range commands() {
		s += k + "\t" + v + "\n"
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

func run(c Command) {
	fmt.Println("running command: ", c)
}

func Execute(args ...string) {
	fmt.Printf("Hello from the CMS module!\n")

	if len(args) < 1 {
		fmt.Println(listCommands())
		return
	}

	for k, v := range commands() {
		if arg == k {
			fmt.Println(v)
			return
		}
	}

}
