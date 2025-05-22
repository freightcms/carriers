package web

import (
	"github.com/freightcms/carriers/db"
	"github.com/labstack/echo/v4"
)

type (
	AppContext struct {
		echo.Context
		db.DbContext
	}
)
