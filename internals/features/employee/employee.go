package employee

import (
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/internals/core/database"
)

func InitEmployee(r *gin.Engine) {
	database.DB.AutoMigrate(&Employee{})

	r.GET("/employees", getAllEmployee)

	r.POST("/employee", createEmployee)

	r.GET("/employee/:id", getEmployee)

	r.PATCH("/employee/:id", updateEmployee)

	r.DELETE("/employee/:id", deleteEmployee)

}
