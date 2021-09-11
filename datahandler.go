package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// DbURL represents the full url for the database connection
var DbURL string

var tableStrings = []string{
	"CREATE TABLE IF NOT EXISTS report404(URL text NOT NULL DEFAULT '', IP text NOT NULL DEFAULT '', Method text NOT NULL DEFAULT '');",
}

func createTables() {
	db, err := sql.Open("postgres", DbURL)
	if err != nil {
		log.WithError(err).Error("Error opening database")
	}
	for _, x := range tableStrings {
		_, err := db.Exec(x)
		if err != nil {
			log.WithError(err).Fatal("Error opening database")
		}
	}
}

func insert(insertString string, sqlArgs ...interface{}) {
	db, err := sql.Open("postgres", DbURL)
	if err != nil {
		log.WithError(err).Error("Error opening database")
	}
	stmt, err := db.Prepare(insertString)
	if err != nil {
		log.WithError(err).Error("Error preparing statement")
	}
	_, err = stmt.Exec(sqlArgs...)
	if err != nil {
		log.WithError(err).Error("Error executing stmt")
	}
	if err = stmt.Close(); err != nil {
		log.WithError(err).Error("Error closing statement")
	}
	if err = db.Close(); err != nil {
		log.WithError(err).Error("Error closing DB")
	}
}

// func query() {

// }
