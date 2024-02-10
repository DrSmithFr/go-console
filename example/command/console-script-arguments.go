package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/DrSmithFr/go-console/input/option"
	"time"
)

func main() {

	//
	// Easy way to create a command with arguments and options
	//

	cmd := go_console.Script{
		Version:     "1.0.0",
		BuildDate:   time.Now(),
		Name:        "app:command",
		Description: "The app:command command.",
		Arguments: []go_console.Argument{
			{
				Name:        "name",
				Value:       argument.Required | argument.List,
				Description: "The name of the user.",
			},
		},
		Options: []go_console.Option{
			{
				Name:        "foo",
				Shortcut:    "f",
				Value:       option.None,
				Description: "The foo option.",
			},
		},
	}

	cmd.Build()
}
