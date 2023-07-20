package main_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ma1y0/go-warehouse/models"
	"github.com/Ma1y0/go-warehouse/router"
	"github.com/stretchr/testify/assert"
)

func TestGETItemsRoute(t *testing.T) {
	// Initialize router
	r := router.CreateRouter()
	// Initialize database
	models.ConnecToDatabase()
	
	// Test request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	r.ServeHTTP(w, req)

	// Parse JSON 
	var response map[string][]string
	if err := json.Unmarshal([]byte(w.Body.String()), &response); err != nil {
		t.FailNow()
	}

	items, exists := response["items"]

	fmt.Println(items)

	// Test 
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.True(t, exists)
} 
