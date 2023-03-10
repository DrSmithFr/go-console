package output

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/verbosity"
)

// constructor
func NewBufferedOutput(decorated bool, format *formatter.OutputFormatter) *BufferedOutput {
	out := &BufferedOutput{
		buffer: "",
	}

	out.doPrint = out.Store
	out.doWrite = out.StoreBytes

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

var _ OutputInterface = (*BufferedOutput)(nil)

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

func (o *BufferedOutput) StoreBytes(p []byte) (n int, err error) {
	if o.IsQuiet() {
		return 0, errors.New("buffered output is quiet")
	}

	o.Store(string(p), verbosity.Normal)

	return len(p), nil
}
