package output

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/formatter"
)

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

type ConsoleOutput struct {
	NullOutput
}

func (o *ConsoleOutput) Write(message string) {
	fmt.Printf(o.format(message))
}