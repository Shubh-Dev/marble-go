package models

import "gorm.io/gorm"

type ProductStock struct {
	gorm.Model

	Quantity        int64 `json:"quantity"`

	Product         Product
	ProductID       int64 `json:"product_id"`
	StockLocation   StockLocation
	StockLocationID int64 `json:"stock_location_id"`
}
