package output

import "fmt"

func NewConsoleOutput() ConsoleOutput {
	out := new(ConsoleOutput)
	out.doWrite = out.Write
	return *out
}

type ConsoleOutput struct {
	NullOutput
}

func (o *ConsoleOutput) Write(message string) {
	fmt.Print(message)
}