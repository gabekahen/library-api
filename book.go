package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	// LibraryPath - temp directory for book storage
	LibraryPath = "/Users/gabekahen/Documents/library/"
)

// Book structure
type Book struct {
	UID         string
	Title       string
	Author      string
	Publisher   string
	PublishDate time.Time
	Rating      int
	Status      int
}

// NewBook creates a new book given a valid JSON []byte slice.
// Returns an error if JSON is not valid, or if book already exists in storage.
func NewBook(data []byte) (*Book, error) {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		return nil, err
	}

	book.genuid()

	err = book.read()
	if err == nil {
		return nil, fmt.Errorf("Book already present: %s", book.UID)
	}

	err = book.write()
	if err != nil {
		return nil, fmt.Errorf("Could not write book to storage: %s", book.UID)
	}

	return &book, nil
}

// returns a JSON-encoded book object
func (book *Book) print() []byte {
	output, _ := json.Marshal(book)
	return output
}

// generates UID for book object
func (book *Book) genuid() {
	hasher := sha1.New()
	hasher.Write([]byte(book.Title + book.Author + book.Publisher))
	uid := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	book.UID = uid
}

// writes the book object to storage
func (book *Book) write() error {
	file, err := os.Create(LibraryPath + book.UID)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(book.print())
	if err != nil {
		return err
	}

	return nil
}

// reads the book object from storage
func (book *Book) read() error {
	content, err := ioutil.ReadFile(LibraryPath + book.UID)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &book)
	if err != nil {
		return err
	}

	return nil
}

// Removes the book object from storage
func (book *Book) delete() error {
	err := os.Remove(LibraryPath + book.UID)
	if err != nil {
		return err
	}

	return nil
}
