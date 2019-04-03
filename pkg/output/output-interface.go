package output

import "github.com/DrSmithFr/go-console/pkg/formatter"

// OutputInterface is the interface implemented by all Output classes
type OutputInterface interface {
	// Formats a message according to the current formatter styles.
	format(message string) string

	// Writes a message to the output.
	Write(message string)

	// Writes a message to the output and adds a newline at the end.
	Writeln(message string)

	// Sets the decorated flag.
	SetDecorated(decorated bool)

	// Gets the decorated flag.
	IsDecorated() bool

	// Sets current output formatter instance.
	SetFormatter(formatter *formatter.OutputFormatter)

	// Gets current output formatter instance.
	GetFormatter() *formatter.OutputFormatter
}
