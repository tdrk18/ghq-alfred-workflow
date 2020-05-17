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
			query := strings.Join(c.Args().Slice(), ",")
			wf.NewItem("First result!").Valid(true)
			wf.NewItem("Second result!").Valid(true)
			wf.NewItem(query).Valid(true)
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
