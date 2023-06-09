package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kuthumipepple/book-management-system/pkg/models"
	"github.com/kuthumipepple/book-management-system/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	err := utils.ParseBody(r, newBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

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
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookUpdate := &models.Book{}
	if err := utils.ParseBody(r, bookUpdate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bookPtr, db := models.GetBookById(id)
	if bookUpdate.Name != "" {
		bookPtr.Name = bookUpdate.Name
	}
	if bookUpdate.Author != "" {
		bookPtr.Author = bookUpdate.Author
	}
	if bookUpdate.Publication != "" {
		bookPtr.Publication = bookUpdate.Publication
	}
	db.Save(bookPtr)

	res, _ := json.Marshal(bookPtr)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
