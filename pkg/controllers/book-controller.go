package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kuthumipepple/book-management-system/pkg/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookPtr, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookPtr)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {}
func DeleteBook(w http.ResponseWriter, r *http.Request) {}
