package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
		Book{ID: 2, Title: "Start With Why", Author: "Simon Sinek", Year: "2009"},
		Book{ID: 3, Title: "Deep Work", Author: "Cal Newport", Year: "2016"},
	)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBooks).Methods("POST")
	router.HandleFunc("/books", updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBooks(w http.ResponseWriter, r *http.Request) {
	var book Book

	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove a book")
}
