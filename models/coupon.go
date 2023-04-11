package models

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey" json:"id"`
	Code       string    `json:"code"`
	ExpiryDate time.Time `json:"expiry_date"`

	DiscountID uint      `json:"discount_id"`
	Discount   Discount  `json:"discount" gorm:"foreignKey:DiscountID"`
}
