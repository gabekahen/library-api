package main

import (
	"encoding/json"
	"os"
)

var (
	// LibraryPath - temp directory for book storage
	LibraryPath = "/Users/gabekahen/Documents/library"
)

// Book structure
type Book struct {
	UID         uint64
	Title       string
	Author      string
	Publisher   string
	PublishDate int64
	Rating      int
	Status      int
}

func (book *Book) print() []byte {
	output, _ := json.Marshal(book)
	return output
}

func (book *Book) write() error {
	file, err := os.Create(LibraryPath + book.Title)
	if err != nil {
		return err
	}

	_, err = file.Write(book.print())
	if err != nil {
		return err
	}

	return nil
}
