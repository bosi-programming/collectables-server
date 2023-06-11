package models

type Book struct {
	ID     int    `json:"id" binding:"required" gorm:"primary_key"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
  PlaceOfBook string `json:"placeOfBook" binding:"required"`
}
