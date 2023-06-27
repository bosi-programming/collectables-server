package typeController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

// DeleteType godoc
// @Summary Delete a type
// @Success 200 {object} Type
// @Router /types/:id [delete]
func DeleteType(c *gin.Context) {
	var collectableType models.Type
	if err := models.DB.Where("id = ?", c.Param("id")).First(&collectableType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&collectableType)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
