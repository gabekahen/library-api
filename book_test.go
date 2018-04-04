package main

import (
	"encoding/json"
	"testing"
	"time"
)

var (
	book = Book{
		Title:       "My Book",
		Author:      "Mr. Author",
		Publisher:   "Mr. Publisher",
		PublishDate: time.Now(),
		Rating:      4,
		Status:      0,
	}
)

func TestPrint(t *testing.T) {
	var outBook Book
	json.Unmarshal(book.print(), &outBook)
	if outBook.Title != "My Book" {
		t.Errorf("Book.Print(): Expected Title 'My Book', got '%v'", outBook.Title)
	}
}

func TestWrite(t *testing.T) {
	_, err := book.create()
	if err != nil {
		t.Error(err)
	}
}

/*
func TestRead(t *testing.T) {
	newBook := Book{UID: book.UID}

	err := newBook.read()
	if err != nil {
		t.Error(err)
	}
	if newBook.Title != "My Book" {
		t.Errorf("Book.Print(): Expected Title 'My Book', got '%v'", newBook.Title)
	}
}

func TestDelete(t *testing.T) {
	err := book.delete()
	if err != nil {
		t.Error(err)
	}
}

// TestNewBook tests the NewBook constructor. It should be able to create and
// write a new Book object to storage, but should return an error if the object
// already exists in storage.
func TestNewBook(t *testing.T) {
	bookString := map[string][]string{
		"Title":       []string{"NewBook Tester"},
		"Author":      []string{"NewBook Author"},
		"Publisher":   []string{"NewBook Publisher"},
		"PublishDate": []string{"00000000"},
		"Rating":      []string{"3"},
		"Status":      []string{"0"},
	}

	book, err := NewBook(bookString)
	if err != nil {
		t.Error(err)
	}

	if len(book.UID) < 1 {
		t.Errorf("No UID Present")
	}
}

func TestBookExist(t *testing.T) {
	err := book.create()
	if err != nil {
		t.Errorf("Error writing book: %s", err)
	}

	if book.exist() == false {
		t.Errorf("Book was written, but exist() returned false")
	}

	book.delete()

	if book.exist() == true {
		t.Errorf("Book was delted, but exist() returned true")
	}
}
*/
