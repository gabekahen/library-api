package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("LIBRARYAPI_DB_USER", "root")
	os.Setenv("LIBRARYAPI_DB_PASS", "")
	os.Setenv("LIBRARYAPI_DB_PROT", "tcp")
	os.Setenv("LIBRARYAPI_DB_HOST", "localhost:3306")
	os.Exit(m.Run())
}

func TestLibraryCreateAndDelete(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	createReq, err := http.NewRequest("GET", `/create?Title=API test create`, nil)
	if err != nil {
		t.Error(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handlerCreate := http.HandlerFunc(libraryCreateHandler)

	// Serve the create request to the response recorder
	handlerCreate.ServeHTTP(rr, createReq)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Create a delete request to pass to the delete handler.
	// The delete request needs the UID generated by the request to /create .
	responseBook := Book{}
	json.Unmarshal(rr.Body.Bytes(), &responseBook)

	deleteReq, err := http.NewRequest("GET", `/delete/`+responseBook.UID, nil)
	if err != nil {
		t.Error(err)
	}

	handlerDelete := http.HandlerFunc(libraryDeleteHandler)

	// Reset the response recorder (rr), and attempt to delete the new book object.
	rr = httptest.NewRecorder()
	handlerDelete.ServeHTTP(rr, deleteReq)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Reset the response recorder (rr), and attempt to delete the
	// already deleted book object. Should return a HTTP 500
	rr = httptest.NewRecorder()
	handlerDelete.ServeHTTP(rr, deleteReq)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// Tests the /read/ endpoint
func TestLibraryRead(t *testing.T) {
	bookData := map[string][]string{
		"Title": []string{"test read endpoint"},
	}

	// create new book object
	book, err := NewBook(bookData)
	if err != nil {
		t.Error(err)
	}

	// write new book to storage for reading
	_, err = book.create()
	if err != nil {
		t.Error(err)
	}

	readReq, err := http.NewRequest("GET", `/read/`+book.UID, nil)
	if err != nil {
		t.Error(err)
	}

	rr := httptest.NewRecorder()
	handlerRead := http.HandlerFunc(libraryReadHandler)

	handlerRead.ServeHTTP(rr, readReq)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"UID":"` + book.UID + `","Title":"test read endpoint","Author":"","Publisher":"","PublishDate":"0001-01-01T00:00:00Z","Rating":0,"Status":0}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// cleanup test book
	book.delete()

	// test read on missing book
	rr = httptest.NewRecorder()
	handlerRead.ServeHTTP(rr, readReq)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
