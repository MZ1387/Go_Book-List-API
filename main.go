package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/lib/pq"

	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

var books []models.Book
var db *sql.DB

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

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBooks).Methods("POST")
	router.HandleFunc("/books", updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	books = []Book{}

	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id=$1", params["id"])
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(book)
}

func addBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)
	err := db.QueryRow(
		"insert into books (title, author, year) values($1, $2, $3) RETURNING ID;",
		book.Title, book.Author, book.Year).Scan(&bookID)

	logFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	result, err := db.Exec(
		"UPDATE books set title=$1, author=$2, year=$3 where ID=$4 RETURNING ID",
		&book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsUpdated)
}

func removeBooks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := db.Exec("DELETE from books where id=$1", params["id"])
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	json.NewEncoder(w).Encode(rowsDeleted)
}
