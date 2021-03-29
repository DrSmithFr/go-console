package argument

import (
	"errors"
	"fmt"
)

const (
	DEFAULT = OPTIONAL

	REQUIRED = 1
	OPTIONAL = 2
	IS_ARRAY = 4
)

// constructor
func New(
	name string, // The argument name
	mode int,    // The argument mode: REQUIRED or OPTIONAL (default: OPTIONAL)
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
func (a *InputArgument) GetName() string {
	return a.name
}

// Returns true if the argument is required.
func (a *InputArgument) IsRequired() bool {
	return REQUIRED == (REQUIRED & a.mode)
}

// Returns true if the argument can take multiple values.
func (a *InputArgument) IsArray() bool {
	return IS_ARRAY == (IS_ARRAY & a.mode)
}

// Sets the default value.
func (a *InputArgument) SetDefault(defaultValue string) *InputArgument {
	if REQUIRED == a.mode && "" != defaultValue {
		panic(errors.New("cannot set a default value except for InputArgument::OPTIONAL mode"))
	}

	if a.IsArray() {
		panic(errors.New("cannot use SetDefault() for InputArgument::IS_ARRAY mode, use SetDefaults() instead"))
	}

	a.defaultValue = defaultValue
	return a
}

// Sets the default value for array args.
func (a *InputArgument) SetDefaults(values []string) *InputArgument {
	if REQUIRED == a.mode && 0 != len(values) {
		panic(errors.New("cannot set a default value except for InputArgument::OPTIONAL mode"))
	}

	if !a.IsArray() {
		panic(errors.New("cannot use SetDefaults() except for InputArgument::IS_ARRAY mode, use SetDefault() instead"))
	}

	a.defaultValues = values

	return a
}

// Returns the default value.
func (a *InputArgument) GetDefault() string {
	if a.IsArray() {
		panic(errors.New("cannot use GetDefault() for InputArgument::IS_ARRAY mode, use GetDefaults() instead"))
	}

	return a.defaultValue
}

// Returns the defaults value.
func (a *InputArgument) GetDefaults() []string {
	if !a.IsArray() {
		panic(errors.New("cannot use GetDefaults() except for InputArgument::IS_ARRAY, use GetDefault() instead"))
	}

	return a.defaultValues
}

// Returns the description text
func (a *InputArgument) GetDescription() string {
	return a.description
}

// Returns the description text
func (a *InputArgument) SetDescription(desc string) *InputArgument {
	a.description = desc
	return a
}
