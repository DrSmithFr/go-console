package output

import (
	"fmt"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/verbosity"
)

// constructor
func NewNullOutput(decorated bool, format *formatter.OutputFormatter) *NullOutput {
	out := new(NullOutput)

	out.doWrite = out.Void

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
	doWrite   func(string, verbosity.Level)
	formatter *formatter.OutputFormatter
	verbosity verbosity.Level
}

func (o *NullOutput) format(message string) string {
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

func (o *NullOutput) Write(message string) {
	o.doWrite(o.format(message), verbosity.Normal)
}

// Writes a message to the output and adds a newline at the end
func (o *NullOutput) Writeln(message string) {
	o.Write(fmt.Sprintf("%s\n", message))
}

func (o *NullOutput) WriteOnVerbose(message string, level verbosity.Level) {
	o.doWrite(o.format(message), level)
}

// Writes a message to the output and adds a newline at the end
func (o *NullOutput) WritelnOnVerbose(message string, level verbosity.Level) {
	o.WriteOnVerbose(fmt.Sprintf("%s\n", message), level)
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
func (o *NullOutput) GetFormatter() *formatter.OutputFormatter {
	return o.formatter
}

func (o *NullOutput) SetVerbosity(verbosity verbosity.Level) {
	o.verbosity = verbosity
}

func (o *NullOutput) GetVerbosity() verbosity.Level {
	return o.verbosity
}

func (o *NullOutput) IsQuiet() bool {
	return o.GetVerbosity() == verbosity.Quiet
}

func (o *NullOutput) IsVerbose() bool {
	return o.GetVerbosity() == verbosity.Verbose
}

func (o *NullOutput) IsVeryVerbose() bool {
	return o.GetVerbosity() == verbosity.VeryVerbose
}

func (o *NullOutput) IsDebug() bool {
	return o.GetVerbosity() == verbosity.Debug
}

func (o *NullOutput) IsVerbosityAllowed(level verbosity.Level) bool {
	return level <= o.GetVerbosity()
}
