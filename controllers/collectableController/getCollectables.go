package collectableController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"go-server/m/utils/token"
	"net/http"
)

// GetCollectables godoc
// @Summary Get all collectables
// @Success 200 {object} []Collectable
// @Router /collectables [get]
func GetCollectables(c *gin.Context) {
	user_id, _ := token.ExtractTokenID(c)
	var collectables []models.Collectable

	var author = c.Query("author")

	if author != "" {
		err := models.DB.Where("author LIKE ? AND user_id = ?", "%"+author+"%", user_id).Find(&collectables).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No collectables by this author"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": collectables})
		return
	}

	var title = c.Query("title")

	if title != "" {
		err := models.DB.Where("title LIKE ? AND user_id = ?", "%"+title+"%", user_id).Find(&collectables).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No collectables by this title", "params": title})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": collectables})
		return
	}

	models.DB.Where("user_id = ?", user_id).Find(&collectables)

	c.JSON(http.StatusOK, gin.H{"data": collectables})
}
