package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pynezz/pynezz_com/cmd"
	"github.com/pynezz/pynezz_com/cmd/cms"
	"github.com/pynezz/pynezz_com/cmd/serve"
	"github.com/pynezz/pynezzentials/ansi"
)

var warning = func(warning string) {
	r, g, b, _ := ansi.HexToRGB("#e64553") // catppuccin latte maroon
	ansi.PrintBold(ansi.HexColor256(r, g, b, warning+"\n"))
}

// ridiculously over-engineered warning message for not implemented features
var warnNotImplemented = func(msg string) {
	r, g, b, _ := ansi.HexToRGB("#e64553")    // catppuccin latte maroon
	wr, wg, wb, _ := ansi.HexToRGB("#f9e2af") // catppuccin latte peach
	l := len(msg)
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			fmt.Print(ansi.HexColor256(r, g, b, ansi.RoundedHoriz))
		} else {
			fmt.Print(ansi.HexColor256(wr, wg, wb, ansi.RoundedHoriz))
		}
	}
	fmt.Printf("\n%s\n", msg)
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			fmt.Print(ansi.HexColor256(r, g, b, ansi.RoundedHoriz))
		} else {
			fmt.Print(ansi.HexColor256(wr, wg, wb, ansi.RoundedHoriz))
		}
	}
	fmt.Println()
}

////go:embed templates/*
// var resources embed.FS

var buildVersion string

// var t = template.Must(template.ParseFS(resources, "templates/*", "templates/layout/*"))

var header = func() string {
	return fmt.Sprintf("%s%s", ansi.FormatRoundedBox("pynezz.dev CLI\n"+buildVersion), "\n")
}

func main() {

	fmt.Println(header())
	args := os.Args[1:]
	Execute(args...)
}

var displayHelp func(args ...string) = func(args ...string) {
	fmt.Printf(`Usage:
	%s [module] [options]
Options:
  cms     Execute the CMS module
  serve   Serve the webapp

  info    Print information about the CLI
  help    Print this help message
`, filepath.Base(os.Args[0]))
}

var info func(...string) = func(...string) {
	if buildVersion == "" {
		buildVersion = "development"
	}
	msg := fmt.Sprintf(`%s is a command line interface for the pynezz.dev website, written in Go.
It's a simple markdown based CMS to manage the content of my website and serve the webapp.`+
		"\n\n", ansi.ColorF(ansi.Cyan, "%s", strings.Split(filepath.Base(os.Args[0]), "_")[0]))

	msg += fmt.Sprintf("version:     %s\n", buildVersion)
	msg += fmt.Sprintf("author:      %s", ansi.ColorF(ansi.Cyan, "Kevin aka. pynezz\n"))
	msg += fmt.Sprintf("website:     %s", ansi.ColorF(ansi.Cyan, "https://pynezz.dev\n"))
	msg += fmt.Sprintf("source code: %s", ansi.ColorF(ansi.Cyan, "%s", "https://github.com/pynezz/pynezz_com\n"))

	fmt.Println(msg)
}

func Execute(args ...string) {

	if len(args) < 1 {
		displayHelp(args...)
		return
	}

	needHelp := false
	for _, arg := range args {
		// fmt.Println("arg: ", arg)
		if arg == "help" {
			if help := cmd.Execute(args...); help != "" {
				needHelp = true
				fmt.Println(help)
				return
			}
		}
	}

	if needHelp {
		fmt.Println("No help available for this module.")
		return
	}
	f := map[string]func(...string){
		"cms":   cms.Execute,
		"serve": serve.Execute,
		"info":  info,
	}

	// check the arguments and execute the function if it exists
	for _, module := range args[:1] {
		if f[module] == nil {
			warning("[!] unknown module: " + module)
			displayHelp(args...)
			return
		}
		f[module](args[1:]...)
	}
}
