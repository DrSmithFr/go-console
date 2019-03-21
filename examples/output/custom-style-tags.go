package main

import "github.com/DrSmithFr/go-console/pkg/output"

func main() {
	// green text
	out := output.NewConsoleOutput(true, nil)

	// black text on a cyan background
	out.Writeln("<fg=green>foo</>")

	// green text
	out.Writeln("<fg=black;bg=cyan>foo</>")

	// bold text on a yellow background
	out.Writeln("<bg=yellow;options=bold>foo</>")

	// bold text with underscore
	out.Writeln("<options=bold,underscore>foo</>")
}