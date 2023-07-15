package main

import (
	"github.com/Ma1y0/go-warehouse/router"
	"github.com/Ma1y0/go-warehouse/router/models"
)

func main() {
	// Initialize Router
	router := router.CreateRouter()

	// Initialize DB
	models.ConnecToDatabase()

	if err := router.Run(":8080"); err != nil {
		panic("Failed to bind server")
	}
}
