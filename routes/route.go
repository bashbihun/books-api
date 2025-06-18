package routes

import (
	"books-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoute(r *gin.Engine, db *gorm.DB) {
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/books", controllers.GetBook)
	r.POST("/books", controllers.CreateBook)
}
