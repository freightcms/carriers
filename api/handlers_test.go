package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
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

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.MatchRegex(t, w.Body.String(), regexp.MustCompile("\"error\":"))
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

	// Assert
	assert.Equal(t, w.Code, http.StatusOK)
	assert.MatchRegex(t, w.Body.String(), regexp.MustCompile("\"next\"\\:(\\s+)?\"(.+)(page\\=\\d+).+\""))
}
