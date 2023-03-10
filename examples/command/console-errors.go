package main

import (
	"errors"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/DrSmithFr/go-console/pkg/style"
)

func main() {
	io := style.
		NewConsoleCommand().
		AddInputArgument(
			argument.
				New("name", argument.Required),
		).
		AddInputOption(
			option.
				New("foo", option.None).
				SetShortcut("f"),
		).
		Build()

	// enable stylish errors
	defer io.HandleRuntimeException()

	panic(errors.New("runtime error !"))
}
