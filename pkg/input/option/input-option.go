package option

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/helper"
	"regexp"
	"strings"
)

const (
	NONE     = 1
	REQUIRED = 2
	OPTIONAL = 4
	IS_ARRAY = 8
)

// constructor
func New(
	name string, // The option name
	mode int,    // The option mode: One of the option constants (default: NONE)
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

	if opt.IsArray() && ! opt.AcceptValue() {
		panic(errors.New("impossible to have an option mode IS_ARRAY if the option does not accept a value"))
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
	return REQUIRED == (REQUIRED & a.mode)
}

// returns true if the option takes an optional value.
func (a *InputOption) IsValueOptional() bool {
	return OPTIONAL == (OPTIONAL & a.mode)
}

// returns true if the option takes an optional value.
func (a *InputOption) IsArray() bool {
	return IS_ARRAY == (IS_ARRAY & a.mode)
}

// Sets the default value.
func (a *InputOption) SetDefault(defaultValue string) *InputOption {
	if NONE == (NONE&a.mode) && "" != defaultValue {
		panic(errors.New("cannot set a default value when using InputOption::VALUE_NONE mode"))
	}

	if a.IsArray() {
		panic(errors.New("cannot use SetDefault() for InputOption::VALUE_IS_ARRAY mode, use SetDefaults() instead"))
	}

	a.defaultValue = defaultValue
	return a
}

func (a *InputOption) SetDefaults(values []string) *InputOption {
	if NONE == (NONE&a.mode) && 0 != len(values) {
		panic(errors.New("cannot set default values when using InputOption::VALUE_NONE mode"))
	}

	if ! a.IsArray() {
		panic(errors.New("cannot use SetDefaults() except for InputOption::IS_ARRAY mode, use SetDefault() instead"))
	}

	a.defaultValues = values

	return a
}

// Returns the default value.
func (a *InputOption) GetDefault() string {
	if a.IsArray() {
		panic(errors.New("cannot use GetDefault() for InputOption::IS_ARRAY mode, use GetDefaults() instead"))
	}

	return a.defaultValue
}

// Returns the defaults value.
func (a *InputOption) GetDefaults() []string {
	if ! a.IsArray() {
		panic(errors.New("cannot use GetDefaults() except for InputOption::IS_ARRAY, use GetDefault() instead"))
	}

	return a.defaultValues
}

func (a *InputOption) Equals(b InputOption) bool {
	if b.IsArray() != a.IsArray() {
		return false
	}

	if a.IsArray() {
		return b.GetName() == a.GetName() &&
			b.GetShortcut() == a.GetShortcut() &&
			b.IsArray() == a.IsArray() &&
			b.IsValueRequired() == a.IsValueRequired() &&
			b.IsValueOptional() == a.IsValueOptional() &&
			helper.IsStringSliceEqual(b.GetDefaults(), a.GetDefaults())
	}

	return b.GetName() == a.GetName() &&
		b.GetShortcut() == a.GetShortcut() &&
		b.GetDefault() == a.GetDefault() &&
		b.IsArray() == a.IsArray() &&
		b.IsValueRequired() == a.IsValueRequired() &&
		b.IsValueOptional() == a.IsValueOptional()
}
