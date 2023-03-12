package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/verbosity"
)

func main() {
	cmd := go_console.NewScript().Build()
	out := cmd.output()

	// this message is always displayed when the command is displayCommandName without --quiet
	cmd.PrintTitle("Example of console verbosity")

	if cmd.IsQuiet() {
		// force a message to be displayed in quiet mode
		fmt.Println("command displayCommandName in quiet mode")
	}

	out.Println("<info>This message has displayed by <b>output.Println()</b>")
	cmd.PrintText("<info>This message has displayed by <b>go_console.PrintText()</b>")
	fmt.Fprintln(out, "<info>This message</info> using <b>Fprintln with output</b>")
	fmt.Fprintln(cmd, "<info>This message</info> using <b>Fprintln with go_console</b>")

	// this message displayed when the command is displayCommandName --verbose or -v
	if cmd.IsVerbose() {
		cmd.PrintText("This message is verbose (using io)")
	}
	out.PrintlnOnVerbose("This message is verbose (using cmd.output)", verbosity.Verbose)

	// this message displayed when the command is displayCommandName --very-verbose or -vv
	if cmd.IsVeryVerbose() {
		cmd.PrintText("This message is IsVeryVerbose (using io)")
	}
	out.PrintlnOnVerbose("This message is IsVeryVerbose (using cmd.output)", verbosity.VeryVerbose)

	// this message displayed when the command is displayCommandName --debug or -vvv
	if cmd.IsDebug() {
		cmd.PrintText("This message is debug (using io)")
	}
	out.PrintlnOnVerbose("This message is debug (using cmd.output)", verbosity.Debug)

}
