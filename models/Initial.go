package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitialModels() {
	db, _ := gorm.Open(sqlite.Open("goShop.db"),
		&gorm.Config{Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{LogLevel: logger.Info})})

	_ = db.AutoMigrate(&User{})
	_ = db.AutoMigrate(&Product{})
	_ = db.AutoMigrate(&Cart{})

	db.Create(&Product{0, "Apple", "a delicious apple"})
	db.Create(&Product{1, "Golabi", "a delicious Golabi"})
	db.Create(&Product{2, "Watermelon", "a delicious Watermelon"})

	db.Create(&User{ID: "0-5", Name: "Admin", Email: "Admin@goShop.com"})
	db.Create(&User{ID: "5", Name: "Test", Email: "Test@goShop.com"})
	db.Create(&User{ID: "0-1", Name: "User", Email: "User@goShop.com"})

	db.Create(&Cart{ID: "Cart-0", Count: 192, ProductID: 0, UserID: "0-0"})
	db.Create(&Cart{ID: "Cart-1", Count: 192, ProductID: 0, UserID: "5"})
	db.Create(&Cart{ID: "Cart-2", Count: 192, ProductID: 0, UserID: "5"})
	db.Create(&Cart{ID: "Cart-3", Count: 192, ProductID: 0, UserID: "5"})
	db.Create(&Cart{ID: "Cart-4", Count: 25, ProductID: 2, UserID: "0-1"})
	db.Create(&Cart{ID: "Cart-5", Count: 128, ProductID: 1, UserID: "0-0"})

	var cart User
	db.Preload("Cart").Preload("Cart.Product").First(&cart, User{ID: "5"})

	fmt.Println(cart)
}
