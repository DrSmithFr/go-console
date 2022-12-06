package output

import "DrSmithFr/go-console/pkg/formatter"

// constructor
func NewChanOutput(channel chan string, decorated bool, format *formatter.OutputFormatter) *ChanOutput {
	out := &ChanOutput{
		channel: channel,
	}

	out.doWrite = out.Write

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

func (o *ChanOutput) Write(message string) {
	o.channel <- o.format(message)
}
