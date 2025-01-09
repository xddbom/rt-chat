package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootRout(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"message": "Welcome to real time chat based on Redis and Websockets!"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}