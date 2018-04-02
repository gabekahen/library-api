package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func libraryCreateHandler(w http.ResponseWriter, r *http.Request) {
	book := Book{}

	err := book.write()
	if err != nil {
		log.Printf("Error writing book: %s\n", err)
	}

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf(`{"UID": "%s"}`, book.UID))
}

func main() {
	http.HandleFunc("/create", libraryCreateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
