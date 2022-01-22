package department

import (
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/lib/core/database"
)

func InitDepartment(r *gin.Engine) {
	database.DB.AutoMigrate(&Department{})

	r.GET("/departments", getAllDepartment)

	r.POST("/department", createDepartment)

	r.GET("/department/:id", getDepartment)

	r.PATCH("/department/:id", updateDepartment)

	r.DELETE("/department/:id", deleteDepartment)
}
