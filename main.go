package main

import (
	"log"
	"os"

	"go-server/m/controllers"
	"go-server/m/controllers/collectableController"
	"go-server/m/controllers/userController"
	"go-server/m/controllers/typeController"
	"go-server/m/middlewares"
	"go-server/m/models"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
  ginMode := os.Getenv("GIN_MODE")
  if ginMode == "release" {
    gin.SetMode(gin.ReleaseMode)
  }
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

	router.Run(":3000")
}
