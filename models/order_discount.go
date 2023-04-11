package models

import "gorm.io/gorm"

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
