package test

import (
	"log"
	"os"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/devaliakbar/gin_gorm_example/internal/core/database"
	"github.com/devaliakbar/gin_gorm_example/internal/features/department"
	"github.com/devaliakbar/gin_gorm_example/internal/features/employee"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	shutDown()
	os.Exit(code)
}

func setUp() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	log.Println("<<<Test Started>>>")

	db, err := gorm.Open(sqlite.Open("testdb.db"), &gorm.Config{})

	if err != nil {
		panic("Connot connect to db")
	}

	database.DB = db

	router := gin.Default()
	department.InitDepartment(router)
	employee.InitEmployee(router)

	unitTest.SetRouter(router)
}

func shutDown() {
	///DROP ALL TABLES
	database.DB.Migrator().DropTable(&department.Department{}, employee.Employee{})

	log.Println("<<<Test End>>>")
}
