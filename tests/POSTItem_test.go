package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Ma1y0/go-warehouse/models"
	"github.com/Ma1y0/go-warehouse/router"
	"github.com/stretchr/testify/assert"
)

func TestPOSTItemROute(t *testing.T) {
	r := router.CreateRouter()
	models.ConnecToDatabase()

	type data struct {
		  Name  string `json:"name"`
			Amount uint    `json:"amount"`
	}

	data1 := data{
		Name: "Item1",
		Amount: 1,
	}

	jsonData1, err := json.Marshal(data1)
	if err != nil {
		panic(err.Error())
	}


	w := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonData1))					
	r.ServeHTTP(w, req1)

	var data1DB models.ItemModel
	a := data1DB
	models.DB.Where("name = ?", "Item1").First(&data1DB)


  // Tests
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.True(t, data1DB != a)

	// Clen Up
	if err := os.Remove("main.db"); err != nil {
		panic("Failed to delete db")
	}
}
