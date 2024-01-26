package models

import (
	"github.com/jinzhu/gorm"
)

type Collectable struct {
	gorm.Model
	Title       string `gorm:"size:255;not null" json:"title" binding:"required"`
	Author      string `gorm:"size:255;not null" json:"author" binding:"required"`
	Category      string `gorm:"size:255;not null" json:"category" binding:"required"`
	SubCategory      string `gorm:"size:255;not null" json:"subCategory" binding:"required"`
	Type        *Type   `gorm:"size:255;not null" json:"type" binding:"required"`
	UserID        int   `json:"user_id" binding:"required"`
	User        *User   `json:"user" binding:"required"`
}
