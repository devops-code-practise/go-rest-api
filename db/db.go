package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
var DB *sql.DB
 
func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "api.db")
 
    if err != nil {
        panic("Could not connect to database.")
    }
 
    // DB.SetMaxOpenConns(10)
    DB.SetMaxOpenConns(10) // better practise when working with SQlite in go.
    DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"email" TEXT NOT NULL UNIQUE,
		"username" TEXT NOT NULL UNIQUE,
		"password" TEXT NOT NULL
	);`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create table.")
	}

	createEventTable := `CREATE TABLE IF NOT EXISTS events (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,	
		"name" TEXT NOT NULL,
		"description" TEXT,
		"location" TEXT,
		"dateTime" DATETIME,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`	
	
	_, err = DB.Exec(createEventTable)
	
	if err != nil {
		panic("Could not create table.")
	}

	createRegistrationTable := `CREATE TABLE IF NOT EXISTS registrations (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"user_id" INTEGER,
		"event_id" INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (event_id) REFERENCES events(id)
	);`

	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("Could not create table.")
	}
}