package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name         string
	UserId       uint
	Categories   []Category
	Locations    []Location
	Transactions []Transaction
}
