package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func libraryCreateHandler(w http.ResponseWriter, r *http.Request) {
	book, err := NewBook(r.URL.Query())
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.exist() {
		err := fmt.Errorf("ERROR creating book - Book already exists in storage: %s", book.UID)
		log.Print(err)
		http.Error(w, err.Error(), http.StatusConflict)
	}

	// If the book was created successfully, return the book's UID
	log.Printf("Created new book: %s", book.UID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, fmt.Sprintf(`{"UID": "%s"}`, book.UID))
}

func libraryDeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/create", libraryCreateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
