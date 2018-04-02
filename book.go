package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
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
	PublishDate int64
	Rating      int
	Status      int
}

// returns a JSON-encoded book object
func (book *Book) print() []byte {
	output, _ := json.Marshal(book)
	return output
}

// writes the book object to storage
func (book *Book) write() error {
	hasher := sha1.New()
	hasher.Write([]byte(book.Title + book.Author + book.Publisher))
	uid := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	book.UID = uid
	file, err := os.Create(LibraryPath + uid)
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
