package test

import (
	"estj/src/router"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController_GET_users(t *testing.T) {
	// Given
	route := router.GetRouter().GetRouter()
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	// When
	route.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.NotEmpty(t, string(responseData))
}
