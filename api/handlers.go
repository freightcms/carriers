package api

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	"github.com/gin-gonic/gin"
)

var (
	pageRegex     = regexp.MustCompile("page=[0-9]+")
	pageSizeRegex = regexp.MustCompile("pageSize=[0-9]+")
)

// GetPaginatedLink will return the current url passed in query parameters `page=` and `pageSize=`
// with the parameters passed in.
func GetPaginatedLink(currentUrl string, page, pageSize int) string {
	nextLink := pageRegex.ReplaceAllString(currentUrl, fmt.Sprintf("page=%d", page))
	nextLink = pageSizeRegex.ReplaceAllString(nextLink, fmt.Sprintf("pageSize=%d", pageSize))
	return nextLink
}

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
		PageSize: pageSize,
		Pages:    len(carriers) / pageSize,
		Entities: &carriers,
		Next:     "",
		Previous: "",
	}
	if len(carriers) >= pageSize {
		// we're going to assume there are more rseults
		resp.Next = GetPaginatedLink(c.Request.RequestURI, query.Page+1, pageSize)
	}
	if query.Page > 0 {
		resp.Previous = GetPaginatedLink(c.Request.RequestURI, query.Page-1, pageSize)
	}
	c.JSON(http.StatusOK, resp)
}

// GetCarrierHandler is a gin request handler for fetching a single Freight Carrier based on their identifier.
// The rout expects to have the carrier database interface attached to the gin context.
func GetCarrierHandler(c *gin.Context) {
	db := c.MustGet("db").(db.CarrierDb)
	id := c.Param("id")
	if id == "" {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	carrier, err := db.GetCarrier(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, carrier)
}

// CreateCarrierHandler deserializes the POST request body as a FreightCarrierModel. The endpoint attempts to create
// a new instance of the carrier. If it fails the function is aborted with the error in the response body and a status
// code of 400. By default the response body does not contain the carrier created. it must be requested
// calling a GET endpoint later on. The status code 201 Created is returned on success.
func CreateCarrierHandler(c *gin.Context) {
	var carrier models.FreightCarrierModel

	if err := c.ShouldBindBodyWithJSON(&carrier); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	db := c.MustGet("db").(db.CarrierDb)

	if err := db.CreateCarrier(c.Request.Context(), &carrier); err != nil {
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
