package server

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Serve(port string) {
	fmt.Println("Serving the webapp on port", port)
	app := echo.New()
	setup(app)

	app.Logger.Fatal(app.Start(":" + port))
}

func setup(app *echo.Echo) {
	app.GET("/", homeHandler)
	app.GET("/about", aboutHandler)
	app.GET("/posts/", newPostsHandler().handleShowLastPosts)
	app.GET("/posts/:id", newPostsHandler().handleShowPosts)
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
