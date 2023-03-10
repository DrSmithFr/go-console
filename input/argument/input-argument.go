package argument

import (
	"errors"
	"fmt"
)

const (
	Default  = Optional
	Required = 1
	Optional = 2
	List     = 4
)

// constructor
func New(
	name string, // The argument name
	mode int, // The argument mode: Required or Optional (default: Optional)
) *InputArgument {
	if mode > 7 || mode < 1 {
		panic(errors.New(fmt.Sprintf("Argument mode '%d' is not valid.", mode)))
	}

	arg := &InputArgument{
		name:          name,
		mode:          mode,
		description:   "",
		defaultValue:  "",
		defaultValues: []string{},
	}

	return arg
}

// Represents a command line argument.
type InputArgument struct {
	name          string
	mode          int
	defaultValue  string
	defaultValues []string
	description   string
}

// Returns the argument name.
func (a *InputArgument) Name() string {
	return a.name
}

// Returns true if the argument is required.
func (a *InputArgument) IsRequired() bool {
	return Required == (Required & a.mode)
}

// Returns true if the argument can take multiple values.
func (a *InputArgument) IsList() bool {
	return List == (List & a.mode)
}

// Sets the default value.
func (a *InputArgument) SetDefault(defaultValue string) *InputArgument {
	if Required == a.mode && "" != defaultValue {
		panic(errors.New("cannot set a default value except for InputArgument::Optional mode"))
	}

	if a.IsList() {
		panic(errors.New("cannot use SetDefaultAnswer() for InputArgument::List mode, use SetDefaults() instead"))
	}

	a.defaultValue = defaultValue
	return a
}

// Sets the default value for array args.
func (a *InputArgument) SetDefaults(values []string) *InputArgument {
	if Required == a.mode && 0 != len(values) {
		panic(errors.New("cannot set a default value except for InputArgument::Optional mode"))
	}

	if !a.IsList() {
		panic(errors.New("cannot use SetDefaults() except for InputArgument::List mode, use SetDefaultAnswer() instead"))
	}

	a.defaultValues = values

	return a
}

// Returns the default value.
func (a *InputArgument) Default() string {
	if a.IsList() {
		panic(errors.New("cannot use GetDefaultAnswer() for InputArgument::List mode, use Defaults() instead"))
	}

	return a.defaultValue
}

// Returns the defaults value.
func (a *InputArgument) Defaults() []string {
	if !a.IsList() {
		panic(errors.New("cannot use Defaults() except for InputArgument::List, use GetDefaultAnswer() instead"))
	}

	return a.defaultValues
}

// Returns the description text
func (a *InputArgument) Description() string {
	return a.description
}

// Returns the description text
func (a *InputArgument) SetDescription(desc string) *InputArgument {
	a.description = desc
	return a
}
