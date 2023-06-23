package bookController

import (
	"github.com/gin-gonic/gin"
	"go-server/m/models"
	"go-server/m/utils/token"
	"net/http"
)

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	PlaceOfBook string `json:"placeOfBook" binding:"required"`
}

// CreateBook godoc
// @Summary Create a book
// @Success 200 {object} Book
// @Router /books [post]
func CreateBook(c *gin.Context) {
	user_id, _ := token.ExtractTokenID(c)
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.DB.Where("id = ?", user_id).First(&models.User{}).Value.(*models.User)

	book := models.Book{Title: input.Title, Author: input.Author, PlaceOfBook: input.PlaceOfBook, User: user}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
