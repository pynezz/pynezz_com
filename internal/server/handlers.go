package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/templates"
)

func homeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, templates.Home())
}

func aboutHandler(c echo.Context) error {
	return c.String(http.StatusOK, "About page\n")
}
