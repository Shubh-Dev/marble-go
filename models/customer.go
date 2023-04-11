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


