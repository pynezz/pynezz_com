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

func handleLogin(c echo.Context) error {
	return c.String(http.StatusOK, "Login page\n")
}

func handleRegister(c echo.Context) error {
	// css
	c.Get("/styles/templ.css")

	return Render(c, http.StatusOK, templates.Register())
}
