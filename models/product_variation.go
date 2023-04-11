package models

import "gorm.io/gorm"

type ProductVariation struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey" json:"id"`

	Product   Product
	ProductID uint   `json:"product_id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}
