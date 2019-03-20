package output

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/formatter"
)

func NewBufferedOutput(decorated bool, format *formatter.OutputFormatter) *BufferedOutput {
	out := new(BufferedOutput)

	out.buffer = ""
	out.doWrite = out.Write

	if nil == format {
		out.formatter = formatter.NewOutputFormatter()
	} else {
		out.formatter = format
	}

	out.SetDecorated(decorated)

	return out
}

type BufferedOutput struct {
	NullOutput
	buffer string
}

func (o *BufferedOutput) Write(message string) {
	o.buffer = fmt.Sprintf("%s%s", o.buffer, message)
}

// Empties buffer and returns its content.
func (o *BufferedOutput) Fetch() string {
	buffer := o.buffer
	o.buffer = ""
	return buffer
}
