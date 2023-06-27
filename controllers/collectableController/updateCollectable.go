package collectableController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

type UpdateCollectableInput struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	PlaceOfCollectable string `json:"placeOfCollectable"`
}

// UpdateCollectable godoc
// @Summary Update a collectable
// @Success 200 {object} Collectable
// @Router /collectables/:id [patch]
func UpdateCollectable(c *gin.Context) {
	var collectable models.Collectable
	if err := models.DB.Where("id = ?", c.Param("id")).First(&collectable).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateCollectableInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&collectable).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": collectable})
}
