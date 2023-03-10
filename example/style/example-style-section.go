package main

import (
	"github.com/DrSmithFr/go-console"
)

func main() {
	// creating default console styler
	cmd := go_console.NewCli()

	// according to my terminal size (default: 120)
	cmd.SetMaxLineLength(80)

	// use simple strings for short messages
	cmd.PrintSection("Lorem ipsum dolor sit amet")
}
