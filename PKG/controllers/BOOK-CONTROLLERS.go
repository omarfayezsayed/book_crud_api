package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/omarfayezsayed/book-crud-api/PKG/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	books = models.GetAllBooks()
	res, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	models.CreateBook(book.Name, book.Author)
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("succes"))

}
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	id, _ := strconv.Atoi(params["bookId"])
	book, err := models.GetBookByID(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("not found"))
	} else {
		res, _ := json.Marshal(book)
		w.WriteHeader(http.StatusAccepted)
		w.Write(res)
	}

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	id, _ := strconv.Atoi(params["bookId"])
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	models.UpdateBook(id, book.Name, book.Author)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["bookId"])
	err := models.DeleteBook(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("book not found"))
	} else {
		w.Write([]byte("book deleted"))
	}
}
