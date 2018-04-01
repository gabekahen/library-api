package main

import (
  "encoding/json"
)
// Book structure
type Book struct {
  UID uint64
  Title string
  Author string
  Publisher string
  PublishDate uint64
  Rating int
  Status int
}

func (book *Book) print() []byte {
  output, _ := json.Marshal(book)
  return output
}
