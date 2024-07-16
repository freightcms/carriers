package e2e

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

func setup(t *testing.T) {
	t.Setenv("CONNECTION_STRING", "mongodb://admin:admin@localhost:27017/")
	t.Setenv("Environment", "testing")
	t.Setenv("Product", "carriers")
}

func Test_GetCarriers_Should_Be_Paginated_Result(t *testing.T) {
	// arrange
	body := make(map[string]interface{})
	setup(t)
	// act
	r, err := http.Get("/carriers")
	err = json.NewDecoder(r.Body).Decode(&body)
	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, body["total"], 0)
	assert.Equal(t, body["page"], 0)
	assert.Equal(t, body["pageSize"], 10)
	assert.Equal(t, body["pages"], 1)
	assert.Equal(t, body["next"], "")
	assert.Equal(t, body["previous"], "")
}

func Test_GetCarriers_Should_Have_Next_Link(t *testing.T) {
	// arrange
	body := make(map[string]interface{})
	setup(t)

	// act
	r, err := http.Get("/carriers?page=0&pageSize=1")

	assert.Equal(t, err, nil)

	err = json.NewDecoder(r.Body).Decode(&body)

	// assert
	assert.Equal(t, err, nil)
	assert.Equal(t, body["total"], 2)
	assert.Equal(t, body["page"], 1)
	assert.Equal(t, body["pageSize"], 1)
	assert.Equal(t, body["pages"], 2)
	assert.Equal(t, body["entities"], make(map[string]interface{}))
	assert.Equal(t, body["next"], "/carriers?page=1&pageSize=1")
	assert.Equal(t, body["previous"], "/carriers?page=0&pageSize=1")
}
