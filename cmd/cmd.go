package cmd

import (
	"fmt"

	"github.com/pynezz/pynezz_com/cmd/help"
)

// var usage func(...string) = func(args ...string) {
// 	f := map[string]func(args ...string){
// 		"help": displayHelp,
// 		"info": func(...string) {
// 			fmt.Println("Pynezz.com CLI")
// 		},
// 		"cms": CMS,
// 	}

// 	for _, arg := range args[:2] {
// 		if f[arg] == nil {
// 			return
// 		}
// 		f[arg](args...)
// 	}

// 	fmt.Printf(`Usage:
// 	%s [module] [options]
// Options:
//    cms		Execute the CMS module
//    serve		Serve the webapp

//    info		Print information about the CLI
//    help		Print this help message
// `, filepath.Base(os.Args[0]))
// }

func Execute(args ...string) string {
	// fmt.Println("Hello from the main package!")
	fmt.Printf("Args: %v\n", args)
	if len(args) < 2 {
		return ""
	}
	return help.Help(args...)
}
