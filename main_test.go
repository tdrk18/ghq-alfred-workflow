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

func TestGetRepos(t *testing.T) {
	bytes := []byte{'a', '\n', 'b', '\n'}
	result := getRepos(bytes)
	if len(result) != 2 {
		t.Fatal(fmt.Sprintf("failed: getRepos() returns %d contents", len(result)))
	}
	if result[0] != "a" {
		t.Fatal(fmt.Sprintf("failed: getRepos().first returns %s", result[0]))
	}
	if result[1] != "b" {
		t.Fatal(fmt.Sprintf("failed: getRepos().second returns %s", result[1]))
	}
}
