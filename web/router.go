package web

import (
	"github.com/labstack/echo/v4"
)

var (
	router *echo.Router
)

func Register(e *echo.Echo) {
	e.GET("/carriers", getAllCarriersHandler)
}
