package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/server/middleware"
	"github.com/pynezz/pynezz_com/templates"
)

func Serve(port string) {
	fmt.Println("Serving the webapp on port", port)
	app := echo.New()
	setup(app)
	c1 := templates.Style()
	handler := templ.NewCSSMiddleware(app, c1)
	// app.GET("/styles/templ.css", func(c echo.Context) error {
	// 	css := templates.RenderStyle()
	// 	return c.Blob(http.StatusOK, "text/css", []byte(css))
	// })

	app.Logger.Fatal(http.ListenAndServe(":"+port, handler))
}

func setup(app *echo.Echo) {

	ctx := app.AcquireContext()
	defer app.ReleaseContext(ctx)

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Response().Header().Set(echo.HeaderServer, "pynezz.dev")
			return next(ctx)
		}
	})

	app.GET("/favicon.ico", func(c echo.Context) error {
		return c.File("pynezz/public/favicon.ico")
	})

	app.GET("/static/svgs/github-icon.svg", func(c echo.Context) error {
		return c.File("pynezz/public/svgs/github-icon.svg")
	})

	app.GET("/", homeHandler)

	// Use Bouncer middleware
	app.GET("/login", middleware.Bouncer(handleLogin))
	app.GET("/register", handleRegister)
	app.POST("/register", middleware.Register(handleLogin))
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
