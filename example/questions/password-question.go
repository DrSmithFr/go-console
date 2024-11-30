package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"os"
)

func main() {
	cmd := go_console.NewScript().Build()
	qh := question.NewHelper(os.Stdin, cmd.Output)

	// Simple question with hidden answer
	pass := qh.Ask(
		question.
			NewQuestion("What is your password?").
			SetHidden(true),
	)

	cmd.PrintText("Password: " + pass)
}
