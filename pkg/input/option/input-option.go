package option

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/helper"
	"regexp"
	"strings"
)

const (
	Default = None

	None     = 1
	Required = 2
	Optional = 4
	List     = 8
)

const (
	Defined   = "true"
	Undefined = "false"
)

// constructor
func New(
	name string, // The option name
	mode int, // The option mode: One of the option constants (default: None)
) *InputOption {
	if "--" == name[:2] {
		name = name[2:]
	}

	if "" == name {
		panic(errors.New("an option name cannot be empty"))
	}

	if mode > 15 || mode < 1 {
		panic(errors.New(fmt.Sprintf("option mode '%d' is not valid", mode)))
	}

	opt := &InputOption{
		name:          name,
		shortcut:      "",
		mode:          mode,
		description:   "",
		defaultValue:  "",
		defaultValues: []string{},
	}

	if opt.IsList() && !opt.AcceptValue() {
		panic(errors.New("impossible to have an option mode List if the option does not accept a value"))
	}

	return opt
}

// Represents a command line option.
type InputOption struct {
	name          string
	shortcut      string
	mode          int
	defaultValue  string
	defaultValues []string
	description   string
}

// Returns the option name.
func (a *InputOption) GetName() string {
	return a.name
}

// set the option description
func (a *InputOption) SetDescription(desc string) *InputOption {
	a.description = desc
	return a
}

// Returns the description text.
func (a *InputOption) GetDescription() string {
	return a.description
}

// The shortcuts, can be empty, or a string of shortcuts delimited by '|'
func (a *InputOption) SetShortcut(shortcut string) *InputOption {
	if "" != shortcut {
		a := regexp.MustCompile(`(\|)-?`)
		shortcuts := a.Split(strings.TrimLeft(shortcut, "-"), -1)
		shortcut = strings.Join(shortcuts, "|")

		if "" == shortcut {
			panic(errors.New("an option shortcut cannot be empty"))
		}
	}

	a.shortcut = shortcut
	return a
}

// Returns the option shortcut.
func (a *InputOption) GetShortcut() string {
	return a.shortcut
}

// Returns true if the option accepts a value.
func (a *InputOption) AcceptValue() bool {
	return a.IsValueRequired() || a.IsValueOptional()
}

// returns true if the option requires a value.
func (a *InputOption) IsValueRequired() bool {
	return Required == (Required & a.mode)
}

// returns true if the option takes an optional value.
func (a *InputOption) IsValueNone() bool {
	return None == (None & a.mode)
}

// returns true if the option takes an optional value.
func (a *InputOption) IsValueOptional() bool {
	return Optional == (Optional & a.mode)
}

// returns true if the option takes an optional value.
func (a *InputOption) IsList() bool {
	return List == (List & a.mode)
}

// Sets the default value.
func (a *InputOption) SetDefault(defaultValue string) *InputOption {
	if !a.AcceptValue() && "" != defaultValue {
		panic(errors.New("cannot set a default value when using InputOption::VALUE_NONE mode"))
	}

	if a.IsList() {
		panic(errors.New("cannot use SetDefaultAnswer() for InputOption::VALUE_IS_ARRAY mode, use SetDefaults() instead"))
	}

	a.defaultValue = defaultValue
	return a
}

// Sets the default value for array options
func (a *InputOption) SetDefaults(values []string) *InputOption {
	if None == (None&a.mode) && 0 != len(values) {
		panic(errors.New("cannot set default values when using InputOption::VALUE_NONE mode"))
	}

	if !a.IsList() {
		panic(errors.New("cannot use SetDefaults() except for InputOption::List mode, use SetDefaultAnswer() instead"))
	}

	a.defaultValues = values

	return a
}

// Returns the default value.
func (a *InputOption) GetDefault() string {
	if a.IsList() {
		panic(errors.New("cannot use GetDefaultAnswer() for InputOption::List mode, use GetDefaults() instead"))
	}

	return a.defaultValue
}

// Returns the defaults value.
func (a *InputOption) GetDefaults() []string {
	if !a.IsList() {
		panic(errors.New("cannot use GetDefaults() except for InputOption::List, use GetDefaultAnswer() instead"))
	}

	return a.defaultValues
}

// compare to another option
func (a *InputOption) Equals(b InputOption) bool {
	if b.IsList() != a.IsList() {
		return false
	}

	if a.IsList() {
		return b.GetName() == a.GetName() &&
			b.GetShortcut() == a.GetShortcut() &&
			b.IsList() == a.IsList() &&
			b.IsValueRequired() == a.IsValueRequired() &&
			b.IsValueOptional() == a.IsValueOptional() &&
			helper.IsStringSliceEqual(b.GetDefaults(), a.GetDefaults())
	}

	return b.GetName() == a.GetName() &&
		b.GetShortcut() == a.GetShortcut() &&
		b.GetDefault() == a.GetDefault() &&
		b.IsList() == a.IsList() &&
		b.IsValueRequired() == a.IsValueRequired() &&
		b.IsValueOptional() == a.IsValueOptional()
}
