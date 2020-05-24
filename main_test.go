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
