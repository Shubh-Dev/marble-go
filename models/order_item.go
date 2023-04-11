package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model

	ID       uint `gorm:"primaryKey" json:"id"`
	Quantity uint `json:"quantity"`

	Order     Order
	OrderID   uint `json:"order_id" gorm:"not null"`
	Product   Product
	ProductID uint `json:"product_id"`
}
