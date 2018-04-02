package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLibraryCreateAndDelete(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	create_req, err := http.NewRequest("GET", `/create?Title=API test create`, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a delete request to pass to the delete handler
	delete_req, err := http.NewRequest("GET", `/delete/LLJvMjsQPn0t79mIhOvLjPkbjTw=`, nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handlerCreate := http.HandlerFunc(libraryCreateHandler)
	handlerDelete := http.HandlerFunc(libraryDeleteHandler)

	// Serve the create request to the response recorder
	handlerCreate.ServeHTTP(rr, create_req)

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
	handlerCreate.ServeHTTP(rr, create_req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusConflict)
	}

	// Reset the response recorder (rr), and attempt to delete the new book object.
	rr = httptest.NewRecorder()
	handlerDelete.ServeHTTP(rr, delete_req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Reset the response recorder (rr), and attempt to delete the
	// already deleted book object. Should return a HTTP 500
	rr = httptest.NewRecorder()
	handlerDelete.ServeHTTP(rr, delete_req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
