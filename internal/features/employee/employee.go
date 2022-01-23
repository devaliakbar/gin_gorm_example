package employee

import (
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/internal/core/database"
)

func InitEmployee(r *gin.Engine) {
	database.DB.AutoMigrate(&Employee{})

	employeeController := EmployeeController{}

	r.GET("/employees", employeeController.getAllEmployee)

	r.POST("/employee", employeeController.createEmployee)

	r.GET("/employee/:id", employeeController.getEmployee)

	r.PATCH("/employee/:id", employeeController.updateEmployee)

	r.DELETE("/employee/:id", employeeController.deleteEmployee)

}
