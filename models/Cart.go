package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Product []Product

	ID     string `gorm:"primaryKey" json:"id"`
	UserID string `json:"userId"`
	Count  byte   `json:"count"`
}
