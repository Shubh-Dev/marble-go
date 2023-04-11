package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name" gorm:"varchar(255);not null"`
	LastName  string `json:"last_name" gorm:"varchar(255);not null"`
	Email     string `json:"email" gorm:"unique_index"`
	Phone     string `json:"phone"`
	UserName  string `json:"user_name" gorm:"unique_index;not null;"`
	Password  string `json:"password"`
}

type CustomerAddress struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey" json:"id"`
	Address     string `json:"address" gorm:"not null"`
	AddressType string `json:"address_type" gorm:"not null"`

	CustomerID uint     `json:"customer_id"`
	Customer   Customer `json:"customer" gorm:"foreignKey:CustomerID"`
}


