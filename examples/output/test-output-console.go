package main

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/color"
	"github.com/MrSmith777/go-console/pkg/formatter"
	"github.com/MrSmith777/go-console/pkg/output"
)

func main() {
	// creating new output
	out := output.NewConsoleOutput(true, nil)

	// enjoy
	out.Writeln("<error>An error</error>")
	out.Writeln("<error>An error</>")

	out.Writeln("<info>An information</info>")
	out.Writeln("<info>An information</>")

	out.Writeln("<comment>An comment</comment>")
	out.Writeln("<comment>An comment</>")

	out.Writeln("<question>A question</question>")
	out.Writeln("<question>A question</>")

	out.Writeln(
		fmt.Sprintf(
			"<bg=%s;fg=%s;options=%s>custom style testing</>",
			color.BLUE,
			color.GREEN,
			color.BOLD,
		),
	)

	s := formatter.NewOutputFormatterStyle(color.RED, color.YELLOW, []string{color.BOLD, color.BLINK})
	out.GetFormatter().SetStyle("fire", *s)

	out.Writeln("<fire>foo</>")
}
