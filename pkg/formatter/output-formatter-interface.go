package formatter

// Formatter interface for console output
type OutputFormatterInterface interface {
	// Sets the decorated flag.
	SetDecorated(decorated bool)

	// Gets the decorated flag.
	IsDecorated() bool

	// Sets a new style to cache.
	SetStyle(name string, style OutputFormatterStyle)

	// Gets style from cache with specified name.
	GetStyle(name string) *OutputFormatterStyle

	// Gets style stack
	GetStyleStack() *OutputFormatterStyleStack

	// Checks if output formatter has style in cache with specified name.
	HasStyle(name string) bool

	// Formats a message according to the given styles.
	Format(message string) string
}
