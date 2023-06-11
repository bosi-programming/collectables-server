package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
