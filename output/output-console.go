package output

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/verbosity"
	"os"
)

// constructor
func NewCliOutput(decorated bool, format *formatter.OutputFormatter) *ConsoleOutput {
	out := new(ConsoleOutput)

	out.doPrint = out.StdOut
	out.doWrite = out.StdOutBytes

	if nil == format {
		out.formatter = formatter.NewOutputFormatter()
	} else {
		out.formatter = format
	}

	out.SetDecorated(decorated)

	return out
}

// Console output classes
type ConsoleOutput struct {
	NullOutput
}

var _ OutputInterface = (*ConsoleOutput)(nil)

func (o *ConsoleOutput) StdOut(message string, level verbosity.Level) {
	if o.IsQuiet() {
		return
	}

	if o.IsVerbosityAllowed(level) {
		fmt.Printf(message)
	}
}

func (o *ConsoleOutput) StdOutBytes(p []byte) (n int, err error) {
	if o.IsQuiet() {
		return 0, errors.New("console output is quiet")
	}

	return fmt.Fprint(os.Stdout, string(p))
}
