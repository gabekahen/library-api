package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func libraryCreateHandler(w http.ResponseWriter, r *http.Request) {
	book := Book{}

	for key, value := range r.URL.Query() {
		switch key {
		case `Title`:
			book.Title = value[0]
		case `Author`:
			book.Author = value[0]
		case `Publisher`:
			book.Publisher = value[0]
		case `PublishDate`:
			i, err := strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Printf("Invalid UNIX date: %s\n", value[0])
				return
			}
			book.PublishDate = time.Unix(i, 0).UTC()
		case `Rating`:
			i, err := strconv.ParseInt(value[0], 10, 0)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Printf("Invalid Rating: %s\n", value[0])
				return
			}
			book.Rating = int(i)
		case `Status`:
			i, err := strconv.ParseInt(value[0], 10, 0)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Printf("Invalid Status: %s\n", value[0])
				return
			}
			book.Status = int(i)
		}
	}

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
