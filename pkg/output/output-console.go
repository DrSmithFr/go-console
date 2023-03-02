package output

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/formatter"
	"github.com/DrSmithFr/go-console/pkg/verbosity"
)

// constructor
func NewConsoleOutput(decorated bool, format *formatter.OutputFormatter) *ConsoleOutput {
	out := new(ConsoleOutput)

	out.doWrite = out.Print

	if nil == format {
		out.formatter = formatter.NewOutputFormatter()
	} else {
		out.formatter = format
	}

	out.SetDecorated(decorated)

	return out
}

// Console output classes
type ConsoleOutput struct {
	NullOutput
}

func (o *ConsoleOutput) Print(message string, level verbosity.Level) {
	if o.IsQuiet() {
		return
	}

	if o.IsVerbosityAllowed(level) {
		fmt.Printf(message)
	}
}
