package config

import (
	"fmt"
	"os"

	"github.com/athunlal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() (*gorm.DB, error) {
	dns := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	/// Barbers - BackEnd red yellow blue green magenta
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Session{})
	DB.AutoMigrate(&models.FoodAndDrink{})
	DB.AutoMigrate(&models.BuffetOrder{})
	DB.AutoMigrate(&models.Room{})
	DB.AutoMigrate(&models.SessionHistory{})

	/// Barbers - BackEnd red yellow blue green magenta

	/// PetCare - BackEnd @deprecated
	// DB.AutoMigrate(&models.User{})
	// DB.AutoMigrate(&models.Admin{})
	// DB.AutoMigrate(&models.Address{})
	// DB.AutoMigrate(&models.Product{})
	// DB.AutoMigrate(&models.Brand{})
	// DB.AutoMigrate(&models.Cart{})
	// DB.AutoMigrate(&models.Image{})
	// DB.AutoMigrate(&models.Payment{})
	// DB.AutoMigrate(&models.OderDetails{})
	// DB.AutoMigrate(&models.Coupon{})
	// DB.AutoMigrate(&models.Wishlist{})
	// DB.AutoMigrate(&models.Catogery{})
	// DB.AutoMigrate(&models.RazorPay{})
	// DB.AutoMigrate(&models.Oder_item{})
	// DB.AutoMigrate(&models.Wallet{})
	// DB.AutoMigrate(&models.WalletHistory{})

	return DB, nil

}
