package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

// LogFatal returning error message
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB connect to DataBase
func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))

	db, err = sql.Open("postgres", pgURL)
	LogFatal(err)

	err = db.Ping()
	LogFatal(err)

	return db

}
