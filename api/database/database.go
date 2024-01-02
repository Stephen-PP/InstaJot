package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./instajot.db")
	if err != nil {
		log.Println("An error occurred opening the database connection.")
		panic(err)
	}

	createNotesDatabase(db)
	createUserDatabase(db)

	return db
}

func createUserDatabase(db *sql.DB) {
	// Create the users table if it doesn't exist
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS Users (
		UserId TEXT PRIMARY KEY,
		Username TEXT,
		Email TEXT,
		Password TEXT,
		CreatedAt TIMESTAMP default (strftime('%s', 'now'))
	)`)

	// Make sure we prepared the table successfully
	if err != nil {
		log.Println("An error occurred while preparing the users table.")
		panic(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("An error occurred creating the users table.")
		panic(err)
	}
}

func createNotesDatabase(db *sql.DB) {
	// Create the notes table if it doesn't exist
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS Notes (
		NoteId TEXT,
		HistoryId TEXT PRIMARY KEY,
		UserId INTEGER,
		Content BLOB,
		CreatedAt TIMESTAMP default (strftime('%s', 'now')),
		FOREIGN KEY (UserId) REFERENCES Users(UserId)
	)`)

	// Make sure we prepared the table successfully
	if err != nil {
		log.Println("An error occurred while preparing the notes table.")
		panic(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println("An error occurred creating the notes table.")
		panic(err)
	}
}
