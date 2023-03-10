package main

import (
	"github.com/DrSmithFr/go-console"
)

func main() {
	// creating default console styler
	io := go_console.NewCli()

	// according to my terminal size (default: 120)
	io.SetMaxLineLength(80)

	// use simple strings for short messages
	io.PrintSection("Lorem ipsum dolor sit amet")
}
