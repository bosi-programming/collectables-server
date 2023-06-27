package collectableController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

// FindCollectableById godoc
// @Summary Find a collectable by Id
// @Success 200 {object} Collectable
// @Router /collectables/:id [get]
func FindCollectableById(c *gin.Context) {
	var collectable models.Collectable

	if err := models.DB.Where("id = ?", c.Param("id")).First(&collectable).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collectable})
}
