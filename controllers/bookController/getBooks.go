package bookController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

// GetBooks godoc
// @Summary Get all books
// @Success 200 {object} []Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book

	var author = c.Query("author")

	if author != "" {
		err := models.DB.Where("author LIKE ?", "%"+author+"%").Find(&books).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No books by this author"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": books})
		return
	}

	var title = c.Query("title")

	if title != "" {
		err := models.DB.Where("title LIKE ?", "%"+title+"%").Find(&books).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No books by this title", "params": title})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": books})
		return
	}

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
