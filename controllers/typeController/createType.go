package typeController

import (
	"go-server/m/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCollectableTypeInput struct {
	Name string `json:"name" binding:"required"`
}

// CreateCollectable godoc
// @Summary Create a collectable
// @Success 200 {object} Collectable
// @Router /collectables [post]
func CreateCollectable(c *gin.Context) {
	var input CreateCollectableTypeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hasCollectableType := models.DB.Where("name = ?", input.Name).First(&models.Type{})

	if hasCollectableType != nil && hasCollectableType.Value.(*models.Type).Name == input.Name {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Collectable Type already exists!", "info": hasCollectableType.Value.(*models.Type).Name, "input": input})
		return
	}

	collectableType := models.Type{Name: input.Name}
	models.DB.Create(&collectableType)

	c.JSON(http.StatusOK, gin.H{"data": collectableType})
}
