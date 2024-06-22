package server

import (
	"fmt"

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
	app.GET("/posts/:id", postsHandler)
}
