package models

import (
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	ImageURL  string `json:"image_url"`


	Product   Product
	ProductID uint   `json:"product_id" gorm:"not null"`
}
