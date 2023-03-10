package main

import (
	"errors"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/DrSmithFr/go-console/input/option"
)

func main() {
	io := go_console.NewCli().
		AddInputArgument(
			argument.New("name", argument.Required),
		).
		AddInputOption(
			option.New("foo", option.None).
				SetShortcut("f"),
		).
		Build()

	// enable stylish errors
	defer io.HandleRuntimeException()

	panic(errors.New("runtime error !"))
}
