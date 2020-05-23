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

func TestGetRepoName(t *testing.T) {
	path := "/path/to/github.com/tdrk18/repo"
	result1 := getRepoName(path, true)
	if result1 != "tdrk18/repo" {
		t.Fatal(fmt.Sprintf("failed: getRepoName(repo, true) returns %s", result1))
	}
	result2 := getRepoName(path, false)
	if result2 != "github.com/tdrk18/repo" {
		t.Fatal(fmt.Sprintf("failed: getRepoName(repo, false) returns %s", result2))
	}
}

func TestGetDomain(t *testing.T) {
	path := "/path/to/github.com/tdrk18/repo"
	result := getDomain(path)
	if result != "github.com" {
		t.Fatal(fmt.Sprintf("failed: getDomain(repo) returns %s", result))
	}
}

func TestGetIconName(t *testing.T) {
	github := getIconName("/path/to/github.com/tdrk18/repo")
	if github != "github.png" {
		t.Fatal(fmt.Sprintf("failed: getIconName(repo) returns %s", github))
	}
	gitlab := getIconName("/path/to/gitlab.com/tdrk18/repo")
	if gitlab != "gitlab.png" {
		t.Fatal(fmt.Sprintf("failed: getIconName(repo) returns %s", gitlab))
	}
	bitbucket := getIconName("/path/to/bitbucket.org/tdrk18/repo")
	if bitbucket != "bitbucket.png" {
		t.Fatal(fmt.Sprintf("failed: getIconName(repo) returns %s", bitbucket))
	}
	other := getIconName("/path/to/other.com/tdrk18/repo")
	if other != "git.png" {
		t.Fatal(fmt.Sprintf("failed: getIconName(repo) returns %s", other))
	}
}

func TestGetRepoURL(t *testing.T) {
	path := "/path/to/github.com/tdrk18/repo"
	result := getRepoURL(path)
	if result != "https://github.com/tdrk18/repo" {
		t.Fatal(fmt.Sprintf("failed: getRepoURL() returns %s", result))
	}
}
