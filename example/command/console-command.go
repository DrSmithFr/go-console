package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/input/argument"
	"time"
)

func main() {

	script := go_console.Command{
		UseNamespace: true,
		Description:  "This Command act as a group of command.",
		AppInfo: &go_console.ApplicationInfo{
			Name:      "app",
			Version:   "1.0.0",
			BuildDate: time.Now(),
		},
		Scripts: []*go_console.Script{
			{
				Name:        "app:user:create",
				Description: "Print hello world form an external runner.",
				Arguments: []go_console.Argument{
					{
						Name:        "email",
						Description: "The email of the user.",
						Value:       argument.Required,
					},
				},
				Runner: displayCommandName,
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
				Name:        "cache:warmup",
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
