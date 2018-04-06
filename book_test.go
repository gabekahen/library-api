package main

import (
	"encoding/json"
	"testing"
	"time"
)

var (
	book = Book{
		Title:       "",
		Author:      "Mr. Author",
		Publisher:   "Mr. Publisher",
		PublishDate: time.Now(),
		Rating:      4,
		Status:      0,
	}
)

func TestPrint(t *testing.T) {
	var outBook Book
	testPrintBook := book
	testPrintBook.Title = "TestPrint"

	json.Unmarshal(testPrintBook.print(), &outBook)
	if outBook.Title != testPrintBook.Title {
		t.Errorf(
			"Book.Print(): Expected Title '%s', got '%s'",
			testPrintBook.Title,
			outBook.Title,
		)
	}
}

func TestCreate(t *testing.T) {
	testWriteBook := book
	testWriteBook.Title = "TestWriteBook"

	t.Run("Create book TestWrite", func(t *testing.T) {
		err := testWriteBook.create()
		t.Logf("Book UID: %d", testWriteBook.UID)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Create book TestWrite2 with same author, publisher", func(t *testing.T) {
		testWriteBook2 := testWriteBook
		testWriteBook2.Title = "TestWriteBook2"
		err := testWriteBook2.create()
		if err != nil {
			t.Error(err)
		}
		testWriteBook2.delete()
	})
	t.Run("Fail on duplicate book TestWrite", func(t *testing.T) {
		err := testWriteBook.create()
		if err == nil {
			t.Error("Expected failure on duplicate book, but got none")
		}
	})

	testWriteBook.delete()
}

func TestRead(t *testing.T) {
	testReadBook := book
	testReadBook.Title = "testReadBook"

	testReadBook.create()

	newReadBook := Book{UID: testReadBook.UID}

	err := newReadBook.read()
	if err != nil {
		t.Error(err)
	}

	if newReadBook.Title != testReadBook.Title {
		t.Logf("%+v", newReadBook)
		t.Errorf("Expected Publisher '%s', got '%s'", testReadBook.Publisher, newReadBook.Publisher)
	}
	newReadBook.delete()
}

func TestDelete(t *testing.T) {
	testDeleteBook := book
	testDeleteBook.Title = "testDeleteBook"

	testDeleteBook.create() // create the book to be deleted
	t.Run("Delete book from TestWrite", func(t *testing.T) {
		err := testDeleteBook.delete()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Fail to delete missing book", func(t *testing.T) {
		err := testDeleteBook.delete()
		if err == nil {
			t.Errorf("No error received on deleting missing book %d", book.UID)
		}
	})
}

// test the update method for the Book struct
func TestUpdate(t *testing.T) {
	testUpdateBook := book
	testUpdateBook.Title = "testUpdateBook"

	err := testUpdateBook.create()
	if err != nil {
		t.Error(err)
	}

	t.Run("Test update Rating", func(t *testing.T) {
		testBook := testUpdateBook
		testBook.Title = "testUpdateTitle"
		testBook.Rating = 1 // previously set to 4, above

		err := testBook.update()
		if err != nil {
			t.Error(err)
		}

		newBook := Book{UID: testBook.UID}
		newBook.read()
		if newBook.Rating != testBook.Rating {
			t.Errorf("Expected rating '%d', but got '%d'", testBook.Rating, newBook.Rating)
		}
	})

	t.Run("Commit the same update multiple times", func(t *testing.T) {
		testBook := testUpdateBook
		testBook.Status = 1
		for i := 0; i < 5; i++ {
			err := testBook.update()
			if err != nil {
				t.Error(err)
			}
		}
	})
	testUpdateBook.delete()
}

// TestNewBook tests the NewBook constructor. It should be able to create and
// write a new Book object to storage, but should return an error if the object
// already exists in storage.
func TestNewBook(t *testing.T) {
	bookString := map[string][]string{
		"Title":       []string{"TestNewBook"},
		"Author":      []string{"NewBook Author"},
		"Publisher":   []string{"NewBook Publisher"},
		"PublishDate": []string{"00000000"},
		"Rating":      []string{"3"},
		"Status":      []string{"0"},
	}

	testNewBook, err := NewBook(bookString)
	if err != nil {
		t.Error(err)
	}

	err = testNewBook.create()
	if err != nil {
		t.Error(err)
	}

	newBook := Book{UID: testNewBook.UID}
	err = newBook.read()
	if err != nil {
		t.Error(err)
	}

	if newBook.Author != testNewBook.Author {
		t.Logf("%+v", newBook)
		t.Errorf("Expected author '%s', but got '%s'", testNewBook.Author, newBook.Author)
	}
	newBook.delete()
}
