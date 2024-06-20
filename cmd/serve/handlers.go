package serve

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func aboutHandler(c echo.Context) error {
	return c.String(http.StatusOK, "About page")
}

func postsHandler(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Post ID: "+id)
}
