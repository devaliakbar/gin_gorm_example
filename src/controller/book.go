package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/src/models"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET ALL BOOKS**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetAllBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    books,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**CREATE BOOK**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    book,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**GET BOOK**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func GetBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    book,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**UPDATE BOOK**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	//	models.DB.Model(&book).Updates(input)

	models.DB.Model(&book).Updates(map[string]interface{}{"title": input.Title, "author": input.Author})

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    book,
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**DELETE BOOK**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func DeleteBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Record not found!",
		})

		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    book,
	})
}
