package main

import (
	"github.com/DrSmithFr/go-console"
)

func main() {

	//
	// Easy way to create a command with arguments and options
	//

	cmd := go_console.Script{
		Name:        "app:command",
		Description: "The app:command command.",

		BuildInfo: &go_console.BuildInfo{
			Version:   "1.0.0",
			BuildFlag: "2024-01-01",
		},
	}

	cmd.Build()
}
