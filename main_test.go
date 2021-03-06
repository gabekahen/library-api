package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	os.Setenv("LIBRARYAPI_DB_USER", "root")
	os.Setenv("LIBRARYAPI_DB_PASS", "")
	os.Setenv("LIBRARYAPI_DB_PROT", "tcp")
	os.Setenv("LIBRARYAPI_DB_HOST", "localhost:3306")

	// make sure the schema exists before we start our tests
	err := initSchema()
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestLibraryCreateAndDelete(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	createReq, err := http.NewRequest(
		"GET",
		`/create?Title=TestLibraryCreateAndDelete&Author=Mr. Author&PublishDate=12345678`,
		nil,
	)
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
			status, http.StatusCreated)
	}

	// Create a delete request to pass to the delete handler.
	// The delete request needs the UID generated by the request to /create .
	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)

	request := fmt.Sprintf("/delete/%s", response["UID"])
	t.Logf("Delete request: %s", request)
	deleteReq, err := http.NewRequest("GET", request, nil)
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
		"Title":       []string{"TestLibraryRead"},
		"Author":      []string{"Author Authington"},
		"Publisher":   []string{"Mr. Publisher"},
		"PublishDate": []string{"1522863678"},
		"Rating":      []string{"5"},
		"Status":      []string{"0"},
	}

	// create new book object
	book, err := NewBook(bookData)
	if err != nil {
		t.Error(err)
	}

	// write new book to storage for reading
	err = book.create()
	if err != nil {
		t.Error(err)
	}

	uidString := fmt.Sprintf("%d", book.UID)
	readReq, err := http.NewRequest("GET", `/read/`+uidString, nil)
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

func TestLibraryUpdate(t *testing.T) {
	baseBook := Book{
		Title:       "TestLibraryUpdate",
		Author:      "Mr. TestLibraryUpdate",
		Publisher:   "TestLibraryUpdate Publishing Co",
		PublishDate: time.Now(),
		Rating:      3,
		Status:      1,
	}

	handlerUpdate := http.HandlerFunc(libraryUpdateHandler)

	// Try updating the status from 1 -> 22, and the rating from 3 -> 2
	t.Run("Update Status & Rating", func(t *testing.T) {
		book := baseBook

		err := book.create()
		if err != nil {
			t.Error(err)
		}

		reqString := fmt.Sprintf(
			"/update?UID=%d&Status=%d&Rating=%d",
			book.UID,
			22,
			2,
		)
		updateReq, err := http.NewRequest("GET", reqString, nil)
		if err != nil {
			t.Error(err)
		}

		rr := httptest.NewRecorder()
		handlerUpdate.ServeHTTP(rr, updateReq)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Logf("%+v", string(rr.Body.Bytes()))
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check if the book was updated
		newBook := Book{UID: book.UID}
		newBook.read()
		if newBook.Status != 22 {
			t.Errorf("Expected book status 22, got %d", newBook.Status)
		}
		if newBook.Rating != 2 {
			t.Errorf("Expected book rating 2, got %d", newBook.Rating)
		}
		book.delete() // cleanup
	})

	// Update the rating from 3 -> 5. Status should remain 1
	t.Run("Update Rating only", func(t *testing.T) {
		book := baseBook

		err := book.create()
		if err != nil {
			t.Error(err)
		}

		reqString := fmt.Sprintf(
			"/update?UID=%d&Rating=%d",
			book.UID,
			5,
		)
		updateReq, err := http.NewRequest("GET", reqString, nil)
		if err != nil {
			t.Error(err)
		}

		rr := httptest.NewRecorder()
		handlerUpdate.ServeHTTP(rr, updateReq)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Logf("%+v", string(rr.Body.Bytes()))
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check if the book was updated
		newBook := Book{UID: book.UID}
		newBook.read()
		if newBook.Status != 1 {
			t.Errorf("Expected book status 1, but got %d", newBook.Status)
		}
		if newBook.Rating != 5 {
			t.Errorf("Expected book rating 5, got %d", newBook.Rating)
		}
		book.delete()
	})
}
