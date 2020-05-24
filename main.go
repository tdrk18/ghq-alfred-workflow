package main

import (
	"log"
	"os"
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
