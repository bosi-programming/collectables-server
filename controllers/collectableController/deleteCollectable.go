package collectableController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

// DeleteCollectable godoc
// @Summary Delete a collectable
// @Success 200 {object} Collectable
// @Router /collectables/:id [delete]
func DeleteCollectable(c *gin.Context) {
	var collectable models.Collectable
	if err := models.DB.Where("id = ?", c.Param("id")).First(&collectable).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&collectable)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
