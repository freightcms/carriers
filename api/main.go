package api

import (
	"net/http"

	"github.com/freightcms/carriers/schemas"
	"github.com/labstack/echo/v4"
)

// CarrierService is the interface that provides carrier methods.
type CarrierService interface {
	// CreateCarrier Creates a new carrier and returns the created carrier. If the carrier
	// could not be created, an error is returned.
	CreateCarrier(schema *schemas.CreateCarrierSchema) (*schemas.CarrierSchema, error)
	// GetCarrier returns a carrier by id.
	GetCarrier(id string) (*schemas.CarrierSchema, error)
	// GetCarriers returns all carriers.
	GetCarriers() ([]*schemas.CarrierSchema, error)
	// UpdateCarrier updates a carrier.
	UpdateCarrier(schema *schemas.CarrierSchema) (*schemas.CarrierSchema, error)
	// DeleteCarrier deletes a carrier.
	DeleteCarrier(id string) error
}

// ServiceMiddleware injects the carrier service into the context
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

	schema := new(schemas.CreateCarrierSchema)
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
