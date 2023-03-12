package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console"
)

func main() {

	script := go_console.Command{
		Description: "This Command act as a group of command.",
		Scripts: []*go_console.Script{
			{
				Name:        "app:user:create",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
			{
				Name:        "app:user:promote",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
			{
				Name:        "app:user:revoke",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
			{
				Name:        "cache:clear",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
			{
				Name:        "cache:clear",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
			{
				Name:        "database:create",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
			{
				Name:        "database:drop",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
			{
				Name:        "database:migration:migrate",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
			{
				Name:        "server:start",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
			{
				Name:        "server:stop",
				Description: "Print hello world form an external runner.",
				Runner:      displayCommandName,
			},
		},
	}

	// This start the command logic.
	script.Run()
	// There must have no code after this line as script.run() call os.Exit() on completion.
}

func displayCommandName(cmd *go_console.Script) go_console.ExitCode {
	cmd.PrintTitle(fmt.Sprintf("Command %s", cmd.Name))
	return go_console.ExitSuccess
}
