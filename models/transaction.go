package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Name        string    `json:"Name"`
	Date        time.Time `json:"Date"`
	Value       int       `json:"Value"`
	Description string    `gorm:"type:text" json:"Description"`
	BookId      uint
	LocationId  uint
	CategoryId  uint
}
