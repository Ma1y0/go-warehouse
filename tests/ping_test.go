package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ma1y0/go-warehouse/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := router.CreateRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)
	
	// Expexted value
	expectedBody := gin.H{"message": "Pong"}

	// Parse response to json
	var response map[string]string
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		t.FailNow()
	}
	// Extract message
	value, exists := response["message"]

	assert.Equal(t, 200, w.Code)
	assert.True(t, exists)
	assert.Equal(t, expectedBody["message"], value)
}
