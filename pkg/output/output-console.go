package output

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/formatter"
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

func (o *ConsoleOutput) Print(message string) {
	fmt.Printf(message)
}
