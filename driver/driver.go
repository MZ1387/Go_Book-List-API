package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
