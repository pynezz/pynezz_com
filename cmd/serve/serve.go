package serve

import (
	"fmt"

	"github.com/pynezz/pynezz_com/internal/helpers"
	"github.com/pynezz/pynezz_com/internal/server"
	"github.com/pynezz/pynezzentials/ansi"
)

var usage func(...string) string = func(args ...string) string {
	fmt.Println("args in serve.go", args)

	return fmt.Sprintln(`Usage: serve [options]

  Serve the web server on a specified port.

Options:
    --help      Print this help message
    --port, -p  Specify the port to listen on

Example:
  serve --port 8080

Visit http://localhost:8080 in your browser to see the webapp.`)
}

func Help(args ...string) string {
	return fmt.Sprintf("Help for serve module: \n%s", usage(args...))
}

func Execute(args ...string) {

	fmt.Println("Hello from the serve package!")

	// Some args parsing
	if len(args) < 1 {
		helpers.Warning("Please provide a command.")
		fmt.Println(usage(args...))
		return
	}

	for i, arg := range args[:1] {
		if arg == "--help" {
			fmt.Println(usage(args...))
			return
		}

		if arg == "--port" || arg == "-p" {
			if len(args) < 2 {
				fmt.Println("Please provide a port number.")
				fmt.Println(usage(args...))
				return
			}

			port := args[i+1]

			server.Serve(port)

			ansi.PrintInfo("Listening on port: " + port)
		}
	}

	ansi.PrintInfo("Waiting for SIGINT (Ctrl+C) to shutdown...")
}
