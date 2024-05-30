package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

type MockCarrierDb struct {
	create func(ctx context.Context, carrier *models.FreightCarrierModel) error
	delete func(ctx context.Context, id string) error
	get    func(ctx context.Context, id string) (*models.FreightCarrierModel, error)
	all    func(ctx context.Context) ([]models.FreightCarrierModel, error)
	update func(ctx context.Context, id string, carrier *models.FreightCarrierModel) error
}

// CreateCarrier implements db.CarrierDb.
func (m *MockCarrierDb) CreateCarrier(ctx context.Context, carrier *models.FreightCarrierModel) error {
	return m.create(ctx, carrier)
}

// DeleteCarrier implements db.CarrierDb.
func (m *MockCarrierDb) DeleteCarrier(ctx context.Context, id string) error {
	return m.delete(ctx, id)
}

// GetCarrier implements db.CarrierDb.
func (m *MockCarrierDb) GetCarrier(ctx context.Context, id string) (*models.FreightCarrierModel, error) {
	return m.get(ctx, id)
}

// GetCarriers implements db.CarrierDb.
func (m *MockCarrierDb) GetCarriers(ctx context.Context) ([]models.FreightCarrierModel, error) {
	return m.all(ctx)
}

// UpdateCarrier implements db.CarrierDb.
func (m *MockCarrierDb) UpdateCarrier(ctx context.Context, id string, carrier *models.FreightCarrierModel) error {
	return m.update(ctx, id, carrier)
}

func createMockDb() db.CarrierDb {
	return &MockCarrierDb{}
}

func TestGetCarriersHandler_Should_Return_BindingError_With_400_StatusCode(t *testing.T) {
	// Arrange
	router := gin.Default()
	mockDb := createMockDb()
	router.Use(func(ctx *gin.Context) {
		// fake mock carrier db middleware
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/carriers?page=a", nil)
	router.ServeHTTP(w, req)
	jsonBody := struct {
		Error string `json:"error"`
	}{}
	err := json.Unmarshal(w.Body.Bytes(), &jsonBody)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, jsonBody.Error)
}

func TestGetCarriersHandler_Should_Set_NextLink_In_Response(t *testing.T) {
	// Arrange
	router := gin.Default()
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 2), nil
		},
	}

	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()

	// TODO: the carrier db needs to send back at least the page size

	// Act
	uri := "http://localhost:3000/carriers?pageSize=2&page=2"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri
	router.ServeHTTP(w, req)
	jsonBody := &PaginatedQueryResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &jsonBody)

	// Assert
	assert.Equal(t, nil, err)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 2, jsonBody.Total)
	assert.Equal(t, 2, jsonBody.Page)
	assert.Equal(t, 2, jsonBody.PageSize)
	assert.Equal(t, jsonBody.Next, "http://localhost:3000/carriers?pageSize=2&page=3")
	assert.Equal(t, jsonBody.Previous, "http://localhost:3000/carriers?pageSize=2&page=1")
}

func TestGetCarriersHandler_Should_Not_Set_NextLink(t *testing.T) {
	// Arrange
	router := gin.Default()
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 1), nil
		},
	}

	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()

	// TODO: the carrier db needs to send back at least the page size

	// Act
	uri := "http://localhost:3000/carriers?pageSize=2&page=2" // because the page size is greater than the results returned we should not set the link
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri
	router.ServeHTTP(w, req)

	// Assert
	jsonBody := &PaginatedQueryResponse{}
	json.Unmarshal(w.Body.Bytes(), &jsonBody)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 1, jsonBody.Total)
	assert.Equal(t, 2, jsonBody.Page)
	assert.Equal(t, 2, jsonBody.PageSize)
	assert.Equal(t, "", jsonBody.Next) // should be no more links since there were less results than requested
	assert.Equal(t, "http://localhost:3000/carriers?pageSize=2&page=1", jsonBody.Previous)
}

func TestGetCarriersHandler_PreviousLink_Should_Be_Empty(t *testing.T) {
	// Arrange
	router := gin.Default()
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 2), nil
		},
	}

	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()

	// TODO: the carrier db needs to send back at least the page size

	// Act
	uri := "http://localhost:3000/carriers?pageSize=2&page=0" // because this is page 0 the previous link should not be set
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri
	router.ServeHTTP(w, req)

	// Assert
	jsonBody := &PaginatedQueryResponse{}
	json.Unmarshal(w.Body.Bytes(), &jsonBody)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 2, jsonBody.Total)
	assert.Equal(t, 0, jsonBody.Page)
	assert.Equal(t, 2, jsonBody.PageSize)
	assert.Equal(t, "http://localhost:3000/carriers?pageSize=2&page=1", jsonBody.Next)
	assert.Equal(t, "", jsonBody.Previous) // should be no more links since there were less results than requested
}

func TestGetCarriersHandler_PreviousLink_Should_Be_Not_Empty(t *testing.T) {
	// Arrange
	router := gin.Default()
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 2), nil
		},
	}

	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()

	// TODO: the carrier db needs to send back at least the page size

	// Act
	uri := "http://localhost:3000/carriers?pageSize=2&page=1" // because this is page 1 (greater than 0) the previous link should be set
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri
	router.ServeHTTP(w, req)

	// Assert
	jsonBody := &PaginatedQueryResponse{}
	json.Unmarshal(w.Body.Bytes(), &jsonBody)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 2, jsonBody.Total)
	assert.Equal(t, 1, jsonBody.Page)
	assert.Equal(t, 2, jsonBody.PageSize)
	assert.Equal(t, "http://localhost:3000/carriers?pageSize=2&page=2", jsonBody.Next)
	assert.Equal(t, "http://localhost:3000/carriers?pageSize=2&page=0", jsonBody.Previous) // should be no more links since there were less results than requested
}

func TestGetCarriersHandler_Should_Default_PageSize_When_Not_Provided(t *testing.T) {
	// Arrange
	router := gin.Default()
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 11), nil // need to create >10 (the default size) to ensure we get the next link back with the page size of 10
		},
	}

	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()

	// TODO: the carrier db needs to send back at least the page size

	// Act
	uri := "http://localhost:3000/carriers?&page=2" // don't pass in `pageSize` as the query parameter to check logic
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri
	router.ServeHTTP(w, req)

	// Assert
	jsonBody := &PaginatedQueryResponse{}
	json.Unmarshal(w.Body.Bytes(), &jsonBody)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 11, jsonBody.Total)
	assert.Equal(t, 2, jsonBody.Page)
	assert.Equal(t, 10, jsonBody.PageSize)
	assert.Equal(t, "http://localhost:3000/carriers?&page=3", jsonBody.Next)
	assert.Equal(t, "http://localhost:3000/carriers?&page=1", jsonBody.Previous) // should be no more links since there were less results than requested
}
