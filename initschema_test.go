package main

import (
	"testing"
)

func TestInitSchema(t *testing.T) {
	_, err := initSchema()
	if err != nil {
		t.Fatal(err)
	}
}
