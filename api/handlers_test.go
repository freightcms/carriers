package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestGetCarriersHandler_Should_Return_BindingError(t *testing.T) {
	// Arrange
	router := gin.Default()
	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/carriers", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
