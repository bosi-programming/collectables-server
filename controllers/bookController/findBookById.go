package bookController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

// FindBookById godoc
// @Summary Find a book by Id
// @Success 200 {object} Book
// @Router /books/:id [get]
func FindBookById(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
