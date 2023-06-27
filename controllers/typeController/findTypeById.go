package typeController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

// FindTypeById godoc
// @Summary Find a collectableType by Id
// @Success 200 {object} Type
// @Router /collectableTypes/:id [get]
func FindTypeById(c *gin.Context) {
	var collectableType models.Type

	if err := models.DB.Where("id = ?", c.Param("id")).First(&collectableType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collectableType})
}
