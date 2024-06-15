package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pynezz/pynezz_com/cmd"
	"github.com/pynezz/pynezz_com/cmd/cms"
	"github.com/pynezz/pynezz_com/cmd/serve"
)

func main() {
	args := os.Args[1:]
	Execute(args...)
}

var displayHelp func(...string) = func(args ...string) {
	fmt.Printf(`Usage:
	%s [module] [options]
Options:
  cms		Execute the CMS module
  serve		Serve the webapp

  info		Print information about the CLI
  help		Print this help message
`, filepath.Base(os.Args[0]))
}

// var usage func(...string) = func(args ...string) {
// 	f := map[string]func(args ...string){
// 		"help": displayHelp,
// 		"info": func(...string) {
// 			fmt.Println("Pynezz.com CLI")
// 		},
// 		"cms": cmd.CMS(args[:1]),
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
//   cms		Execute the CMS module
//   serve		Serve the webapp

//   info		Print information about the CLI
//   help		Print this help message
// `, filepath.Base(os.Args[0]))
// }

var info func(...string) = func(...string) {
	fmt.Println("Pynezz.com CLI")
}

func Execute(args ...string) {

	if len(args) < 1 {
		displayHelp(args...)
		return
	}

	for _, arg := range args[:1] {
		if arg == "help" {
			if help := cmd.Execute(args...); help != "" {
				fmt.Println(help)
				return
			}
		}
	}

	f := map[string]func(...string){
		"cms":   cms.Execute,
		"serve": serve.Execute,
		"info":  info,
	}

	// check the arguments and execute the function if it exists
	for _, module := range args[:1] {
		if f[module] == nil {
			fmt.Printf("Unknown module %s\n", module)
			displayHelp(args...)
			return
		}
		f[module](args[1:]...)
	}
}
