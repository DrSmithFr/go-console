package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/verbosity"
)

func main() {
	cmd := go_console.NewCli().Build()
	out := cmd.Output()

	// this message is always displayed when the command is run without --quiet
	cmd.PrintTitle("Example of console verbosity")

	if cmd.IsQuiet() {
		// force a message to be displayed in quiet mode
		fmt.Println("command run in quiet mode")
	}

	out.Println("<info>This message has displayed by <b>Output.Println()</b>")
	cmd.PrintText("<info>This message has displayed by <b>go_console.PrintText()</b>")
	fmt.Fprintln(out, "<info>This message</info> using <b>Fprintln with Output</b>")
	fmt.Fprintln(cmd, "<info>This message</info> using <b>Fprintln with go_console</b>")

	// this message displayed when the command is run --verbose or -v
	if cmd.IsVerbose() {
		cmd.PrintText("This message is verbose (using io)")
	}
	out.PrintlnOnVerbose("This message is verbose (using cmd.Output)", verbosity.Verbose)

	// this message displayed when the command is run --very-verbose or -vv
	if cmd.IsVeryVerbose() {
		cmd.PrintText("This message is IsVeryVerbose (using io)")
	}
	out.PrintlnOnVerbose("This message is IsVeryVerbose (using cmd.Output)", verbosity.VeryVerbose)

	// this message displayed when the command is run --debug or -vvv
	if cmd.IsDebug() {
		cmd.PrintText("This message is debug (using io)")
	}
	out.PrintlnOnVerbose("This message is debug (using cmd.Output)", verbosity.Debug)

}
