package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/DrSmithFr/go-console/pkg/style"
	"os"
)

func main() {
	fmt.Printf("%v\n", os.Args)

	io := style.
		NewConsoleStyler().
		AddInputArgument(
			argument.
				New("name", argument.REQUIRED),
		).
		AddInputOption(
			option.
				New("foo", option.OPTIONAL).
				SetShortcut("f"),
		).
		ParseInput().
		ValidateInput()

	fmt.Printf(
		"name argument value '%s'\n",
		io.GetInput().GetArgument("name"),
	)
}
