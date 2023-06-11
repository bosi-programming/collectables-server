package userController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

// GetUsers godoc
// @Summary Get all users
// @Success 200 {object} []User
// @Router /users [get]

func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
