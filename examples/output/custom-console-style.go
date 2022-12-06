package main

import (
	"DrSmithFr/go-console/pkg/color"
	"DrSmithFr/go-console/pkg/formatter"
	"DrSmithFr/go-console/pkg/output"
)

func main() {
	// creating new output
	out := output.NewConsoleOutput(true, nil)

	// create new style
	s := formatter.NewOutputFormatterStyle(color.RED, color.YELLOW, []string{color.BOLD, color.BLINK})

	// add style to formatter
	out.GetFormatter().SetStyle("fire", *s)

	// use the new style
	out.Writeln("<fire>foo</>")
}
