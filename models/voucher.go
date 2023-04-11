package models

import "gorm.io/gorm"

type Voucher struct {
	gorm.Model
	ID    uint    `gorm:"primaryKey" json:"id"`
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}
