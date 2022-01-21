package server

import (
	"github.com/devaliakbar/gin_gorm_example/lib/core/database"
	"github.com/devaliakbar/gin_gorm_example/lib/core/middleware"
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/lib/features/department"
	"github.com/devaliakbar/gin_gorm_example/lib/features/employee"
)

func RunServer() {

	database.InitializeDb()

	r := gin.Default()

	r.Use(middleware.Logger())

	department.InitRoutes(r)
	employee.InitRoutes(r)

	r.Run("localhost:8080")
}
