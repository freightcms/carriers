package api

import (
	"net/http"

	"github.com/freightcms/carriers/schemas"
	"github.com/freightcms/carriers/services"
	"github.com/labstack/echo/v4"
)

// CarrierService is the interface that provides carrier methods.
type Service interface {
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
func ServiceMiddleware(service Service) echo.MiddlewareFunc {
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
	c.Bind(schema)
	model, err := service.CreateCarrier(schema)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, &model)
}

func getCarriers(c echo.Context) error {
	service, ok := c.Get("carrierService").(services.CarrierService)
	if !ok {
		c.Logger().Debug("Could not get carrier service")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	models, err := service.GetCarriers()

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, &models)
}

func getCarrier(c echo.Context) error {
	service, ok := c.Get("carrierService").(services.CarrierService)
	if !ok {
		c.Logger().Debug("Could not get carrier service")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	id := c.Param("id")
	model, err := service.GetCarrier(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, &model)
}

func deleteCarrier(c echo.Context) error {
	service, ok := c.Get("carrierService").(services.CarrierService)
	if !ok {
		c.Logger().Debug("Could not get carrier service")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	id := c.Param("id")
	err := service.DeleteCarrier(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.NoContent(http.StatusOK)
}

func updateCarrier(c echo.Context) error {
	service, ok := c.Get("carrierService").(services.CarrierService)
	if !ok {
		c.Logger().Debug("Could not get carrier service")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	id := c.Param("id")
	schema := new(schemas.CarrierSchema)
	c.Bind(schema)
	schema.ID = id

	model, err := service.UpdateCarrier(schema)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, &model)
}

// CreateApp creates a new echo app for the Carrier API to run on.
// Service sould be the carrier service used to interact between any external web service(s) or data stores.
func CreateApp(service Service) *echo.Echo {
	e := echo.New()

	// Routes
	e.GET("/", health)
	e.GET("/carriers", getCarriers, ServiceMiddleware(service))          // paginate through carriers
	e.GET("/carriers/:id", getCarrier, ServiceMiddleware(service))       // get a specific carrier by id
	e.DELETE("/carriers/:id", deleteCarrier, ServiceMiddleware(service)) // delete a specific carrier by id
	e.PUT("/carriers/:id", updateCarrier, ServiceMiddleware(service))    // update a specific carrier by id
	e.POST("/carriers", create, ServiceMiddleware(service))              // create a new carrier

	return e
}
