package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/src/models"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET ALL DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAllDepartment(c *gin.Context) {
	var departments []models.Department

	models.DB.Find(&departments)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    departments,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**CREATE DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type CreateDepartmentInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateDepartment(c *gin.Context) {
	var input CreateDepartmentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	department := models.Department{Name: input.Name}
	models.DB.Create(&department)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    department,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetDepartment(c *gin.Context) {
	var department models.Department

	if err := models.DB.Where("id = ?", c.Param("id")).First(&department).Error; err != nil {
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
type UpdateDepartmentInput struct {
	Name string `json:"name"`
}

func UpdateDepartment(c *gin.Context) {
	var department models.Department
	if err := models.DB.Where("id = ?", c.Param("id")).First(&department).Error; err != nil {

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

	//models.DB.Model(&department).Updates(input)

	updateBody := map[string]interface{}{}
	if input.Name != "" {
		updateBody["name"] = input.Name
	}

	models.DB.Model(&department).Updates(updateBody)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    department,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**DELETE DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func DeleteDepartment(c *gin.Context) {
	var department models.Department

	if err := models.DB.Where("id = ?", c.Param("id")).First(&department).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	models.DB.Delete(&department)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    department,
	})
}
