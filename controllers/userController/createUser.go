package userController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

type CreateUserInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUser godoc
// @Summary Create a user
// @Success 200 {object} User
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var input CreateUserInput

	if hasUser := models.DB.Where("username C ?", input.UserName).First(&models.User{}); hasUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exists!"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.UserName, Password: input.Password}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
