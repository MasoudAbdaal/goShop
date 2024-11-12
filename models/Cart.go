package models

type Cart struct {
	ID        string `gorm:"primaryKey" json:"id"`
	UserID    string `json:"userId"`
	ProductID int16  `json:"productId"`
	Count     byte   `json:"count"`

	Product Product `gorm:"foreignKey:ProductID"`
}
