package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kuthumipepple/book-management-system/pkg/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {}
func UpdateBook(w http.ResponseWriter, r *http.Request)  {}
func DeleteBook(w http.ResponseWriter, r *http.Request)  {}
