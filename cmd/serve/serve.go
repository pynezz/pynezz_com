package serve

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
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

func Serve(args ...string) {

	app := echo.New()
	setup(app)
}

func setup(app *echo.Echo) {
	app.GET("/", homeHandler)
	app.GET("/about", aboutHandler)
	app.GET("posts/:id", postsHandler)

}

func Help(args ...string) string {
	return fmt.Sprintf("Help for serve module: \n%s", usage(args...))
}

func Execute(args ...string) {
	fmt.Println("Hello from the serve package!")

	var port string

	// Some args parsing
	if len(args) < 1 {
		fmt.Println("Please provide a command.")
		usage(args...)
		return
	}

	for i, arg := range args[:1] {
		if arg == "--help" {
			usage(args...)
			return
		}

		if arg == "--port" || arg == "-p" {
			if len(args) < 2 {
				fmt.Println("Please provide a port number.")
				usage(args...)
				return
			}

			port := args[i+1]

			fmt.Println("Listening on port", port)
			return
		}
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
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
