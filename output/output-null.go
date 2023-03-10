package output

import (
	"fmt"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/verbosity"
)

// constructor
func NewNullOutput(decorated bool, format *formatter.OutputFormatter) *NullOutput {
	out := new(NullOutput)

	out.doPrint = out.Void

	if nil == format {
		out.formatter = formatter.NewOutputFormatter()
	} else {
		out.formatter = format
	}

	out.SetDecorated(decorated)

	return out
}

// Null output classes (~abstract)
type NullOutput struct {
	doPrint   func(string, verbosity.Level)
	doWrite   func([]byte) (int, error)
	formatter *formatter.OutputFormatter
	verbosity verbosity.Level
}

var _ OutputInterface = (*NullOutput)(nil)

func (o *NullOutput) Format(message string) string {
	if nil == o.formatter {
		return message
	}

	return (*o.formatter).Format(message)
}

func (o *NullOutput) preWriteEvent(message string) {

}

func (o *NullOutput) Void(message string, level verbosity.Level) {
	// do nothing
}

func (o *NullOutput) Print(message string) {
	o.doPrint(o.Format(message), verbosity.Normal)
}

// Writes a message to the output and adds a newline at the end
func (o *NullOutput) Println(message string) {
	o.Print(fmt.Sprintf("%s\n", message))
}

func (o *NullOutput) PrintOnVerbose(message string, level verbosity.Level) {
	o.doPrint(o.Format(message), level)
}

// Writes a message to the output and adds a newline at the end
func (o *NullOutput) PrintlnOnVerbose(message string, level verbosity.Level) {
	o.PrintOnVerbose(fmt.Sprintf("%s\n", message), level)
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
func (o *NullOutput) Formatter() *formatter.OutputFormatter {
	return o.formatter
}

func (o *NullOutput) SetVerbosity(verbosity verbosity.Level) {
	o.verbosity = verbosity
}

func (o *NullOutput) Verbosity() verbosity.Level {
	return o.verbosity
}

func (o *NullOutput) IsQuiet() bool {
	return o.Verbosity() == verbosity.Quiet
}

func (o *NullOutput) IsVerbose() bool {
	return o.Verbosity() == verbosity.Verbose
}

func (o *NullOutput) IsVeryVerbose() bool {
	return o.Verbosity() == verbosity.VeryVerbose
}

func (o *NullOutput) IsDebug() bool {
	return o.Verbosity() == verbosity.Debug
}

func (o *NullOutput) IsVerbosityAllowed(level verbosity.Level) bool {
	return level <= o.Verbosity()
}

func (o *NullOutput) Write(raw []byte) (n int, err error) {
	formatted := o.Format(string(raw))
	message := []byte(formatted)
	return o.doWrite(message)
}
