package main

import "github.com/Ma1y0/go-warehouse/router"

func main() {
	// Initialize Router
	router := router.CreateRouter()

	if err := router.Run(":8080"); err != nil {
		panic("Failed to bind server")
	}
}
