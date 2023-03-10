package output

import (
	"fmt"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/verbosity"
)

// constructor
func NewBufferedOutput(decorated bool, format *formatter.OutputFormatter) *BufferedOutput {
	out := &BufferedOutput{
		buffer: "",
	}

	out.doWrite = out.Store

	if nil == format {
		out.formatter = formatter.NewOutputFormatter()
	} else {
		out.formatter = format
	}

	out.SetDecorated(decorated)

	return out
}

// Buffered output classes
type BufferedOutput struct {
	NullOutput
	buffer string
}

func (o *BufferedOutput) Store(message string, level verbosity.Level) {
	if o.IsQuiet() {
		return
	}

	if o.IsVerbosityAllowed(level) {
		o.buffer = fmt.Sprintf("%s%s", o.buffer, message)
	}
}

// Empties buffer and returns its content.
func (o *BufferedOutput) Fetch() string {
	buffer := o.buffer
	o.buffer = ""
	return buffer
}
