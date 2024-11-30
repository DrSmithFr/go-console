package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"os"
)

func main() {
	cmd := go_console.NewScript().Build()
	qh := question.NewHelper(os.Stdin, cmd.Output)

	// Simple question with default answer
	name := qh.Ask(
		question.
			NewQuestion("What is your name?").
			SetDefaultAnswer("John Doe"),
	)

	cmd.PrintText("Hello " + name)
}
