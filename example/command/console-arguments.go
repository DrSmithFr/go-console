package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/DrSmithFr/go-console/input/option"
)

func main() {

	//
	// Easy way to create a command with arguments and options
	//

	cmd := go_console.Cli{
		Name: "app:command",
		Desc: "The app:command command.",
		Args: []go_console.Arg{
			{
				Name:        "name",
				Mode:        argument.Required,
				Description: "The name of the user.",
			},
		},
		Opts: []go_console.Opt{
			{
				Name:        "foo",
				Shortcut:    "f",
				Mode:        option.None,
				Description: "The foo option.",
			},
		},
	}

	cmd.Build()

}
