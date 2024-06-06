package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/freightcms/carriers/db"
	"github.com/freightcms/carriers/models"
	"github.com/freightcms/carriers/validators"
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

func Test_GetCarriersHandler_Should_Return_BindingError_With_400_StatusCode(t *testing.T) {
	// arrange
	mockDb := createMockDb()
	w := httptest.NewRecorder()
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		// fake mock carrier db middleware
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/carriers?page=a", nil)

	// act
	router.ServeHTTP(w, req)

	// assert
	jsonBody := struct {
		Error string `json:"error"`
	}{}
	json.NewDecoder(w.Body).Decode(&jsonBody)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.NotEqual(t, nil, jsonBody.Error)
}

func Test_GetCarriersHandler_Should_Set_NextLink_In_Response(t *testing.T) {
	// arrange
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 2), nil
		},
	}
	writer := httptest.NewRecorder()
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	uri := "http://localhost:3000/carriers?pageSize=2&page=2"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri

	// act
	router.ServeHTTP(writer, req)

	// assert
	var jsonBody PaginatedQueryResponse
	err := json.NewDecoder(writer.Body).Decode(&jsonBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, writer.Code, http.StatusOK)
	assert.Equal(t, 2, jsonBody.Total)
	assert.Equal(t, 2, jsonBody.Page)
	assert.Equal(t, 2, jsonBody.PageSize)
	assert.Equal(t, jsonBody.Next, "http://localhost:3000/carriers?pageSize=2&page=3")
	assert.Equal(t, jsonBody.Previous, "http://localhost:3000/carriers?pageSize=2&page=1")
}

func Test_GetCarriersHandler_Should_Not_Set_NextLink(t *testing.T) {
	// arrange
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 1), nil
		},
	}
	w := httptest.NewRecorder()
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})
	router.GET("/carriers", GetCarriersHandler)
	uri := "http://localhost:3000/carriers?pageSize=2&page=2" // because the page size is greater than the results returned we should not set the link
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri

	// act
	router.ServeHTTP(w, req)

	// assert
	var jsonBody PaginatedQueryResponse
	json.Unmarshal(w.Body.Bytes(), &jsonBody)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 1, jsonBody.Total)
	assert.Equal(t, 2, jsonBody.Page)
	assert.Equal(t, 2, jsonBody.PageSize)
	assert.Equal(t, "", jsonBody.Next) // should be no more links since there were less results than requested
	assert.Equal(t, "http://localhost:3000/carriers?pageSize=2&page=1", jsonBody.Previous)
}

func Test_GetCarriersHandler_PreviousLink_Should_Be_Empty(t *testing.T) {
	// arrange
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 2), nil
		},
	}
	w := httptest.NewRecorder()
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	uri := "http://localhost:3000/carriers?pageSize=2&page=0" // because this is page 0 the previous link should not be set
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri

	// act
	router.ServeHTTP(w, req)

	// assert
	var jsonBody PaginatedQueryResponse
	json.Unmarshal(w.Body.Bytes(), &jsonBody)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 2, jsonBody.Total)
	assert.Equal(t, 0, jsonBody.Page)
	assert.Equal(t, 2, jsonBody.PageSize)
	assert.Equal(t, "http://localhost:3000/carriers?pageSize=2&page=1", jsonBody.Next)
	assert.Equal(t, "", jsonBody.Previous) // should be no more links since there were less results than requested
}

func Test_GetCarriersHandler_PreviousLink_Should_Be_Not_Empty(t *testing.T) {
	// arrange
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 2), nil
		},
	}
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()
	uri := "http://localhost:3000/carriers?pageSize=2&page=1" // because this is page 1 (greater than 0) the previous link should be set
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri

	// act
	router.ServeHTTP(w, req)

	// Assert
	var jsonBody PaginatedQueryResponse
	json.NewDecoder(w.Body).Decode(&jsonBody)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 2, jsonBody.Total)
	assert.Equal(t, 1, jsonBody.Page)
	assert.Equal(t, 2, jsonBody.PageSize)
	assert.Equal(t, "http://localhost:3000/carriers?pageSize=2&page=2", jsonBody.Next)
	assert.Equal(t, "http://localhost:3000/carriers?pageSize=2&page=0", jsonBody.Previous) // should be no more links since there were less results than requested
}

func Test_GetCarriersHandler_Should_Default_PageSize_When_Not_Provided(t *testing.T) {
	// Arrange
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 11), nil // need to create >10 (the default size) to ensure we get the next link back with the page size of 10
		},
	}
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers", GetCarriersHandler)
	w := httptest.NewRecorder()

	uri := "http://localhost:3000/carriers?&page=2" // don't pass in `pageSize` as the query parameter to check logic
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	req.RequestURI = uri
	// act

	router.ServeHTTP(w, req)

	// assert
	var jsonBody PaginatedQueryResponse
	json.NewDecoder(w.Body).Decode(&jsonBody)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, 11, jsonBody.Total)
	assert.Equal(t, 2, jsonBody.Page)
	assert.Equal(t, 10, jsonBody.PageSize)
	assert.Equal(t, "http://localhost:3000/carriers?&page=3", jsonBody.Next)
	assert.Equal(t, "http://localhost:3000/carriers?&page=1", jsonBody.Previous) // should be no more links since there were less results than requested
}

func Test_GetCarrierHandler_Should_Have_Status_NotFound_When_Id_Missing_From_Query(t *testing.T) {
	// arrange
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 11), nil // need to create >10 (the default size) to ensure we get the next link back with the page size of 10
		},
		get: func(ctx context.Context, id string) (*models.FreightCarrierModel, error) {
			return &models.FreightCarrierModel{}, nil
		},
	}
	writer := httptest.NewRecorder()
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})
	router.GET("/carriers/:id", GetCarrierHandler)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/carriers/", nil)

	// act
	router.ServeHTTP(writer, req)

	// assert
	assert.Equal(t, http.StatusNotFound, writer.Result().StatusCode)
}

func Test_GetCarrierHandler_Should_Have_StatusInternalServerError(t *testing.T) {
	// arrange
	jsonBody := &struct {
		Error string `json:"error"`
	}{}
	mockDb := &MockCarrierDb{
		all: func(ctx context.Context) ([]models.FreightCarrierModel, error) {
			return make([]models.FreightCarrierModel, 11), nil // need to create >10 (the default size) to ensure we get the next link back with the page size of 10
		},
		get: func(ctx context.Context, id string) (*models.FreightCarrierModel, error) {
			return nil, errors.New("This error should be output in response")
		},
	}

	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})

	router.GET("/carriers/:id", GetCarrierHandler)
	writer := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/carriers/24325272-c9b4-425a-8316-2b6657aa8bf6", nil)

	// act
	router.ServeHTTP(writer, req)

	// assert
	json.NewDecoder(writer.Body).Decode(&jsonBody)
	assert.Equal(t, http.StatusInternalServerError, writer.Result().StatusCode)
	assert.Equal(t, "This error should be output in response", jsonBody.Error)
}

func Test_GetCarrierHandler_Should_Have_Status_OK_And_Carrier_Response_Body(t *testing.T) {
	// arrange
	mockDb := &MockCarrierDb{
		get: func(ctx context.Context, id string) (*models.FreightCarrierModel, error) {
			return &models.FreightCarrierModel{
				ID: "01HZ8TT8D0DMKD12K1YMWP7TF3",
			}, nil
		},
	}
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", mockDb)
	})
	router.GET("/carriers/:id", GetCarrierHandler)
	writer := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/carriers/01HZ8TT8D0DMKD12K1YMWP7TF3", nil)

	// act
	router.ServeHTTP(writer, req)

	// assert
	var jsonBody models.FreightCarrierModel
	json.NewDecoder(writer.Body).Decode(&jsonBody)
	assert.Equal(t, http.StatusOK, writer.Result().StatusCode)
	assert.Equal(t, "01HZ8TT8D0DMKD12K1YMWP7TF3", jsonBody.ID)
}

func Test_CreateCarrierHandler_Should_Have_Status_400_BadRequest_When_Service_Fails(t *testing.T) {
	// arrange
	errBody := &struct {
		Error string `json:"error"`
	}{}
	requestBody := &models.FreightCarrierModel{}
	body, _ := json.Marshal(requestBody)
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		mockDb := MockCarrierDb{
			create: func(ctx context.Context, carrier *models.FreightCarrierModel) error {
				return errors.New("this error happened when creating a carrier")
			},
		}
		ctx.Set("db", &mockDb)
	})
	router.POST("/carriers", CreateCarrierHandler)
	req, _ := http.NewRequest(http.MethodPost, "/carriers", bytes.NewReader(body))
	response := httptest.NewRecorder()

	// act
	router.ServeHTTP(response, req)

	// assert
	s := response.Body.String()
	json.NewDecoder(response.Body).Decode(&errBody)
	assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	assert.NotEqual(t, nil, s)
}

func Test_CreateCarrierHandler_Should_Have_Status_400_BadRequest_When_Validation_Fails(t *testing.T) {
	// arrange
	errBody := &struct {
		Validations map[string][]string `json:"validations"`
	}{}
	requestBody := &models.FreightCarrierModel{}
	body, _ := json.Marshal(requestBody)
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		mockDb := MockCarrierDb{
			create: func(ctx context.Context, carrier *models.FreightCarrierModel) error {
				return nil
			},
		}
		ctx.Set("db", &mockDb)
	})
	router.Use(func(ctx *gin.Context) {
		ctx.Set(string(CreateCarrierValidatorKey), validators.CreateValidatorFunc(func(ctx context.Context, val any) []error {
			return []error{errors.New("Failed")}
		}))
	})
	router.POST("/carriers", CreateCarrierHandler)
	req, _ := http.NewRequest(http.MethodPost, "/carriers", bytes.NewReader(body))
	response := httptest.NewRecorder()

	// act
	router.ServeHTTP(response, req)

	// assert
	s := response.Body.String()
	json.NewDecoder(response.Body).Decode(&errBody)
	assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	assert.NotEqual(t, nil, s)
	assert.Equal(t, 1, len(errBody.Validations["errors"]))
	assert.Equal(t, "Failed", errBody.Validations["errors"][0])
}
