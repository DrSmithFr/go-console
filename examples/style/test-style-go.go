package main

import (
	"github.com/MrSmith777/go-console/pkg/style"
)

func main() {
	// creating default console styler
	io := style.NewConsoleGoStyler()

	// use Go styler
	io.Title("Lorem Ipsum Dolor Sit Amet")

	// access OutputInterface
	io.GetOutput().Write("<info>some info</>")
}
