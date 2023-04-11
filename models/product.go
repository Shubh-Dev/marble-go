package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name" gorm:"not null"`
	Description string  `gorm:"size:1000 not null" json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
}
