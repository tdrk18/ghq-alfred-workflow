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
			wf.NewItem("First result!").Valid(true)
			wf.NewItem("Second result!").Valid(true)
			wf.NewItem(query).Valid(true)
			for _, repo := range repos {
				wf.NewItem(repo).Valid(true)
			}
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

func execGhq() []byte {
	command := os.Getenv("ghq")
	out, err := exec.Command(command, "list", "-p").Output()
	if err != nil {
		return []byte{}
	}
	return out
}
