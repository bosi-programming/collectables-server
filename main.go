package main

import (
	"go-server/m/controllers/bookController"
	"go-server/m/controllers/userController"
	"go-server/m/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDataBase()

	router.GET("/users", userController.GetUsers)
	router.POST("/register", userController.CreateUser)

	router.GET("/books", bookController.GetBooks)
	router.POST("/books", bookController.CreateBook)
	router.GET("/books/:id", bookController.FindBookById)
	router.PATCH("/books/:id", bookController.UpdateBook)
	router.DELETE("/books/:id", bookController.DeleteBook)

	router.Run()
}
