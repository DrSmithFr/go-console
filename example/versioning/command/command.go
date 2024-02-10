package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console"
)

var (
	Name    string
	Version string
	Flag    string
)

func main() {

	script := go_console.Command{
		UseNamespace: true,
		Description:  "This Command act as a group of command.",
		Info: &go_console.Info{
			Name:      Name,
			Version:   Version,
			BuildFlag: Flag,
		},
		Scripts: []*go_console.Script{
			{
				Name:        "hello:world",
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
