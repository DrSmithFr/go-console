package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/verbosity"
)

func main() {
	io := go_console.NewCli().Build()
	out := io.Output()

	// this message is always displayed when the command is run without --quiet
	io.PrintTitle("Example of console verbosity")

	if io.IsQuiet() {
		// force a message to be displayed in quiet mode
		fmt.Println("command run in quiet mode")
	}

	// this message is always displayed when the command is run without --quiet
	io.PrintText("This message has normal verbosity (using io)")
	out.Writeln("This message has normal verbosity (using io.Output)")

	// this message displayed when the command is run --verbose or -v
	if io.IsVerbose() {
		io.PrintText("This message is verbose (using io)")
	}
	out.WritelnOnVerbose("This message is verbose (using io.Output)", verbosity.Verbose)

	// this message displayed when the command is run --very-verbose or -vv
	if io.IsVeryVerbose() {
		io.PrintText("This message is IsVeryVerbose (using io)")
	}
	out.WritelnOnVerbose("This message is IsVeryVerbose (using io.Output)", verbosity.VeryVerbose)

	// this message displayed when the command is run --debug or -vvv
	if io.IsDebug() {
		io.PrintText("This message is debug (using io)")
	}
	out.WritelnOnVerbose("This message is debug (using io.Output)", verbosity.Debug)

}
