package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name         string `json:"name"`
	UserId       uint   `json:"user_id"`
	Categories   []Category
	Locations    []Location
	Transactions []Transaction
}
