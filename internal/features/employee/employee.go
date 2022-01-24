package employee

import (
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/internal/core/database"
)

func InitEmployee(r *gin.Engine) {
	database.DB.AutoMigrate(&Employee{})

	employeeCntr := employeeController{}

	r.GET("/employees", employeeCntr.getAllEmployee)

	r.POST("/employee", employeeCntr.createEmployee)

	r.GET("/employee/:id", employeeCntr.getEmployee)

	r.PATCH("/employee/:id", employeeCntr.updateEmployee)

	r.DELETE("/employee/:id", employeeCntr.deleteEmployee)

}
