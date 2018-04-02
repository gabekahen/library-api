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
		PublishDate: time.Now().Unix(),
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
	t.Log(newBook)
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
