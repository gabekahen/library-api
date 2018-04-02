package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
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

func (book *Book) print() []byte {
	output, _ := json.Marshal(book)
	return output
}

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
