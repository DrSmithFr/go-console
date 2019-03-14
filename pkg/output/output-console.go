package output

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/formatter"
)

func NewConsoleOutput() *ConsoleOutput {
	out := new(ConsoleOutput)

	out.doWrite = out.Write

	out.formatter = formatter.NewOutputFormatter()

	return out
}

type ConsoleOutput struct {
	NullOutput
}

func (o *ConsoleOutput) Write(message string) {
	fmt.Printf(o.format(message))
}