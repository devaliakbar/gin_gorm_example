package server

import (
	"github.com/devaliakbar/gin_gorm_example/internals/core/database"
	"github.com/devaliakbar/gin_gorm_example/internals/core/middleware"
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/internals/features/department"
	"github.com/devaliakbar/gin_gorm_example/internals/features/employee"
)

func RunServer() {

	database.InitializeDb()

	r := gin.Default()

	r.Use(middleware.Logger())

	department.InitDepartment(r)
	employee.InitEmployee(r)

	r.Run("localhost:8080")
}
