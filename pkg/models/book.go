package models

import (
	"github.com/kuthumipepple/book-management-system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book{
	var books []Book
	db.Find(&books)
	return books
}