package output

import (
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/verbosity"
)

// OutputInterface is the interface implemented by all output classes
type OutputInterface interface {
	// Formats a message according to the current formatter styles.
	Format(message string) string

	// Writes a message to the output.
	Print(message string)

	// Writes a message to the output and adds a newline at the end.
	Println(message string)

	// Writes a message to the output.
	PrintOnVerbose(message string, verbosity verbosity.Level)

	// Writes a message to the output and adds a newline at the end.
	PrintlnOnVerbose(message string, verbosity verbosity.Level)

	// Sets the decorated flag.
	SetDecorated(decorated bool)

	// Gets the decorated flag.
	IsDecorated() bool

	// Sets current output formatter instance.
	SetFormatter(formatter *formatter.OutputFormatter)

	// Gets current output formatter instance.
	Formatter() *formatter.OutputFormatter

	// Sets the verbosity of the output.
	SetVerbosity(verbosity verbosity.Level)

	// Gets the current verbosity of the output.
	Verbosity() verbosity.Level

	// Returns whether verbosity is quiet (-q)
	IsQuiet() bool

	// Returns whether verbosity is verbose (-v)
	IsVerbose() bool

	// Returns whether verbosity is very verbose (-vv)
	IsVeryVerbose() bool

	// Returns whether verbosity is debug (-vvv)
	IsDebug() bool

	// Implements io.Writer
	Write(p []byte) (n int, err error)
}
