package serve

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/helpers"
	"github.com/pynezz/pynezz_com/internal/server"
	"github.com/pynezz/pynezzentials/ansi"
)

var usage func(...string) string = func(args ...string) string {
	fmt.Println("args in serve.go", args)

	return fmt.Sprintln(`Usage: serve [options]
	Options:
		--help			Print this help message
		--port, -p	Specify the port to listen on

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

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
