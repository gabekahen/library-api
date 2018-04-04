package main

import (
	"testing"
)

func TestInitSchema(t *testing.T) {
	err := initSchema()
	if err != nil {
		t.Error(err)
	}
}
