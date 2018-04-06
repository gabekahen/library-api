package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func getDataSource() string {
	return fmt.Sprintf("%s:%s@%s(%s)/?parseTime=true",
		os.Getenv("LIBRARYAPI_DB_USER"),
		os.Getenv("LIBRARYAPI_DB_PASS"),
		os.Getenv("LIBRARYAPI_DB_PROT"),
		os.Getenv("LIBRARYAPI_DB_HOST"),
	)
}

func initSchema() error {
	db, err := sql.Open("mysql", getDataSource())
	if err != nil {
		return err
	}

	defer db.Close()

	// Create the library database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS library_api")
	if err != nil {
		return err
	}

	// Select the library_api database
	_, err = db.Exec("USE library_api")
	if err != nil {
		return err
	}

	// Setup the table schema for books
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS books (
			uid MEDIUMINT NOT NULL AUTO_INCREMENT,
			title VARCHAR(140) NOT NULL,
			author VARCHAR(60) NOT NULL,
			publisher VARCHAR(60),
			publishdate DATETIME,
			rating TINYINT,
			status TINYINT,
			PRIMARY KEY (uid)
		)`)
	if err != nil {
		return err
	}

	return nil
}

// dbConnect is a helper function that returns a *sql.DB object used to
// issue database queries.
func dbConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", getDataSource())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// use the library_api database
	_, err = db.Exec(`USE library_api`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
