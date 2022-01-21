package database

import (
	departmentModels "github.com/devaliakbar/gin_gorm_example/lib/features/department/models"
	employeeModels "github.com/devaliakbar/gin_gorm_example/lib/features/employee/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Connot connect to db")
	}

	db.AutoMigrate(&departmentModels.Department{}, &employeeModels.Employee{})

	DB = db
}
