package database

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() {
	db, err := sql.Open("sqlite3", "./instajot.db")
	if err != nil {
		log.Println("An error occurred opening the database connection.")
		panic(err)
	}

	db.
}