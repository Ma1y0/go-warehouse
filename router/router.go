package router

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/ping", HanldePing)

	return r
}
