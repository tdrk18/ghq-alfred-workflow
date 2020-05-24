package main

import (
	"github.com/urfave/cli/v2"
	"strings"
)

func createApp() *cli.App {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			for _, repo := range getRepos(execGhq()) {
				addItem(repo)
			}
			filter(getQuery(c.Args().First()))
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
