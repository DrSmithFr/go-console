package output

import "github.com/MrSmith777/go-console/pkg/formatter"

func NewChanOutput(channel chan string) *ChanOutput {
	out := new(ChanOutput)

	out.channel = channel
	out.doWrite = out.Write

	out.formatter = formatter.NewOutputFormatter()

	return out
}

type ChanOutput struct {
	NullOutput
	channel chan string
}

func (o *ChanOutput) Write(message string) {
	o.channel <- o.format(message)
}