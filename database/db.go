package database

import (
	"fmt"
	"keshika/models"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port, os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")))

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database successfully")

	DB.AutoMigrate(&models.Employ{})
	
	fmt.Println("Database Migration done..")

}
