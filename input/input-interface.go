package input

import (
	"github.com/DrSmithFr/go-console/input/definition"
)

// InputInterface is the interface implemented by all input classes.
type InputInterface interface {
	// Returns the first argument from the raw parameters (not parsed).
	FirstArgument() string

	// Returns true if the raw parameters (not parsed) contain a value.
	//
	// This method is to be used to introspect the input parameters
	// before they have been validated. It must be used carefully.
	// Does not necessarily return the correct result for short options
	// when multiple flags are combined in the same option.
	HasParameterOption(values []string, onlyParams bool) bool

	// Returns the value of a raw option (not parsed).
	//
	// This method is to be used to introspect the input parameters
	// before they have been validated. It must be used carefully.
	// Does not necessarily return the correct result for short options
	// when multiple flags are combined in the same option.
	ParameterOption(values []string, defaultValue string, onlyParams bool)

	// Binds the current input instance with the given arguments and options.
	Bind(definition definition.InputDefinition)

	// Validates the input.
	Validate()

	// Parse the input data
	Parse()

	// Returns all the given arguments merged with the default values.
	Arguments() map[string]string

	// Returns all the given array arguments merged with the default values.
	ArgumentArrays() map[string][]string

	// Returns the argument value for a given argument name.
	Argument(name string) string

	// Returns the argument array value for a given array argument name.
	ArgumentList(name string) []string

	// Set the argument value for a given argument name.
	SetArgument(name string, value string)

	// Set the argument value for a given array argument name.
	SetArgumentList(name string, value []string)

	// Returns the argument value for a given argument name.
	HasArgument(name string) bool

	// Returns all the given options merged with the default values.
	Options() map[string]string

	// Returns all the given array options merged with the default values.
	OptionLists() map[string][]string

	// Returns the option value for a given option name.
	Option(name string) string

	// Returns the option array value for a given array option name.
	OptionList(name string) []string

	// Sets an option value by name.
	SetOption(name string, value string)

	// Sets an array option value by name.
	SetOptionList(name string, value []string)

	// Returns true if an InputOption object exists by name.
	HasOption(name string) bool

	// Is this input means interactive?
	IsInteractive() bool

	// Sets the input interactivity.
	SetInteractive(bool)

	// Get the input definition
	Definition() *definition.InputDefinition
}
