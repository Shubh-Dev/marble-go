package models

import "gorm.io/gorm"

type StockLocation struct {
	gorm.Model

	ID      uint   `gorm:"primary key" json:"id"`
	Name    string `json:"name" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
}
