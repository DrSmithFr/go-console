package output

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/formatter"
)

// constructor
func NewNullOutput(decorated bool, format *formatter.OutputFormatter) *NullOutput {
	out := new(NullOutput)

	out.doWrite = out.Write

	if nil == format {
		out.formatter = formatter.NewOutputFormatter()
	} else {
		out.formatter = format
	}

	out.SetDecorated(decorated)

	return out
}

// Null output classes (eq. toi abstract)
type NullOutput struct {
	doWrite   func(string)
	formatter *formatter.OutputFormatter
}

func (o *NullOutput) format(message string) string {
	if nil == o.formatter {
		return message
	}

	return (*o.formatter).Format(message)
}

func (o *NullOutput) Write(message string) {
	// do nothing
}

// Writes a message to the output and adds a newline at the end
func (o *NullOutput) Writeln(message string) {
	o.doWrite(fmt.Sprintf("%s\n", message))
}

// Sets the decorated flag
func (o *NullOutput) SetDecorated(decorated bool) {
	if nil == o.formatter {
		return
	}

	(*o.formatter).SetDecorated(decorated)
}

// Gets the decorated flag
func (o *NullOutput) IsDecorated() bool {
	if nil == o.formatter {
		return false
	}

	return (*o.formatter).IsDecorated()
}

// Set current output formatter instance
func (o *NullOutput) SetFormatter(formatter *formatter.OutputFormatter) {
	o.formatter = formatter
}

// Returns current output formatter instance
func (o *NullOutput) GetFormatter() *formatter.OutputFormatter {
	return o.formatter
}
