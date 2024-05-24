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
}

func createMockCarrierDb() db.CarrierDb {
	return &MockCarrierDb{}
}

func (m *MockCarrierDb) GetCarriers(ctx context.Context) ([]models.FreightCarrierModel, error) {
	return nil, nil
}

func (m *MockCarrierDb) GetCarrier(ctx context.Context, id string) (*models.FreightCarrierModel, error) {
	return nil, nil
}

func (m *MockCarrierDb) CreateCarrier(ctx context.Context, carrier *models.FreightCarrierModel) error {
	return nil
}

func (m *MockCarrierDb) UpdateCarrier(ctx context.Context, id string, carrier *models.FreightCarrierModel) error {
	return nil
}

func (m *MockCarrierDb) DeleteCarrier(ctx context.Context, id string) error {
	return nil
}

func TestGetCarriersHandler_Should_Return_BindingError_With_400_StatusCode(t *testing.T) {
	// Arrange
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		// fake mock carrier db middleware
		c.Set("db", createMockCarrierDb())
	})

	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/carriers?page=a", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.MatchRegex(t, w.Body.String(), regexp.MustCompile("\"error\":"))
}
