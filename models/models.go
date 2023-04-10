package models

import "gorm.io/gorm"

// Book model
type Marble struct {
	gorm.Model

	Title  string `json:"title"`
}