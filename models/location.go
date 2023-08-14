package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name         string `json:"Name"`
	BookId       uint
	Book         Book `json:"-"`
	Transactions []Transaction
}
