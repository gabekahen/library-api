package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "github.com/docker/go-healthcheck"
)

// HTTP handler creates new book records.
// Book attributes (Title, Publisher, Rating, etc) are passed as URI arguments
func libraryCreateHandler(w http.ResponseWriter, r *http.Request) {
	book, err := NewBook(r.URL.Query())
	if err != nil {
		log.Printf("libraryCreateHandler: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = book.create()
	if err != nil {
		log.Printf("libraryCreateHandler: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the book was created successfully, return the book's UID
	log.Printf("Created new book: %d", book.UID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, fmt.Sprintf(`{"UID": "%d"}`, book.UID))
}

// HTTP handler reads existing book records.
// /read/<UID>
// Throws an error if the book cannot be found / accessed from storage.
func libraryReadHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.ParseInt(r.URL.Path[len(`/read/`):], 10, 0)
	if err != nil {
		log.Printf("libraryReadHandler: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book := Book{
		UID: int(uid),
	}

	err = book.read()
	if err != nil {
		log.Printf("libraryReadHandler: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the book was created successfully, return the book's UID
	log.Printf("Read request for book: %d", book.UID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(book.print())
	if err != nil {
		log.Printf("libraryReadHandler: %s", err)
	}
}

// HTTP handler updates status, rating for existing book records.
// /update?UID=<int>&rating=<int>&status=<int>
// Throws an error if the book cannot be found / accessed from storage.
func libraryUpdateHandler(w http.ResponseWriter, r *http.Request) {
	book, err := NewBook(r.URL.Query())
	if err != nil {
		log.Printf("libraryCreateHandler: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = book.update()
	if err != nil {
		log.Printf("libraryUpdateHandler: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the book was created successfully, return the book's UID
	log.Printf("Update request for book %d", book.UID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(book.print())
	if err != nil {
		log.Printf("libraryReadHandler: %s", err)
	}
}

// Handler deletes books from storage.
// /delete/<UID>
// Throws errors if the book is cannot be deleted.
// TODO: better error handling
func libraryDeleteHandler(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.ParseInt(r.URL.Path[len(`/delete/`):], 10, 0)
	if err != nil {
		log.Printf("libraryDeleteHandler: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book := Book{
		UID: int(uid),
	}

	err = book.read()
	if err != nil {
		log.Printf("libraryDeleteHandler: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = book.delete()
	if err != nil {
		log.Printf("libraryDeleteHandler: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the book was created successfully, return the book's UID
	log.Printf("Delete request for book %d", book.UID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(book.print())
	if err != nil {
		log.Printf("libraryReadHandler: %s", err)
	}
}

func main() {
	err := initSchema()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/create", libraryCreateHandler)
	http.HandleFunc("/read/", libraryReadHandler)
	http.HandleFunc("/update", libraryUpdateHandler)
	http.HandleFunc("/delete/", libraryDeleteHandler)
	log.Print("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
