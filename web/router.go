package web

import "github.com/labstack/echo/v4"

var (
	router *echo.Router
)

func Router(e *echo.Echo) *echo.Router {
	if router != nil {
		return router
	}
	router = echo.NewRouter(e)

	router.Add(echo.POST, "/carreirs", getAllCarriersHandler)

	return router
}
