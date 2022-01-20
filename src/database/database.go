package database

import (
	"github.com/devaliakbar/gin_gorm_example/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Connot connect to db")
	}

	db.AutoMigrate(&models.Department{}, &models.Employee{})

	DB = db
}
