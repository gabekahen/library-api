package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
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

	book.genuid()

	return &book, nil
}

// returns a JSON-encoded book object
func (book *Book) print() []byte {
	output, _ := json.Marshal(book)
	return output
}

// generates UID for book object
func (book *Book) genuid() {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	uid := make([]byte, 32)

	for i := range uid {
		uid[i] = charset[seededRand.Intn(len(charset))]
	}
	book.UID = string(uid)
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

// reads the book object from storage.
// Throws errors if object does not exist or is inaccessible.
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

// Helper function checks for presence of book in storage
func (book *Book) exist() bool {
	_, err := os.Stat(LibraryPath + book.UID)
	if err == nil {
		return true
	}
	return false
}
