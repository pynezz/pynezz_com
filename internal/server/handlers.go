package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func aboutHandler(c echo.Context) error {
	return c.String(http.StatusOK, "About page\n")
}

func postsHandler(c echo.Context) error {
	id := c.Param("id")
	response := fmt.Sprintf("Post ID: %s\n", id)
	return c.String(http.StatusOK, response)
}
