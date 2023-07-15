package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Life check route
// GET /ping
func HanldePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Pong"})
}
