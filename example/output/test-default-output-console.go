package main

import (
	"github.com/DrSmithFr/go-console/output"
)

func main() {
	// creating new output
	out := output.NewCliOutput(true, nil)

	// white text on a red background
	out.Println("<error>An error</error>")

	// green text
	out.Println("<info>An information</info>")

	// yellow text
	out.Println("<comment>An comment</comment>")

	// black text on a cyan background
	out.Println("<question>A question</question>")

	// underscore text
	out.Println("<u>Some underscore text</u>")

	// bold text
	out.Println("<b>Some bold text</b>")
}
