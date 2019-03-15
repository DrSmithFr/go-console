package main

import (
	"github.com/MrSmith777/go-console/pkg/output"
	"github.com/MrSmith777/go-console/pkg/style"
)

func main() {
	// creating new output
	out := output.NewConsoleOutput(true, nil)

	// creating new styler
	styler := style.NewGoStyler(out)

	out.Writeln("<bg=red>boby</> est cool")
	styler.Title("boby est cool")
}
