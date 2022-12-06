package main

import (
	"DrSmithFr/go-console/pkg/style"
)

func main() {
	// creating default console styler
	io := style.NewConsoleStyler()

	// according to my terminal size (default: 120)
	io.SetMaxLineLength(80)

	// use simple strings for short messages
	io.Title("Lorem ipsum dolor sit amet")
}
