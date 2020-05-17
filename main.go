package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/urfave/cli/v2"
)

var (
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

func run() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			query := getQuery(c.Args().First())
			repos := getRepos(execGhq())
			for _, repo := range repos {
				addItem(repo)
			}
			filter(query)
			wf.WarnEmpty("No matched repository", "Please try new query")
			wf.SendFeedback()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	wf.Run(run)
}

func getQuery(arg string) string {
	return strings.Trim(arg, "\n")
}

func getRepos(bytes []byte) []string {
	trim := strings.Trim(string(bytes), "\n")
	return strings.Split(trim, "\n")
}

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

func execGhq() []byte {
	command := os.Getenv("ghq")
	out, err := exec.Command(command, "list", "-p").Output()
	if err != nil {
		return []byte{}
	}
	return out
}

func filter(query string) {
	wf.Filter(query)
}

func addItem(repo string) {
	wf.NewItem(getRepoName(repo, true)).
		Valid(true)
}
