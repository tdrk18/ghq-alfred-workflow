package main

import (
	"fmt"
	"strings"
)

func getRepoName(repo string, excludeDomain bool) string {
	repoPath := strings.Split(repo, "/")
	length := len(repoPath)
	var i int
	if excludeDomain {
		i = 2
	} else {
		i = 3
	}
	return strings.Join(repoPath[length-i:length], "/")
}

func getDomain(repo string) string {
	repoPath := strings.Split(repo, "/")
	length := len(repoPath)
	return repoPath[length-3]
}

func getRepoURL(repo string) string {
	repoName := getRepoName(repo, false)
	return fmt.Sprintf("https://%s", repoName)
}

func getIconName(repo string) string {
	domain := getDomain(repo)
	switch {
	case strings.Contains(domain, "github"):
		return "github.png"
	case strings.Contains(domain, "bitbucket"):
		return "bitbucket.png"
	case strings.Contains(domain, "gitlab"):
		return "gitlab.png"
	default:
		return "git.png"
	}
}
