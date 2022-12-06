package output

import (
	"fmt"
	"DrSmithFr/go-console/pkg/formatter"
)

// constructor
func NewConsoleOutput(decorated bool, format *formatter.OutputFormatter) *ConsoleOutput {
	out := new(ConsoleOutput)

	out.doWrite = out.Write

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

func (o *ConsoleOutput) Write(message string) {
	fmt.Printf(o.format(message))
}
