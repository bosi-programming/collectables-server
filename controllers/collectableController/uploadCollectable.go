package collectableController

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/zenthangplus/goccm"
	"go-server/m/models"
	"go-server/m/utils/token"
	"net/http"
)

func UploadCollectable(c *gin.Context) {
	user_id, _ := token.ExtractTokenID(c)
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file := form.File["file"][0]
	typeId := form.Value["typeId"]
	if typeId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "typeId is required"})
		return
	}

	user := models.DB.Where("id = ?", user_id).First(&models.User{}).Value.(*models.User)

	csvFile, err := file.Open()

	records, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Creating and updating collectables from csv file"})
	comm, _, err := c.Writer.Hijack()
	comm.Close()
	goc := goccm.New(10)

	for index, record := range records {
		if index == 0 {
			continue
		}
		goc.Wait()
		go func(record []string) {
			createOrUpdateCollectable(record, user)
			goc.Done()
		}(record)
	}

	goc.WaitAllDone()
}

func createOrUpdateCollectable(record []string, user *models.User) {
	title, author, category, subCategory, collectableType := record[0], record[1], record[2], record[3], record[4]
	collectable := models.Collectable{Title: title, Author: author, Category: category, SubCategory: subCategory, Type: collectableType, User: user}

	hasCollectable := models.DB.Where("title = ? AND user_id = ?", title, user.ID).Find(&models.Collectable{})
	println(title, hasCollectable)
	if hasCollectable != nil && hasCollectable.Value.(*models.Collectable).Title == title {
		models.DB.Model(&models.Collectable{}).Where("title = ? AND user_id = ?", title, user.ID).Updates(&collectable)
	} else {
		models.DB.Create(&collectable)
	}
}
