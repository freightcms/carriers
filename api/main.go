package api

import (
	"net/http"

	"github.com/freightcms/carriers/schemas"
	"github.com/freightcms/carriers/services"
	"github.com/labstack/echo/v4"
)

// ServiceMiddleware injects the carrier service into the context
func ServiceMiddleware(service services.CarrierService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("carrierService", service)
			return next(c)
		}
	}
}

func health(c echo.Context) error {
	return c.String(200, "OK")
}

func create(c echo.Context) error {
	service, ok := c.Get("carrierService").(services.CarrierService)
	if !ok {
		c.Logger().Debug("Could not get carrier service")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	schema := new(schemas.CreateCarrierSchema)
	c.Bind(&schema)
	model, err := service.CreateCarrier(schema)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, &model)
}

func getAllCarier(c echo.Context) error {
	service, ok := c.Get("carrierService").(services.CarrierService)
	if !ok {
		c.Logger().Debug("Could not get carrier service")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	carriers, err := service.GetCarriers()

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, &carriers)
}

func CreateApp(service services.CarrierService) *echo.Echo {
	e := echo.New()

	e.Use(ServiceMiddleware(service))

	// Routes
	e.GET("/", getAllCarier, ServiceMiddleware(service))
	e.GET("/healthcheck", health)
	e.POST("/", create, ServiceMiddleware(service))

	return e
}
