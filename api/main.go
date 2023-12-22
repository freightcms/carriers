package api

import (
	"github.com/labstack/echo/v4"
)

func health(c echo.Context) error {
	return c.String(200, "OK")
}

func createApp() *echo.Echo {
	e := echo.New()

	// Routes
	e.GET("/", health)

	return e
}
