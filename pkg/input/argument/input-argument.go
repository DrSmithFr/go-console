package argument

import (
	"errors"
	"fmt"
)

const (
	REQUIRED = 1
	OPTIONAL = 2
	IS_ARRAY = 4
)

func NewInputArgument(name string, mode int, description string, defaultValue string) *InputArgument {
	if mode > 7 || mode < 1 {
		panic(errors.New(fmt.Sprintf("Argument mode '%d' is not valid.", mode)))
	}

	arg := &InputArgument{
		name:        name,
		mode:        mode,
		description: description,
	}

	arg.SetDefault(defaultValue)

	return arg
}

// Represents a command line argument.
type InputArgument struct {
	name         string
	mode         int
	defaultValue string
	description  string
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
func (a *InputArgument) SetDefault(defaultValue string) {
	if REQUIRED == a.mode && "" != defaultValue {
		panic(errors.New("cannot set a default value except for InputArgument::OPTIONAL mode"))
	}

	if a.IsArray() {
		// TODO find a way to enable array default values
	}

	a.defaultValue = defaultValue
}

// Returns the default value.
func (a *InputArgument) GetDefault() string {
	return a.defaultValue
}

// Returns the description text
func (a *InputArgument) GetDescription() string {
	return a.description
}
