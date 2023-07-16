package router

import (
	"net/http"

	"github.com/Ma1y0/go-warehouse/models"
	"github.com/gin-gonic/gin"
)

// Get all items
// GET /items
func GETItems(c *gin.Context) {
	var items []models.ItemModel

	models.DB.Find(&items)

	c.JSON(http.StatusOK, gin.H{"items": items})
}


// Create item 
// POST /items
func POSTItems(c *gin.Context) {
	var input models.ItemModel

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Body", "error": err.Error()})
	}

	item := models.ItemModel{
		Name: input.Name,
		Amount: input.Amount,
	}

	if err := models.DB.Create(&item); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create item", "error": err.Error.Error()})
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created a new item", "item": item})
}
