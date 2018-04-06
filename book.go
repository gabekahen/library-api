package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Book structure
type Book struct {
	UID         int
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

// create write a new book object to the database. It returns
// any errors returned by the database.
func (book *Book) create() error {
	db, err := dbConnect()
	if err != nil {
		return err
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
		return err
	}

	// Grab the UID from the last insert ID
	uid, err := result.LastInsertId()
	if err != nil {
		return err
	}

	book.UID = int(uid)

	return nil
}

// reads the book object from storage.
// Throws errors if object does not exist or is inaccessible.
func (book *Book) read() error {
	db, err := dbConnect()
	if err != nil {
		return err
	}

	row := db.QueryRow(
		`SELECT
			uid,
			title,
			author,
			publisher,
			publishdate,
			rating,
			status
		FROM books
		WHERE uid = ?`,
		book.UID,
	)

	err = row.Scan(
		&book.UID,
		&book.Title,
		&book.Author,
		&book.Publisher,
		&book.PublishDate,
		&book.Rating,
		&book.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

// update commits changes to a Book object's status or rating.
// Returns an error on failure.
func (book *Book) update() error {
	db, err := dbConnect()
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`UPDATE books SET rating = ?, status = ? WHERE uid = ?`,
		book.Rating, book.Status, book.UID,
	)
	if err != nil {
		return err
	}

	return nil
}

// Removes the book object from the database. Returns an error on failure.
func (book *Book) delete() error {
	db, err := dbConnect()
	if err != nil {
		return err
	}

	response, err := db.Exec(`DELETE FROM books WHERE uid = ?`, book.UID)
	if err != nil {
		return err
	}

	if rows, _ := response.RowsAffected(); rows == 0 {
		return fmt.Errorf("DELETE FAILED: Book not found: %d", book.UID)
	}
	return nil
}
