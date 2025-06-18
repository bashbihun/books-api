package main

import (
	"books-api/config"
	"books-api/models"
	"books-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Book{})

	r := gin.Default()
	routes.SetUpRoute(r, db)
	r.Run(":8080")
}
