package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name         string `json:"Name"`
	BookId       uint
	Book         Book `json:"-"`
	Transactions []Transaction
}
