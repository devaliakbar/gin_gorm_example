package server

import (
	"github.com/devaliakbar/gin_gorm_example/lib/core/database"
	"github.com/devaliakbar/gin_gorm_example/lib/core/middleware"
	"github.com/gin-gonic/gin"

	departmentControllers "github.com/devaliakbar/gin_gorm_example/lib/features/deparment/controllers"
	employeeControllers "github.com/devaliakbar/gin_gorm_example/lib/features/employee/controllers"
)

func RunServer() {

	database.InitializeDb()

	r := gin.Default()

	r.Use(middleware.Logger())

	departmentControllers.InitDepartmentRoutes(r)
	employeeControllers.InitEmployeeRoutes(r)

	r.Run("localhost:8080")
}
