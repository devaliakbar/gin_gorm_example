package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/src/models"
)

func InitEmployeeRoutes(r *gin.Engine) {
	r.GET("/employees", GetAllEmployee)

	r.POST("/employee", CreateEmployee)

	r.GET("/employee/:id", GetEmployee)

	r.PATCH("/employee/:id", UpdateEmployee)

	r.DELETE("/employee/:id", DeleteEmployee)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET ALL EMPLOYEE**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAllEmployee(c *gin.Context) {
	type EmployeeSelection struct {
		EmployeeId         string `json:"employee_id"`
		EmployeeName       string `json:"employee_name"`
		EmployeeDepartment string `json:"employee_department"`
	}

	var employees []EmployeeSelection

	models.DB.Table("employees").
		Joins("inner join departments on departments.id = employees.id").
		Select("employees.id as employee_id, employees.name as employee_name, departments.name as employee_department").
		Find(&employees)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employees,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**CREATE EMPLOYEE**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type CreateEmployeeInput struct {
	Name         string `json:"name" binding:"required"`
	DepartmentId int    `json:"department_id"`
}

func CreateEmployee(c *gin.Context) {
	var input CreateEmployeeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	var department models.Department

	if err := models.DB.Where("id = ?", input.DepartmentId).First(&department).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Department not found!",
		})

		return
	}

	employee := models.Employee{Name: input.Name, Department: department}
	models.DB.Create(&employee)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET EMPLOYEE**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetEmployee(c *gin.Context) {
	var employee models.Employee

	if err := models.DB.Preload("Department").First(&employee, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**UPDATE EMPLOYEE**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type UpdateEmployeeInput struct {
	Name         string `json:"name"`
	DepartmentId int    `json:"department_id"`
}

func UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := models.DB.Where("id = ?", c.Param("id")).Preload("Department").First(&employee).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	var input UpdateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	if input.Name != "" {
		employee.Name = input.Name
	}

	if input.DepartmentId != 0 {
		var department models.Department

		if err := models.DB.Where("id = ?", input.DepartmentId).First(&department).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Department not found!",
			})

			return
		}

		employee.Department = department
	}

	models.DB.Save(&employee)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**DELETE EMPLOYEE**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func DeleteEmployee(c *gin.Context) {
	var employee models.Employee

	if err := models.DB.Where("id = ?", c.Param("id")).Preload("Department").First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	models.DB.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}
