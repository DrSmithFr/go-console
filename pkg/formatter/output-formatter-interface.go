package formatter

type OutputFormatterInterface interface {
	// Sets the decorated flag.
	SetDecorated(decorated bool)

	// Gets the decorated flag.
	IsDecorated() bool

	// Sets a new style.
	SetStyle(name string, styleInterface OutputFormatterStyleInterface)

	// Gets style options from style with specified name.
	GetStyle(name string) *OutputFormatterStyleInterface

	// Checks if output formatter has style with specified name.
	HasStyle(name string) bool

	// Formats a message according to the given styles.
	Format(message string) string
}
