package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLibraryCreateAndDelete(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	createReq, err := http.NewRequest("GET", `/create?Title=API test create`, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a delete request to pass to the delete handler
	deleteReq, err := http.NewRequest("GET", `/delete/LLJvMjsQPn0t79mIhOvLjPkbjTw=`, nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handlerCreate := http.HandlerFunc(libraryCreateHandler)
	handlerDelete := http.HandlerFunc(libraryDeleteHandler)

	// Serve the create request to the response recorder
	handlerCreate.ServeHTTP(rr, createReq)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"UID": "LLJvMjsQPn0t79mIhOvLjPkbjTw="}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Serve the same request again to the response recorder.
	// This will attempt to create a duplicate book in storage.
	rr = httptest.NewRecorder()
	handlerCreate.ServeHTTP(rr, createReq)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusConflict)
	}

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
		t.Fatal(err)
	}

	// write new book to storage for reading
	err = book.write()
	if err != nil {
		t.Fatal(err)
	}

	readReq, err := http.NewRequest("GET", `/read/r_PJOOOqUjZ6SCCTVnywIiuaRS8=`, nil)
	if err != nil {
		t.Fatal(err)
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
	expected := `{"UID":"r_PJOOOqUjZ6SCCTVnywIiuaRS8=","Title":"test read endpoint","Author":"","Publisher":"","PublishDate":"0001-01-01T00:00:00Z","Rating":0,"Status":0}`
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
