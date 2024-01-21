package api

import (
	"net/http"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/schemas"
	"github.com/freightcms/carriers/services"
	"github.com/labstack/echo/v4"
)

// ServiceMiddleware injects the carrier service into the context
func ServiceMiddleware(db db.CarrierDb) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("carrierService", services.NewCarrierService(c.Request().Context(), db))
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

	var schema schemas.CreateCarrierSchema
	c.Bind(&schema)
	model, err := service.CreateCarrier(&schema)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
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
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &carriers)
}

func delete(c echo.Context) error {
	service, ok := c.Get("carrierService").(services.CarrierService)
	if !ok {
		c.Logger().Debug("Could not get carrier service")
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	id := c.Param("id")
	err := service.DeleteCarrier(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

// Creates a new echo app with all routes defined.
// The carrierRepository is injected into the app. The carrierRepository is
// used to create a new carrier service which is injected into the context
// of each request. This allows the handlers to access the carrier service
// without having to create a new one for each request.
// The carrier service is created in the services package.
func CreateApp(carrierRepository db.CarrierDb) *echo.Echo {
	e := echo.New()

	e.Use(ServiceMiddleware(carrierRepository))

	// Routes
	e.GET("/", getAllCarier, ServiceMiddleware(carrierRepository))
	e.GET("/healthcheck", health)
	e.POST("/", create, ServiceMiddleware(carrierRepository))
	e.DELETE("/:id", delete, ServiceMiddleware(carrierRepository))

	return e
}
