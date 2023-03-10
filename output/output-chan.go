package output

import (
	"errors"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/verbosity"
)

// constructor
func NewChanOutput(channel chan string, decorated bool, format *formatter.OutputFormatter) *ChanOutput {
	out := &ChanOutput{
		channel: channel,
	}

	out.doPrint = out.Send
	out.doWrite = out.SendBytes

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

var _ OutputInterface = (*ChanOutput)(nil)

func (o *ChanOutput) Send(message string, level verbosity.Level) {
	if o.IsQuiet() {
		return
	}

	if o.IsVerbosityAllowed(level) {
		o.channel <- message
	}
}

func (o *ChanOutput) SendBytes(p []byte) (n int, err error) {
	if o.IsQuiet() {
		return 0, errors.New("chat output is quiet")
	}

	o.channel <- string(p)

	return len(p), nil
}
