package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Ma1y0/go-warehouse/models"
	"github.com/Ma1y0/go-warehouse/router"
	"github.com/stretchr/testify/assert"
)

func Test_DELETE_item_by_id(t *testing.T) {
	r := router.CreateRouter()
	models.ConnecToDatabase()

	// Crete record
	test_item := models.ItemModel{
		Name: "Test item",
	}

	if err := models.DB.Create(&test_item); err.Error != nil {
		panic(err.Error.Error())
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/items/1", nil)
	r.ServeHTTP(w, req)
	//
	// // Test
	assert.Equal(t, http.StatusOK, w.Code)
	assert.False(t, checkDB(1))


	// Delete DB 
	os.Remove("main.db")
}


func checkDB(id uint) bool {
	var item models.ItemModel
	if err := models.DB.First(&item, id); err.Error != nil {
		return false
	}

	return true
}
