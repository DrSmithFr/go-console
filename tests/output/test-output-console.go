package main

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/color"
	formatter2 "github.com/MrSmith777/go-console/pkg/formatter"
	"github.com/MrSmith777/go-console/pkg/output"
)

func main() {
	out := output.NewConsoleOutput()
	out.Writeln("Ceci est un test")
	out.Write("Ceci est un test\n")

	formatter := out.GetFormatter()
	stack := formatter.GetStyleStack()
	stack.Push(formatter2.NewOutputFormatterStyle(color.DEFAULT, color.CYAN))

	// enable color
	out.SetDecorated(true)
	out.Writeln("PAPAPAPAPAPAPAPAPAPA")

	stack.Pop(nil)
	stack.Pop(nil)
	fmt.Printf("bob")

	out.Writeln("POPOPOPOPOPOPOPOPOPO")
}