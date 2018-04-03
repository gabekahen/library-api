package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func initSchema() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:netsys1!@tcp(localhost:3306)/")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS library")
	if err != nil {
		return nil, err
	}

	return db, nil
}
