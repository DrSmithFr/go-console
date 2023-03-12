package main

import (
	"github.com/DrSmithFr/go-console"
)

func main() {
	// creating default console styler
	cmd := go_console.NewScript()

	// according to my terminal size (default: 120)
	cmd.SetMaxLineLength(80)

	// use simple strings for short messages
	cmd.PrintTitle("Lorem ipsum dolor sit amet")
}
