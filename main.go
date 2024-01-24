package main

import (
	"go-server/m/controllers"
	"go-server/m/controllers/collectableController"
	"go-server/m/controllers/userController"
	"go-server/m/controllers/typeController"
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

	protected.GET("/collectables", collectableController.GetCollectables)
	protected.POST("/collectables", collectableController.CreateCollectable)
	protected.POST("/collectables/upload", collectableController.UploadCollectable)
	protected.GET("/collectables/:id", collectableController.FindCollectableById)
	protected.PATCH("/collectables/:id", collectableController.UpdateCollectable)
	protected.DELETE("/collectables/:id", collectableController.DeleteCollectable)

	protected.GET("/types", typeController.GetTypes)
	protected.POST("/types", typeController.CreateCollectableType)
	protected.DELETE("/types/:id", typeController.DeleteType)
	protected.GET("/types/:id", typeController.FindTypeById)

	router.Run(":3000")
}
