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

	// examples of working call

	// ./script bob --foo --bar --val="myValue"
	// ./script bob --foo -b -v "myValue"
	// ./script bob --foo -bmyValue -vothervale
	// ./script bob --bar=myValue -fvlol

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
		AddInputOption(
			option.
				New("bar", option.OPTIONAL).
				SetShortcut("b"),
		).
		AddInputOption(
			option.
				New("val", option.REQUIRED).
				SetShortcut("v"),
		).
		ParseInput().
		ValidateInput()

	//
	// Do what you want with console and args !
	//

	fmt.Printf(
		"name argument value '%s'\n",
		io.GetInput().GetArgument("name"),
	)

	if option.DEFINED == io.GetInput().GetOption("foo") {
		fmt.Printf("foo option is set\n")
	} else {
		fmt.Printf("foo option not used\n")
	}

	fmt.Printf(
		"bar option value '%s'\n",
		io.GetInput().GetOption("bar"),
	)

	fmt.Printf(
		"val option value '%s'\n",
		io.GetInput().GetOption("val"),
	)
}
