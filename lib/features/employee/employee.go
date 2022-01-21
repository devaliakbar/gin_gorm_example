package employee

import (
	"github.com/gin-gonic/gin"

	employeeControllers "github.com/devaliakbar/gin_gorm_example/lib/features/employee/controllers"
)

func InitRoutes(r *gin.Engine) {

	r.GET("/employees", employeeControllers.GetAllEmployee)

	r.POST("/employee", employeeControllers.CreateEmployee)

	r.GET("/employee/:id", employeeControllers.GetEmployee)

	r.PATCH("/employee/:id", employeeControllers.UpdateEmployee)

	r.DELETE("/employee/:id", employeeControllers.DeleteEmployee)

}
