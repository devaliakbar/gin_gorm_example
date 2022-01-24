package department

import (
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/internal/core/database"
)

func InitDepartment(r *gin.Engine) {
	database.DB.AutoMigrate(&Department{})

	departmentCntr := departmentController{}

	r.GET("/departments", departmentCntr.getAllDepartment)

	r.POST("/department", departmentCntr.createDepartment)

	r.GET("/department/:id", departmentCntr.getDepartment)

	r.PATCH("/department/:id", departmentCntr.updateDepartment)

	r.DELETE("/department/:id", departmentCntr.deleteDepartment)
}
