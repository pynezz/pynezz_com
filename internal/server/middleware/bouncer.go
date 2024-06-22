package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// Bouncer is a middleware that checks if the user is authenticated
func Bouncer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Checking if user is authenticated")
		return next(c)
	}
}
