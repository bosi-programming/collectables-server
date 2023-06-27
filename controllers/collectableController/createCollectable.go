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
	PlaceOfCollectable string `json:"placeOfCollectable" binding:"required"`
	TypeID      int    `json:"type_id" binding:"required"`
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
	collectableType := models.DB.Where("id = ?", input.TypeID).First(&models.Type{}).Value.(*models.Type)

	collectable := models.Collectable{Title: input.Title, Author: input.Author, PlaceOfCollectable: input.PlaceOfCollectable, User: user, Type: collectableType}
	models.DB.Create(&collectable)

	c.JSON(http.StatusOK, gin.H{"data": collectable})
}
