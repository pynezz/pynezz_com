package server

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/server/middleware"
	"github.com/pynezz/pynezz_com/templates"
	"github.com/pynezz/pynezzentials/ansi"
)

func Serve(port string) {
	fmt.Println("Serving the webapp on port", port)
	app := echo.New()

	setup(app)
	app.Start(":" + port)
}

func setup(app *echo.Echo) {
	ctx := app.AcquireContext()
	defer app.ReleaseContext(ctx)

	// set server header
	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Response().Header().Set(echo.HeaderServer, "pynezz.dev")
			return next(ctx)
		}
	})

	// vips := make(map[string]bool)

	// set security headers
	app.Use(middleware.SecurityHeaders)
	// app.Use(middleware.Sec)
		
	// static
	app.Static("/static", "pynezz/public/")
	app.Static("/css", "pynezz/public/css/")
	app.Static("/", "pynezz/public/")
	app.GET("/static/svgs/github-icon.svg", func(c echo.Context) error {
		return c.File("pynezz/public/svgs/github-icon.svg")
	})

	postsHandler := newPostsHandler()

	app.GET("/", homeHandler)

	// Use Bouncer middleware
	app.GET("/login", middleware.Bouncer(handleLogin))
	app.POST("/login", middleware.Login(middleware.Bouncer(gotoDashboard)))

	app.GET("/register", handleRegister)
	app.POST("/register", middleware.Register(middleware.Bouncer(gotoDashboard)))

	app.GET("/dashboard", middleware.Bouncer(gotoDashboard))
	app.GET("/about", aboutHandler)

	app.GET("/posts", postsHandler.handleShowLastPosts)
	app.GET("/posts/", postsHandler.handleShowLastPosts)
	app.GET("/posts/:slug", postsHandler.GetPostBySlug)

	// BUG! This does not fetch the posts
	app.GET("/tags/:tag", postsHandler.GetPostsByTag)

	// BUG! This does not fetch the tags (only one)
	app.GET("/tags/", postsHandler.GetTags)
	app.GET("/tags", postsHandler.GetTags)

	// Todo - consider doing this, or just managing it via the CLI in the backend
	// app.GET("/posts/:slug/edit", postsHandler.EditPost)

	app.POST("csp-report", func(c echo.Context) error {
		cpsReport(c.FormValue("csp-report"))
		return c.String(200, "CSP Report Received")
	})
}

func cpsReport(report string) {
	ansi.PrintWarning(fmt.Sprintf("CSP Report: %v ", report))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	nonce, err := middleware.GenerateNonce()
	if err != nil {
		ansi.PrintError("Error generating nonce: " + nonce)
	}
	if err := templates.Root(t, nonce, ctx.Path()).Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
