package main

import (
	"go-server/m/controllers"
	"go-server/m/controllers/bookController"
	"go-server/m/controllers/userController"
	"go-server/m/middlewares"
	"go-server/m/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDataBase()

	router.POST("/register", userController.CreateUser)
	router.POST("/login", controllers.Login)

	protected := router.Group("/")

	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("/users", userController.GetUsers)
	protected.GET("/books", bookController.GetBooks)
	protected.POST("/books", bookController.CreateBook)
	protected.GET("/books/:id", bookController.FindBookById)
	protected.PATCH("/books/:id", bookController.UpdateBook)
	protected.DELETE("/books/:id", bookController.DeleteBook)

	router.Run(":8080")
}
