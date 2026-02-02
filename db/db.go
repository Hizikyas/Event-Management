package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)
var DB *sql.DB

func InitDB() {
	var err error
	DB , err = sql.Open("sqlite" , "app.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetConnMaxIdleTime(5)
	DB.SetMaxOpenConns(10)
	CreateTable()
}

func CreateTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime  TEXT NOT NULL,
	user_id INTEGER
	)`
	_, err := DB.Exec(createEventsTable)

	if err !=nil {
		panic("could not create table")
	}
}