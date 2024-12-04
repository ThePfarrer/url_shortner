package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./test_db.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the users table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS url_model (
	id INTEGER NOT NULL,
	"key" VARCHAR,
	long_url VARCHAR,
	short_url VARCHAR,
	PRIMARY KEY (id),
	UNIQUE ("key")
);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Database initialized successfully")
}
