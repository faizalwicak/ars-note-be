package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup() (*gorm.DB, error) {

	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&User{}, &Grocery{}, &Book{}, &Category{}, &Location{}, &Transaction{}); err != nil {
		log.Println(err)
	}

	return db, err
}
