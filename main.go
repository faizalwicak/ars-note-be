package main

import (
	"log"
	"os"

	"github.com/faizalwicak/ars-note-be/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func DbInit() *gorm.DB {
	db, err := models.Setup()
	if err != nil {
		log.Println("Problem setting up database")
	}
	return db
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	r := SetupRouter()

	log.Fatal(r.Run(host + ":" + port))
}
