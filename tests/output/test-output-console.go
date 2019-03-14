package main

import (
	"github.com/MrSmith777/go-console/pkg/color"
	Formatter "github.com/MrSmith777/go-console/pkg/formatter"
	"github.com/MrSmith777/go-console/pkg/output"
)

func main() {
	out := output.NewConsoleOutput()
	out.Writeln("Ceci est un test")
	out.Write("Ceci est un test\n")

	formatter := out.GetFormatter()
	stack := formatter.GetStyleStack()
	style := Formatter.NewOutputFormatterStyle(color.DEFAULT, color.CYAN)
	stack.Push(style)

	// enable color
	out.SetDecorated(true)
	out.Writeln("PAPAPAPAPAPAPAPAPAPA")

	style.SetOption(color.BLINK)

	out.Writeln("POPOPOPOPOPOPOPOPOPO")
}