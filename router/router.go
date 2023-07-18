package router

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/ping", HanldePing)
	// Items routes
	r.GET("/items", GETItems)
	r.POST("/items", POSTItems)
	r.DELETE("/items/:id", DELETEItemById)

	return r
}
