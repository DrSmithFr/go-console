package main

import (
	"github.com/DrSmithFr/go-console/color"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/output"
)

func main() {
	// creating new output
	out := output.NewCliOutput(true, nil)

	// create new style
	s := formatter.NewOutputFormatterStyle(color.Red, color.Yellow, []string{color.Bold, color.Blink})

	// add style to formatter
	out.Formatter().SetStyle("fire", *s)

	// use the new style
	out.Println("<fire>foo</>")
}
