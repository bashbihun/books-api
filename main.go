package main

import (
	"books-api/config"
	"books-api/models"
	"books-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db := config.ConnectDB()
	db.AutoMigrate(&models.Book{}, &models.User{})

	r := gin.Default()
	routes.SetUpRoute(r, db)
	r.Run(":8080")
}
