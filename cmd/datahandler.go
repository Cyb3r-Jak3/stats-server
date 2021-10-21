package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// DbURL represents the full url for the database connection
var DbURL string

var tableStrings = []string{
	"CREATE TABLE IF NOT EXISTS report404(URL text NOT NULL DEFAULT '', IP text NOT NULL DEFAULT '', Method text NOT NULL DEFAULT '', time integer);",
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

func insert(insertString string, sqlArgs ...interface{}) error {
	db, err := sql.Open("postgres", DbURL)
	if err != nil {
		log.WithError(err).Error("Error opening database")
		return err
	}
	stmt, err := db.Prepare(insertString)
	if err != nil {
		log.WithError(err).Error("Error preparing statement")
		return err
	}

	_, err = stmt.Exec(sqlArgs...)
	if err != nil {
		log.WithError(err).Error("Error executing stmt")
		return err
	}
	if err = stmt.Close(); err != nil {
		log.WithError(err).Error("Error closing statement")
		return err
	}
	if err = db.Close(); err != nil {
		log.WithError(err).Error("Error closing DB")
		return err
	}
	return nil
}

//func query(querystring string, sqlArgs ...interface{}) (rows []*sql.Rows, err error) {
//	db, err := sql.Open("postgres", DbURL)
//	if err != nil {
//		log.WithError(err).Error("Error opening database")
//	}
//	stmt, err := db.Prepare(querystring)
//	if err != nil {
//		log.WithError(err).Error("Error preparing statement")
//	}
//	rows= stmt.QueryRow(sqlArgs)
//	return
//}
