package models

import "gorm.io/gorm"

type CustomerAddress struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey" json:"id"`
	Address     string `json:"address" gorm:"not null"`
	AddressType string `json:"address_type" gorm:"not null"`

	CustomerID uint     `json:"customer_id"`
	Customer   Customer `json:"customer" gorm:"foreignKey:CustomerID"`
}
