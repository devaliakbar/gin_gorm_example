package department

import (
	"github.com/gin-gonic/gin"

	departmentControllers "github.com/devaliakbar/gin_gorm_example/lib/features/department/controllers"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/departments", departmentControllers.GetAllDepartment)

	r.POST("/department", departmentControllers.CreateDepartment)

	r.GET("/department/:id", departmentControllers.GetDepartment)

	r.PATCH("/department/:id", departmentControllers.UpdateDepartment)

	r.DELETE("/department/:id", departmentControllers.DeleteDepartment)
}
