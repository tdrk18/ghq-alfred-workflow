package main

import (
	aw "github.com/deanishe/awgo"
)

var (
	wf *aw.Workflow
)

func init() {
	wf = aw.New()
}

func run() {
	wf.NewItem("First result!").Valid(true)
	wf.NewItem("Second result!").Valid(true)
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
