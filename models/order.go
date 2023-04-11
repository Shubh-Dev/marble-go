package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	ID         uint     `gorm:"primaryKey" json:"id"`
	Date       time.Time
	TotalPrice float64 `json:"total_price"`
	OrderType  string  `json:"order_type"`
	
	CustomerID uint     `json:"customer_id" gorm:"not null"`
	Customer   Customer `gorm:"foreignKey:CustomerID"`
}
