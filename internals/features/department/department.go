package department

import (
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/internals/core/database"
)

func InitDepartment(r *gin.Engine) {
	database.DB.AutoMigrate(&Department{})

	departmentController := DepartmentController{}

	r.GET("/departments", departmentController.getAllDepartment)

	r.POST("/department", departmentController.createDepartment)

	r.GET("/department/:id", departmentController.getDepartment)

	r.PATCH("/department/:id", departmentController.updateDepartment)

	r.DELETE("/department/:id", departmentController.deleteDepartment)
}
