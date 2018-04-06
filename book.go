package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var (
	// LibraryPath - temp directory for book storage
	LibraryPath = "/Users/gabekahen/Documents/library/"
)

// Book structure
type Book struct {
	UID         int64
	Title       string
	Author      string
	Publisher   string
	PublishDate time.Time
	Rating      int
	Status      int
}

// NewBook creates a new book given a valid JSON []byte slice.
// Performs some basic validation on input data
func NewBook(data map[string][]string) (*Book, error) {
	book := Book{}

	for key, value := range data {
		switch key {
		case `Title`:
			book.Title = value[0]
		case `Author`:
			book.Author = value[0]
		case `Publisher`:
			book.Publisher = value[0]
		case `PublishDate`:
			i, err := strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("NewBook: Invalid UNIX date: %s", value[0])
			}
			book.PublishDate = time.Unix(i, 0).UTC()
		case `Rating`:
			i, err := strconv.ParseInt(value[0], 10, 0)
			if err != nil {
				return nil, fmt.Errorf("NewBook: Invalid Rating: %s", value[0])
			}
			book.Rating = int(i)
		case `Status`:
			i, err := strconv.ParseInt(value[0], 10, 0)
			if err != nil {
				return nil, fmt.Errorf("NewBook: Invalid Status: %s", value[0])
			}
			book.Status = int(i)
		}
	}

	return &book, nil
}

// returns a JSON-encoded book object
func (book *Book) print() []byte {
	output, _ := json.Marshal(book)
	return output
}

// create write a new book object to the database. It returns the UID of the
// object, and any errors returned by the database. If an error is returned,
// create() returns int64(0) with the error.
func (book *Book) create() (int64, error) {
	db, err := connectDB()
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(
		"INSERT INTO books (title, author, publisher, publishdate, rating, status) VALUES (?, ?, ?, FROM_UNIXTIME(?), ?, ?)",
		book.Title,
		book.Author,
		book.Publisher,
		book.PublishDate.Unix(),
		book.Rating,
		book.Status,
	)

	if err != nil {
		return 0, err
	}

	// Grab the UID from the last insert ID
	uid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uid, nil
}

// reads the book object from storage.
// Throws errors if object does not exist or is inaccessible.
func (book *Book) read() error {
	return nil
}

// Removes the book object from the database. Returns an error on failure.
func (book *Book) delete() error {
	db, err := connectDB()
	if err != nil {
		return err
	}

	result, err := db.Exec(
		`DELETE FROM books WHERE uid = ?`,
		book.UID,
	)
	if err != nil {
		return err
	}

	retUID, err := result.LastInsertId()

	if retUID != book.UID {
		return fmt.Errorf("delete() looking for UID '%d', but got '%d'", book.UID, retUID)
	}

}

// Helper function checks for presence of book in storage
func (book *Book) exist() bool {
	return false
}
