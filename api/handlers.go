package api

import (
	"net/http"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	"github.com/gin-gonic/gin"
)

// GetCarriersHandler is a handler function that retrieves all carriers from the database.
// The response is a JSON array of carrier objects.
//
// Query Params: `page`, `pageSize`, `include`
func GetCarriersHandler(c *gin.Context) {
	db := c.MustGet("db").(db.CarrierDb)
	var query PaginatedQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	carriers, err := db.GetCarriers(c.Request.Context())
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pageSize := query.PageSize
	if pageSize <= 0 { // avoid divide by zero issues
		pageSize = 10
	}
	resp := &PaginatedQueryResponse{
		Total:    len(carriers),
		Page:     query.Page,
		PageSize: query.PageSize,
		Pages:    len(carriers) / query.PageSize,
		Entities: &carriers,
	}
	c.JSON(http.StatusOK, &resp)
}

func GetCarrierHandler(c *gin.Context) {
	db := c.MustGet("db").(db.CarrierDb)
	id := c.Query("id")
	carrier, err := db.GetCarrier(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, carrier)
}

func CreateCarrierHandler(c *gin.Context) {
	var carrier models.FreightCarrierModel

	if err := c.ShouldBindBodyWithJSON(&carrier); err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	db := c.MustGet("db").(db.CarrierDb)

	if err := db.CreateCarrier(c.Request.Context(), &carrier); err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusCreated)
}

func UpdateCarrierHandler(c *gin.Context) {
	var carrier models.FreightCarrierModel
	id := c.Param("id")
	if err := c.ShouldBindBodyWithJSON(&carrier); err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	db := c.Value("db").(db.CarrierDb)
	if err := db.UpdateCarrier(c.Request.Context(), id, &carrier); err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func DeleteCarrierHandler(c *gin.Context) {
	id := c.Param("id")
	db := c.Value("db").(db.CarrierDb)
	if err := db.DeleteCarrier(c.Request.Context(), id); err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}
