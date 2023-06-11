package bookController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"net/http"
)

// DeleteBook godoc
// @Summary Delete a book
// @Success 200 {object} Book
// @Router /books/:id [delete]
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
