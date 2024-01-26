package collectableController

import (
	"go-server/m/models"
	"go-server/m/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCollectableInput struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Category  string `json:"category" binding:"required"`
	SubCategory  string `json:"subCategory" binding:"required"`
	Type  string `json:"type" binding:"required"`
}

// CreateCollectable godoc
// @Summary Create a collectable
// @Success 200 {object} Collectable
// @Router /collectables [post]
func CreateCollectable(c *gin.Context) {
	user_id, _ := token.ExtractTokenID(c)
	var input CreateCollectableInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.DB.Where("id = ?", user_id).First(&models.User{}).Value.(*models.User)

	collectable := models.Collectable{Title: input.Title, Author: input.Author, Category: input.Category, SubCategory: input.SubCategory, User: user, Type: input.Type}
	models.DB.Create(&collectable)

	c.JSON(http.StatusOK, gin.H{"data": collectable})
}
