package department

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/lib/core/database"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET ALL DEPARTMENT**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func getAllDepartment(c *gin.Context) {
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
type CreateDepartmentInput struct {
	Name string `json:"name" binding:"required"`
}

func createDepartment(c *gin.Context) {
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
func getDepartment(c *gin.Context) {
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
type UpdateDepartmentInput struct {
	Name string `json:"name"`
}

func updateDepartment(c *gin.Context) {
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

	//models.DB.Model(&department).Updates(input)

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
func deleteDepartment(c *gin.Context) {
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