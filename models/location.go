package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name         string `json:"name"`
	BookId       uint   `json:"book_id"`
	Book         Book   `json:"-"`
	Transactions []Transaction
}
