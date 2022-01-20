package main

import (
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/src/controller"
	"github.com/devaliakbar/gin_gorm_example/src/middleware"
	"github.com/devaliakbar/gin_gorm_example/src/models"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	models.ConnectToDb()

	r.Use(middleware.Logger())

	controller.InitDepartmentRoutes(r)

	controller.InitEmployeeRoutes(r)

	r.Run("localhost:8080")
}
