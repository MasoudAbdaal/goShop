package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Cart Cart

	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"mail"`
}
