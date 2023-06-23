package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title       string `gorm:"size:255;not null" json:"title" binding:"required"`
	Author      string `gorm:"size:255;not null" json:"author" binding:"required"`
	PlaceOfBook string `gorm:"size:255;not null" json:"placeOfBook" binding:"required"`
}
