package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lib/pq"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   string
}

var books []Book

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))

	logFatal(err)
	log.Println(pgURL)
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBooks).Methods("POST")
	router.HandleFunc("/books", updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {

}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func addBooks(w http.ResponseWriter, r *http.Request) {

}

func updateBooks(w http.ResponseWriter, r *http.Request) {

}

func removeBooks(w http.ResponseWriter, r *http.Request) {

}
