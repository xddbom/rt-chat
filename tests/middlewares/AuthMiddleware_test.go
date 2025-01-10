package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xddbom/rt-chat/Middlewares"
	"github.com/xddbom/rt-chat/Handlers"
)
func TestAuthMiddleware(t *testing.T) {
	validToken, err := Handlers.GenerateToken("Dude")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	invalidToken := "invalidToken"

	r := gin.Default()
	r.GET("/protected", middlewares.AuthMiddleware(), func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	req, _ := http.NewRequest("GET", "/protected?token="+validToken, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req, _ = http.NewRequest("GET", "/protected?token="+invalidToken, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}