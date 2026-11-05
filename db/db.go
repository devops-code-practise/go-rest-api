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
	createTableSQL := `CREATE TABLE IF NOT EXISTS events (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,	
		"name" TEXT NOT NULL,
		"description" TEXT,
		"location" TEXT,
		"dateTime" DATETIME,
		user_id INTEGER
	);`	
	
	_, err := DB.Exec(createTableSQL)
	
	if err != nil {
		panic("Could not create table.")
	}
}