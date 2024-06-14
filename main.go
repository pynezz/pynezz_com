package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pynezz/pynezz_com/cmd"
)

func main() {
	args := os.Args[1:]
	Execute(args...)
}

var usage func() = func() {
	fmt.Printf(`Usage:
	%s [...options]
Options:
	help       				Print this help message
	bivrost 				Execute bivrost
	test_module				Execute test_module module
	threat_detection 		Execute threat_detection module
Example:
	%s bivrost test_module	Will execute both bivrost and test_module`, filepath.Base(os.Args[0]), filepath.Base(os.Args[0]))
}

var info func() = func() {
	fmt.Println("Pynezz.com CLI")
}

func Execute(args ...string) {

	f := map[string]func(){
		"cms":   cmd.CMS,
		"serve": cmd.Serve,
		"help":  usage,
		"info":  info,
	}

	// check the arguments and execute the function if it exists
	for _, module := range args {
		if f[module] == nil {
			fmt.Printf("Unknown module %s\n", module)
			usage()
			return
		}
		f[module]()
	}
}
