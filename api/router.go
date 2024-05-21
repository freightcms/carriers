package api

import (
	"github.com/gin-gonic/gin"
)

// RegisterHandlers appends the carrier routes to the gin router passed
// as the first argument. Creates a router group with the prefix "/carriers"
// and registers the following routes:
//
// - GET /carriers
//
// - GET /carriers/:id
//
// - POST /carriers
//
// - PUT /carriers/:id
//
// - DELETE /carriers/:id
//
// Returns the router group created.
func CreateRouterGroup(router gin.IRouter) *gin.RouterGroup {
	r := router.Group("/carriers")
	r.GET("/", GetCarriersHandler)
	r.GET("/:id", GetCarrierHandler)
	r.POST("/", CreateCarrierHandler)
	r.PUT("/:id", UpdateCarrierHandler)
	r.DELETE("/:id", DeleteCarrierHandler)

	return r
}
