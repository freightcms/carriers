package web

import (
	"net/http"

	"github.com/freightcms/carriers/db"
	"github.com/labstack/echo/v4"
)

var (
	getAllCarriersHandler = echo.HandlerFunc(func(c echo.Context) error {
		mgr := c.(AppContext).CarrierResourceManager
		var reqQuery GetCarriersRequest
		if err := c.Bind(&reqQuery); err != nil {
			return err
		}
		dbQuery := db.NewQuery().SetPage(reqQuery.Page).SetPageSize(reqQuery.Limit)

		carriers, count, err := mgr.Get(dbQuery)
		if err != nil {
			return err
		}

		resp := GetCarriersResponse{
			Total:    count,
			Carriers: carriers,
		}

		return c.JSON(http.StatusOK, resp)
	})
)
