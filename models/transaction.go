package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Date        time.Time `json:"date"`
	Value       int       `json:"value"`
	Description string    `json:"description" gorm:"type:text"`
	BookId      uint      `json:"book_id"`
	LocationId  *uint     `json:"location_id" gorm:"default:null"`
	CategoryId  *uint     `json:"category_id" gorm:"default:null"`

	Book     Book     `json:"-"`
	Location Location `json:"-"`
	Category Category `json:"-"`
}
