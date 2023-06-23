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

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hasUser := models.DB.Where("username = ?", input.UserName).First(&models.User{})

	if hasUser != nil && hasUser.Value.(*models.User).Username == input.UserName {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exists!", "info": hasUser.Value.(*models.User).Username, "input": input})
		return
	}

	user := models.User{}
	user.Username = input.UserName
	user.Password = input.Password

	_, err := user.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
