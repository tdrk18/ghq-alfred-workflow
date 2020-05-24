package main

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"testing"
)

func TestGetQuery(t *testing.T) {
	args := "text\n"
	result := getQuery(args)
	if result != "text" {
		t.Fatal(fmt.Sprintf("failed: getQuery() returns %s", result))
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

func TestGetArgWithModifier(t *testing.T) {
	path := "/path/to/github.com/tdrk18/repo"
	resultCmd := getArgWithModifier(path, aw.ModCmd)
	if resultCmd != "/path/to/github.com/tdrk18/repo" {
		t.Fatal(fmt.Sprintf("failed: getArgWithModifier() returns %s", resultCmd))
	}
	resultCtrl := getArgWithModifier(path, aw.ModCtrl)
	if resultCtrl != "/path/to/github.com/tdrk18/repo" {
		t.Fatal(fmt.Sprintf("failed: getArgWithModifier() returns %s", resultCtrl))
	}
	resultFn := getArgWithModifier(path, aw.ModFn)
	if resultFn != "tdrk18/repo" {
		t.Fatal(fmt.Sprintf("failed: getArgWithModifier() returns %s", resultFn))
	}
	resultOpt := getArgWithModifier(path, aw.ModOpt)
	if resultOpt != "/path/to/github.com/tdrk18/repo" {
		t.Fatal(fmt.Sprintf("failed: getArgWithModifier() returns %s", resultOpt))
	}
	resultShift := getArgWithModifier(path, aw.ModShift)
	if resultShift != "https://github.com/tdrk18/repo" {
		t.Fatal(fmt.Sprintf("failed: getArgWithModifier() returns %s", resultShift))
	}
	resultDefault := getArgWithModifier(path, "")
	if resultDefault != "https://github.com/tdrk18/repo" {
		t.Fatal(fmt.Sprintf("failed: getArgWithModifier() returns %s", resultDefault))
	}
}

func TestGetSubWithModifier(t *testing.T) {
	resultCmd := getSubWithModifier(aw.ModCmd)
	if resultCmd != "Reveal in Finder." {
		t.Fatal(fmt.Sprintf("failed: getSubWithModifier() returns %s", resultCmd))
	}
	resultCtrl := getSubWithModifier(aw.ModCtrl)
	if resultCtrl != "Browse in terminal." {
		t.Fatal(fmt.Sprintf("failed: getSubWithModifier() returns %s", resultCtrl))
	}
	resultFn := getSubWithModifier(aw.ModFn)
	if resultFn != "Search in browser." {
		t.Fatal(fmt.Sprintf("failed: getSubWithModifier() returns %s", resultFn))
	}
	resultOpt := getSubWithModifier(aw.ModOpt)
	if resultOpt != "Open files in VSCode." {
		t.Fatal(fmt.Sprintf("failed: getSubWithModifier() returns %s", resultOpt))
	}
	resultShift := getSubWithModifier(aw.ModShift)
	if resultShift != "Open URL." {
		t.Fatal(fmt.Sprintf("failed: getSubWithModifier() returns %s", resultShift))
	}
	resultDefault := getSubWithModifier("")
	if resultDefault != "Open URL." {
		t.Fatal(fmt.Sprintf("failed: getSubWithModifier() returns %s", resultDefault))
	}
}
