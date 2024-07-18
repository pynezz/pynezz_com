package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pynezz/pynezz_com/internal/server/middleware"
	"github.com/pynezz/pynezz_com/templates"
)

func homeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, templates.Home())
}

func aboutHandler(c echo.Context) error {
	// get local storage for user

	return Render(c, http.StatusOK, templates.About())
}

func handleLogin(c echo.Context) error {
	cookie, _ := c.Cookie("Authorization")
	if cookie != nil {
		valid, _ := middleware.VerifyJWTToken(cookie.Value)
		if valid.Valid {
			return c.Redirect(http.StatusMovedPermanently, "/dashboard")
		}
	}

	return Render(c, http.StatusOK, templates.Login())
}

func handleRegister(c echo.Context) error {
	return Render(c, http.StatusOK, templates.Register())
}

func gotoDashboard(c echo.Context) error {
	cc := c.(*middleware.CustomContext)

	if cc.User.Username == "" {
		return Render(c, http.StatusUnauthorized, templates.DashboardError())
	}
	return Render(c, http.StatusOK, templates.Dashboard(cc.User))
}

// func gotoDashboard(c echo.Context, user models.User) error {
// 	c.Redirect(http.StatusMovedPermanently, "/dashboard")

// 	return Render(c, http.StatusOK, templates.Dashboard(user))
// }
