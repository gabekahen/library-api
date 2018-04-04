package main

import (
	"os"
	"testing"
)

func TestInitSchema(t *testing.T) {
	os.Setenv("LIBRARYAPI_DB_USER", "root")
	os.Setenv("LIBRARYAPI_DB_PASS", "")
	os.Setenv("LIBRARYAPI_DB_PROT", "tcp")
	os.Setenv("LIBRARYAPI_DB_HOST", "localhost:3306")
	t.Log(os.Environ())
	err := initSchema()
	if err != nil {
		t.Error(err)
	}
}
