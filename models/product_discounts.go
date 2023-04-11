package models

import "gorm.io/gorm"

type ProductDiscount struct {
	gorm.Model

	ID         uint `gorm:"primaryKey" json:"id"`

	Product    Product
	ProductID  uint `json:"product_id"`
	Discount   Discount
	DiscountID uint `json:"discount_id"`
}
