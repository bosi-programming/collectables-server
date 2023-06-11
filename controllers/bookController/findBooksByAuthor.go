package bookController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

// FindBooksByAuthor godoc
// @Summary Find books by an author
// @Success 200 {object} []Book
// @Router /books/:id [get]
func FindBooksByAuthor(c *gin.Context) {
	var books []models.Book

	if err := models.DB.Where("author = ?", c.Param("author")).Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}
