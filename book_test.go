package main

import (
  "testing"
  "encoding/json"
)

func TestPrint (t *testing.T) {
  var outBook Book
  title := "My Book"
  inBook := Book{Title: title}
  json.Unmarshal(inBook.print(), &outBook)
  if outBook.Title != title {
    t.Fatalf("Book.Print(): Expected Title 'My Book', got '%v'", outBook.Title)
  }
}
