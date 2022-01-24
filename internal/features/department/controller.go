package department

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/internal/core/database"
)

type departmentController struct{}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET ALL DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (departmentController) getAllDepartment(c *gin.Context) {
	var departments []Department

	database.DB.Find(&departments)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    departments,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**CREATE DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (departmentController) createDepartment(c *gin.Context) {
	var input CreateDepartmentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	department := Department{Name: input.Name}
	database.DB.Create(&department)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    department,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (departmentController) getDepartment(c *gin.Context) {
	var department Department

	if err := database.DB.Where("id = ?", c.Param("id")).First(&department).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    department,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**UPDATE DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (departmentController) updateDepartment(c *gin.Context) {
	var department Department
	if err := database.DB.Where("id = ?", c.Param("id")).First(&department).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	var input UpdateDepartmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	updateBody := map[string]interface{}{}
	if input.Name != "" {
		updateBody["name"] = input.Name
	}

	database.DB.Model(&department).Updates(updateBody)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    department,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**DELETE DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (departmentController) deleteDepartment(c *gin.Context) {
	var department Department

	if err := database.DB.Where("id = ?", c.Param("id")).First(&department).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	database.DB.Delete(&department)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    department,
	})
}
