package api

import (
	"net/http"

	"github.com/freightcms/carriers/models"
	"github.com/labstack/echo/v4"
)

type CarrierService interface {
	CreateCarrier(schema *CreateCarrierSchema) (models.FreightCarrier, error)
}

func ServiceMiddleware(service *CarrierService) echo.MiddlewareFunc {
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
	service, ok := c.Get("carrierService").(CarrierService)
	if !ok {
		c.Logger().Debug("Could not get carrier service")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	schema := new(CreateCarrierSchema)
	c.Bind(schema)
	model, err := service.CreateCarrier(schema)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, &model)
}

func CreateApp(service *CarrierService) *echo.Echo {
	e := echo.New()

	// Routes
	e.GET("/", health)
	e.POST("/carriers", create, ServiceMiddleware(service))

	return e
}
