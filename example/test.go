package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/DrSmithFr/go-console/input/option"
)

func main() {

	// Declare the command with a struct
	cmd1 := go_console.Cli{
		Args: []go_console.Arg{
			{
				Name: "name",
				Mode: argument.Required,
			},
		},
		Opts: []go_console.Opt{
			{
				Name:     "foo",
				Shortcut: "f",
				Mode:     option.None,
			},
		},
	}
	// Build the command before using it
	cmd1.Build()

	// Declare the command with fluent interface
	cmd2 := go_console.
		NewCli().
		AddInputArgument(
			argument.
				New("name", argument.Required).
				SetDescription("Who do you want to greet?"),
		).
		AddInputArgument(
			argument.
				New("last_name", argument.Optional).
				SetDescription("Your last name?"),
		).Build()

	//
	// You now have access to a last_name argument in your command:
	//

	text := fmt.Sprintf("Hi %s", cmd1.Input().Argument("name"))

	lastName := cmd2.Input().Argument("last_name")

	if lastName != "" {
		text = fmt.Sprintf("%s %s", text, lastName)
	}

	cmd1.Output().Write(text)
	cmd2.Output().Writeln("!")
}
