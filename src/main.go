package main

import (
	"github.com/gin-gonic/gin"

	"github.com/devaliakbar/gin_gorm_example/src/controller"
	"github.com/devaliakbar/gin_gorm_example/src/models"
)

func main() {
	r := gin.Default()

	models.ConnectToDb()

	r.GET("/books", controller.GetAllBooks)

	r.POST("/book", controller.CreateBook)

	r.GET("/book/:id", controller.GetBook)

	r.PATCH("/book/:id", controller.UpdateBook)

	r.DELETE("/book/:id", controller.DeleteBook)

	r.Run("localhost:8080")
}
