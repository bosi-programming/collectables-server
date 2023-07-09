package typeController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"go-server/m/utils/token"
	"net/http"
)

func GetTypes(c *gin.Context) {
	user_id, _ := token.ExtractTokenID(c)
	var types []models.Type
	models.DB.Find(&types)

	var name = c.Query("name")

	if name != "" {
		err := models.DB.Where("name LIKE ? AND user_id = ?", "%"+name+"%", user_id).Find(&types).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No collectables by this author"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": types})
		return
	}

	models.DB.Where("user_id = ?", user_id).Find(&types)
	c.JSON(http.StatusOK, gin.H{"data": types})
}
