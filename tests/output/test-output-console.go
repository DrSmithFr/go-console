package main

import (
	"github.com/MrSmith777/go-console/pkg/output"
)

func main() {
	// creating new output
	out := output.NewConsoleOutput()

	// enable color
	out.SetDecorated(true)

	// enjoy
	out.Writeln("<error>An error</error>")
	out.Writeln("<error>An error</>")

	out.Writeln("<info>An information</info>")
	out.Writeln("<info>An information</>")

	out.Writeln("<comment>An comment</comment>")
	out.Writeln("<comment>An comment</>")

	out.Writeln("<question>A question</question>")
	out.Writeln("<question>A question</>")
}