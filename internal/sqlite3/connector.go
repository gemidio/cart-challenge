package sqlite3

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func Connect() error {
	database, _ = sql.Open("sqlite3", os.Getenv("SQLITE_DATABASE"))
	err := database.Ping()

	if err != nil {
		return err
	}

	return nil
}
