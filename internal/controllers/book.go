package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rxedu/go-bookstore/v1/internal/marshal"
	"github.com/rxedu/go-bookstore/v1/internal/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	marshal.ParseBody(r, book)
	b := book.CreateBook()
	jsonRes(w, b, http.StatusCreated)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	jsonRes(w, books, http.StatusOK)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, ok := marshal.ParseID(bookId)
	if !ok {
		return
	}
	book, _ := models.GetBookByID(ID)
	jsonRes(w, book, http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {}

func DeleteBook(w http.ResponseWriter, r *http.Request) {}

func jsonRes(w http.ResponseWriter, data interface{}, code int) {
	res, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
