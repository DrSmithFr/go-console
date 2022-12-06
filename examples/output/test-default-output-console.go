package main

import "DrSmithFr/go-console/pkg/output"

func main() {
	// creating new output
	out := output.NewConsoleOutput(true, nil)

	// white text on a red background
	out.Writeln("<error>An error</error>")

	// green text
	out.Writeln("<info>An information</info>")

	// yellow text
	out.Writeln("<comment>An comment</comment>")

	// black text on a cyan background
	out.Writeln("<question>A question</question>")

	// underscore text
	out.Writeln("<u>Some underscore text</u>")

	// bold text
	out.Writeln("<b>Some bold text</b>")
}
