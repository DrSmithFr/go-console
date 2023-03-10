package main

import "github.com/DrSmithFr/go-console"

func main() {
	// creating default console styler
	cmd := go_console.NewCli()

	// according to my terminal size (default: 120)
	cmd.SetMaxLineLength(80)

	cmd.PrintListing([]string{
		"Lorem Ipsum Dolor Sit Amet",
		"Lorem Ipsum Dolor Sit Amet",
		"Lorem Ipsum Dolor Sit Amet",
	})
}
