package models

import (
	"github.com/kuthumipepple/book-management-system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var b Book
	db := db.Where("id = ?", id).Find(&b)
	return &b, db
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func DeleteBook(id int64) Book {
	var b Book
	db.Where("id = ?", id).Delete(&b)
	return b
}
