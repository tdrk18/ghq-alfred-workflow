package main

import (
	"os"
	"os/exec"
	"strings"
)

func execGhq() []byte {
	command := os.Getenv("ghq")
	out, err := exec.Command(command, "list", "-p").Output()
	if err != nil {
		return []byte{}
	}
	return out
}

func getRepos(bytes []byte) []string {
	trim := strings.Trim(string(bytes), "\n")
	return strings.Split(trim, "\n")
}
