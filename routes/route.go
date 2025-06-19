package routes

import (
	"books-api/controllers"
	"books-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoute(r *gin.Engine, db *gorm.DB) {
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		r.GET("/books", controllers.GetBook)
		r.POST("/books", controllers.CreateBook)
	}
}
