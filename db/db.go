package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// --- Initialize the Database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	// --- Error check and crash app if no connection
	if err != nil {
		panic("Could not connect to Database.")
	}

	// pool size for ongoing DB connections -- any more than what is defined will have to wait
	DB.SetMaxOpenConns(10)
	// how many DB connections stay open if unused
	DB.SetMaxIdleConns(5)

	// create tables (if necessary) on initial run
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create Users Table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create Events Table.")
	}
}
