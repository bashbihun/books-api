package controllers

import (
	"books-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBook(c *gin.Context) {
	var books []models.Book
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})

}

func CreateBook(c *gin.Context) {
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}
