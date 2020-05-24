package main

import (
	"github.com/urfave/cli/v2"
	"strings"
)

func createApp() *cli.App {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			query := getQuery(c.Args().First())
			repos := getRepos(execGhq())
			for _, repo := range repos {
				addItem(repo)
			}
			filter(query)
			warnEmpty()
			sendFeedback()
			return nil
		},
	}
	return app
}

func getQuery(arg string) string {
	return strings.Trim(arg, "\n")
}
