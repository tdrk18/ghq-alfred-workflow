package main

import (
	"fmt"
	"testing"
)

func TestGetQuery(t *testing.T) {
	args := "text\n"
	result := getQuery(args)
	if result != "text" {
		t.Fatal(fmt.Sprintf("failed: getQuery() returns %s", result))
	}
}
