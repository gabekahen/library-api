package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// HTTP handler creates new book records.
// Book attributes (Title, Publisher, Rating, etc) are passed as URI arguments
func libraryCreateHandler(w http.ResponseWriter, r *http.Request) {
	book, err := NewBook(r.URL.Query())
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.exist() {
		err = fmt.Errorf("ERROR creating book - Book already exists in storage: %s", book.UID)
		log.Print(err)
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	err = book.write()
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the book was created successfully, return the book's UID
	log.Printf("Created new book: %s", book.UID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, fmt.Sprintf(`{"UID": "%s"}`, book.UID))
}

// HTTP handler reads existing book records.
// /read/<UID>
// Throws an error if the book cannot be found / accessed from storage.
func libraryReadHandler(w http.ResponseWriter, r *http.Request) {
	book := Book{
		UID: r.URL.Path[len(`/read/`):],
	}

	err := book.read()
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the book was created successfully, return the book's UID
	log.Printf("Read request for book: %s", book.UID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(book.print())
	if err != nil {
		log.Print(err)
	}
}

// Handler deletes books from storage.
// /delete/<UID>
// Throws errors if the book is cannot be deleted.
// TODO: better error handling
func libraryDeleteHandler(w http.ResponseWriter, r *http.Request) {
	book := Book{
		UID: r.URL.Path[len(`/delete/`):],
	}

	err := book.delete()
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/create", libraryCreateHandler)
	http.HandleFunc("/read/", libraryReadHandler)
	http.HandleFunc("/delete/", libraryDeleteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
