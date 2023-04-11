package models

import (
	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	ID         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `json:"name"`
	Percentage float64 `json:"percentage"`
}
