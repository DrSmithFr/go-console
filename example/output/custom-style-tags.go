package main

import (
	"github.com/DrSmithFr/go-console/output"
)

func main() {
	// green text
	out := output.NewCliOutput(true, nil)

	// black text on a cyan background
	out.Println("<fg=green>foo</>")

	// green text
	out.Println("<fg=black;bg=cyan>foo</>")

	// bold text on a yellow background
	out.Println("<bg=yellow;options=bold>foo</>")

	// bold text with underscore
	out.Println("<options=bold,underscore>foo</>")
}
