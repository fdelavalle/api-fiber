package models

import "gorm.io/gorm"

// Book is a struct that represents a book in the database
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}
