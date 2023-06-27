package models

import (
	"github.com/jinzhu/gorm"
)

type Type struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name" binding:"required"`
}
