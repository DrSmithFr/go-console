package main

import (
	"github.com/DrSmithFr/go-console/pkg/style"
)

func main() {
	// creating default console styler
	io := style.NewConsoleGoStyler()

	// according to my terminal size (default: 120)
	io.SetMaxLineLength(80)

	io.Listing([]string{
		"Lorem Ipsum Dolor Sit Amet",
		"Lorem Ipsum Dolor Sit Amet",
		"Lorem Ipsum Dolor Sit Amet",
	})
}
