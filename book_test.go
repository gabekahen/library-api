package main

import (
	"encoding/json"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	var outBook Book
	title := "My Book"
	inBook := Book{Title: title}
	json.Unmarshal(inBook.print(), &outBook)
	if outBook.Title != title {
		t.Fatalf("Book.Print(): Expected Title 'My Book', got '%v'", outBook.Title)
	}
}

func TestWrite(t *testing.T) {
	book := Book{
		Title:       "My Book",
		Author:      "Mr. Author",
		Publisher:   "Mr. Publisher",
		PublishDate: time.Now().Unix(),
		Rating:      4,
		Status:      0,
	}
	err := book.write()
	if err != nil {
		t.Fatal(err)
	}
}
