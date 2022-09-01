package test

import (
	"estj/src/router"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController(t *testing.T) {
	// Start app.
	StartingTest()

	// Given
	route := router.GetRouter().GetRouter()
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	// When
	route.ServeHTTP(w, req)

	// Then
	assert.Equal(t, http.StatusOK, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, "haha", string(responseData))

	// End app.
	EndTest()
}
