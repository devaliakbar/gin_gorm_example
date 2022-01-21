package employee

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/lib/core/database"
	"github.com/devaliakbar/gin_gorm_example/lib/features/department"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET ALL EMPLOYEE**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func getAllEmployee(c *gin.Context) {
	type EmployeeSelection struct {
		EmployeeId         string `json:"employee_id"`
		EmployeeName       string `json:"employee_name"`
		EmployeeDepartment string `json:"employee_department"`
	}

	var employees []EmployeeSelection

	database.DB.Table("employees").
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

func createEmployee(c *gin.Context) {
	var input CreateEmployeeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	var department department.Department

	if err := database.DB.Where("id = ?", input.DepartmentId).First(&department).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Department not found!",
		})

		return
	}

	employee := Employee{Name: input.Name, Department: department}
	database.DB.Create(&employee)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET EMPLOYEE**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func getEmployee(c *gin.Context) {
	var employee Employee

	if err := database.DB.Preload("Department").First(&employee, "id = ?", c.Param("id")).Error; err != nil {
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

func updateEmployee(c *gin.Context) {
	var employee Employee
	if err := database.DB.Where("id = ?", c.Param("id")).Preload("Department").First(&employee).Error; err != nil {

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
		var department department.Department

		if err := database.DB.Where("id = ?", input.DepartmentId).First(&department).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Department not found!",
			})

			return
		}

		employee.Department = department
	}

	database.DB.Save(&employee)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**DELETE EMPLOYEE**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func deleteEmployee(c *gin.Context) {
	var employee Employee

	if err := database.DB.Where("id = ?", c.Param("id")).Preload("Department").First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	database.DB.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}
