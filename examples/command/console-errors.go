package main

import (
	"errors"
	"DrSmithFr/go-console/pkg/input/argument"
	"DrSmithFr/go-console/pkg/input/option"
	"DrSmithFr/go-console/pkg/style"
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
