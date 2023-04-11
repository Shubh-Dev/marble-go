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

type OrderItem struct {
	gorm.Model

	ID       uint `gorm:"primaryKey" json:"id"`
	Quantity uint `json:"quantity"`

	Order     Order
	OrderID   uint `json:"order_id" gorm:"not null"`
	Product   Product
	ProductID uint `json:"product_id"`
}

type OrderDiscount struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"id"`

	OrderID    uint     `json:"order_id"`
	Order      Order    `json:"order" gorm:"foreignKey:OrderID"`
	DiscountID uint     `json:"discount_id"`
	Discount   Discount `json:"discount" gorm:"foreignKey:DiscountID"`
	CouponID   uint     `json:"coupon_id"`
	Coupon     Coupon   `json:"coupon" gorm:"foreignKey:CouponID"`
}
