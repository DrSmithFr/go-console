package output

import (
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/verbosity"
)

// constructor
func NewChanOutput(channel chan string, decorated bool, format *formatter.OutputFormatter) *ChanOutput {
	out := &ChanOutput{
		channel: channel,
	}

	out.doWrite = out.Send

	if nil == format {
		out.formatter = formatter.NewOutputFormatter()
	} else {
		out.formatter = format
	}

	out.SetDecorated(decorated)

	return out
}

// Chan output classes
type ChanOutput struct {
	NullOutput
	channel chan string
}

func (o *ChanOutput) Send(message string, level verbosity.Level) {
	if o.IsQuiet() {
		return
	}

	if o.IsVerbosityAllowed(level) {
		o.channel <- message
	}
}
