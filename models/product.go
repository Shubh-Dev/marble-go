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

type ProductDiscount struct {
	gorm.Model

	ID         uint `gorm:"primaryKey" json:"id"`

	Product    Product
	ProductID  uint `json:"product_id"`
	Discount   Discount
	DiscountID uint `json:"discount_id"`
}

type ProductStock struct {
	gorm.Model

	Quantity        int64 `json:"quantity"`

	Product         Product
	ProductID       int64 `json:"product_id"`
	StockLocation   StockLocation
	StockLocationID int64 `json:"stock_location_id"`
}

type ProductVariation struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey" json:"id"`

	Product   Product
	ProductID uint   `json:"product_id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}

type ProductImage struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	ImageURL  string `json:"image_url"`


	Product   Product
	ProductID uint   `json:"product_id" gorm:"not null"`
}