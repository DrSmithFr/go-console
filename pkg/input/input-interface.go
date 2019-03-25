package input

// InputInterface is the interface implemented by all input classes.
type InputInterface interface {
	// Returns the first argument from the raw parameters (not parsed).
	GetFirstArgument() string

	// Returns true if the raw parameters (not parsed) contain a value.
	//
	// This method is to be used to introspect the input parameters
	// before they have been validated. It must be used carefully.
	// Does not necessarily return the correct result for short options
	// when multiple flags are combined in the same option.
	HasParameterOption(values []string, onlyParams bool)

	// Returns the value of a raw option (not parsed).
	//
	// This method is to be used to introspect the input parameters
	// before they have been validated. It must be used carefully.
	// Does not necessarily return the correct result for short options
	// when multiple flags are combined in the same option.
	GetParameterOption(values []string, defaultValue string, onlyParams bool)

	// Binds the current Input instance with the given arguments and options.
	Bind(definition InputDefinition)

	// Validates the input.
	Validate()

	// Returns all the given arguments merged with the default values.
	GetArguments() []string

	// Returns the argument value for a given argument name.
	GetArgument(name string) []string

	// Returns the argument value for a given argument name.
	SetArgument(name string, values []string)

	// Returns the argument value for a given argument name.
	HasArgument(name string)

	// Returns all the given options merged with the default values.
	GetOptions() []string

	// Returns the option value for a given option name.
	GetOption(name string) []string

	// Sets an option value by name.
	SetOption(name string, values []string)

	// Returns true if an InputOption object exists by name.
	HasOption(name string)

	// Is this input means interactive?
	IsInteractive() bool

	// Sets the input interactivity.
	SetInteractive(bool)
}
