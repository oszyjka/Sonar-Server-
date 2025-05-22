package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-project/models"
)

var DBTest *gorm.DB

func ConnectTestDB() {
	var err error
	DBTest, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	err = DBTest.AutoMigrate(&models.Cart{}, &models.Product{}, &models.Category{}, &models.Payment{})
	if err != nil {
		panic("Failed to migrate test database")
	}
}
