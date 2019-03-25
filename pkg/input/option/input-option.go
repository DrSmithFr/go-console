package option

import "errors"

const (
	VALUE_NONE = 1
	VALUE_REQUIRED = 2
	VALUE_OPTIONAL = 4
	VALUE_IS_ARRAY = 4
)

func NewInputOption(name string, shortcut []string, mode int, description string, defaultValue string) *InputOption {
	if "--" == name[2:] {
		name = name[2:]
	}

	if "" == name {
		panic(errors.New("an option name cannot be empty"))
	}

	if nil == shortcut {

	}

	opt := & InputOption{
		name:name,
		shortcut:shortcut,
		mode:mode,
		description:description,
		defaultValue:defaultValue,
	}

	return opt
}

type InputOption struct {
	name string
	shortcut []string
	mode int
	defaultValue string
	description string
}

// Returns the argument name.
func (a *InputOption) GetName() string {
	return a.name
}