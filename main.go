package main

import (
	"keshika/database"
	"log"

	"keshika/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	routers := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	router.RegisterRoutes(routers)

	routers.Run(":4000")

}
