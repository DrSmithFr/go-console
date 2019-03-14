package output

import "fmt"

func NewBufferedOutput() BufferedOutput {
	out := new(BufferedOutput)
	out.buffer = ""
	out.doWrite = out.Write
	return *out
}

type BufferedOutput struct {
	NullOutput
	buffer string
}

func (o *BufferedOutput) Write(message string) {
	o.buffer = fmt.Sprintf("%s%s", o.buffer, message)
}

// Empties buffer and returns its content.
func (o* BufferedOutput) Fetch() string {
	buffer := o.buffer
	o.buffer = ""
	return buffer
}