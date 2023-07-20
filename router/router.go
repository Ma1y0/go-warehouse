package router

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/ping", HanldePing)
	// Items routes
	r.GET("/items", GETItems)
	r.GET("/items/:id", GETItemById)
	r.POST("/items", POSTItems)
	r.DELETE("/items/:id", DELETEItemById)

	return r
}
