package main

import (
	"github.com/DrSmithFr/go-console"
)

func main() {

	//
	// Easy way to create a command with arguments and options
	//

	cmd := go_console.Script{
		Version:     "1.0.0",
		BuildFlag:   "2024-01-01",
		Name:        "app:command",
		Description: "The app:command command.",
	}

	cmd.Build()
}
