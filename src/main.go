package main

import (
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/src/controller"
	"github.com/devaliakbar/gin_gorm_example/src/models"
)

func main() {
	r := gin.Default()

	models.ConnectToDb()

	r.GET("/departments", controller.GetAllDepartment)

	r.POST("/department", controller.CreateDepartment)

	r.GET("/department/:id", controller.GetDepartment)

	r.PATCH("/department/:id", controller.UpdateDepartment)

	r.DELETE("/department/:id", controller.DeleteDepartment)

	////

	r.GET("/employees", controller.GetAllEmployee)

	r.POST("/employee", controller.CreateEmployee)

	r.GET("/employee/:id", controller.GetEmployee)

	r.PATCH("/employee/:id", controller.UpdateEmployee)

	r.DELETE("/employee/:id", controller.DeleteEmployee)

	r.Run("localhost:8080")
}
