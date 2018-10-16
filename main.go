package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   string
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(
		books,
		Book{ID: 1, Title: "The Alchemist", Author: "Paulo Coelho", Year: "1988"},
		Book{ID: 1, Title: "Start With Why", Author: "Simon Sinek", Year: "2009"},
		Book{ID: 1, Title: "Deep Work", Author: "Cal Newport", Year: "2016"},
	)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBooks).Methods("POST")
	router.HandleFunc("/books", updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("XXX", w)
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get one book")
}

func addBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book")
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Update books")
}

func removeBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove a book")
}
