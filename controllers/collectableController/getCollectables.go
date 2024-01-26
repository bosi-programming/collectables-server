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

	var category = c.Query("category")

	if category != "" {
		err := models.DB.Where("category LIKE ? AND user_id = ?", "%"+category+"%", user_id).Find(&collectables).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No collectables by this category", "params": category})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": collectables})
		return
	}

	var subCategory = c.Query("subCategory")

	if subCategory != "" {
		err := models.DB.Where("sub_category LIKE ? AND user_id = ?", "%"+subCategory+"%", user_id).Find(&collectables).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No collectables by this subCategory", "params": subCategory})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": collectables})
		return
	}

	var collectableType = c.Query("type")

	if collectableType != "" {
		err := models.DB.Where("type LIKE ? AND user_id = ?", "%"+collectableType+"%", user_id).Find(&collectables).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No collectables by this type", "params": collectableType})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": collectables})
		return
	}

	models.DB.Where("user_id = ?", user_id).Find(&collectables)

	c.JSON(http.StatusOK, gin.H{"data": collectables})
}
