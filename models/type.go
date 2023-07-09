package models

import (
	"github.com/jinzhu/gorm"
)

type Type struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name" binding:"required"`
	UserID        int   `json:"user_id" binding:"required"`
	User        *User   `json:"user" binding:"required"`
}
