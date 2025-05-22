package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-project/models"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database
	DB.AutoMigrate(&models.Product{}, &models.Cart{}, &models.Category{}, &models.Payment{})
}
