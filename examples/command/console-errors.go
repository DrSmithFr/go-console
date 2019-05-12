package main

import (
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/DrSmithFr/go-console/pkg/style"
	"github.com/pkg/errors"
)

func main() {
	io := style.
		NewConsoleStyler().
		AddInputArgument(
			argument.
				New("name", argument.REQUIRED),
		).
		AddInputOption(
			option.
				New("foo", option.NONE).
				SetShortcut("f"),
		).
		ParseInput().
		ValidateInput()

	// enable stylish errors
	defer io.HandleRuntimeException()

	panic(errors.New("runtime error !"))
}