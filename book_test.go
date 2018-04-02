package main

import (
	"encoding/json"
	"testing"
	"time"
)

var (
	book = Book{
		UID:         "12345",
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
		t.Fatalf("Book.Print(): Expected Title 'My Book', got '%v'", outBook.Title)
	}
}

func TestWrite(t *testing.T) {
	err := book.write()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRead(t *testing.T) {
	newBook := Book{UID: book.UID}

	err := newBook.read()
	if err != nil {
		t.Fatal(err)
	}
	if newBook.Title != "My Book" {
		t.Fatalf("Book.Print(): Expected Title 'My Book', got '%v'", newBook.Title)
	}
}

func TestDelete(t *testing.T) {
	err := book.delete()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewBook(t *testing.T) {
	bookString := []byte(`{"Title": "NewBook Tester"}`)

	book, err := NewBook(bookString)
	if err != nil {
		t.Fatal(err)
	}

	if len(book.UID) < 1 {
		t.Fatalf("No UID Present")
	}

	_, err = NewBook(bookString)
	if err == nil {
		t.Fatal("Expected error on duplicate write, but got none")
	}

	book.delete()
}
