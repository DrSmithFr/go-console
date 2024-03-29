package input

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/input/definition"
	"github.com/DrSmithFr/go-console/input/option"
)

type abstractInput struct {
	definition definition.InputDefinition

	interactive bool

	arguments      map[string]string
	argumentArrays map[string][]string

	options      map[string]string
	optionArrays map[string][]string

	doParse    func()
	doValidate func()
}

// get the input definition
func (i *abstractInput) Definition() *definition.InputDefinition {
	return &i.definition
}

// get all parsed arguments
func (i *abstractInput) Arguments() map[string]string {
	return i.arguments
}

// get all parsed arguments array
func (i *abstractInput) ArgumentArrays() map[string][]string {
	return i.argumentArrays
}

// Returns true if an InputArgument object exists by name or position
func (i *abstractInput) HasArgument(name string) bool {
	return i.definition.HasArgument(name)
}

// Returns the argument value for a given argument name
func (i *abstractInput) Argument(name string) string {
	if !i.definition.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	arg := i.definition.Argument(name)

	if arg.IsList() {
		panic(errors.New(fmt.Sprintf("the '%s' argument is an array, use ArgumentList() instead", name)))
	}

	if val, ok := i.arguments[name]; ok {
		return val
	}

	return arg.Default()
}

// Returns the argument array value for a given argument name
func (i *abstractInput) ArgumentList(name string) []string {
	if !i.definition.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	arg := i.definition.Argument(name)

	if !arg.IsList() {
		panic(errors.New(fmt.Sprintf("the '%s' argument is not an array, use Argument() instead", name)))
	}

	if val, ok := i.argumentArrays[name]; ok {
		return val
	}

	return arg.Defaults()
}

// Sets an argument value by name
func (i *abstractInput) SetArgument(name string, value string) {
	if !i.definition.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	arg := i.definition.Argument(name)

	if arg.IsList() {
		panic(errors.New(fmt.Sprintf("the '%s' argument is an array, use SetArgumentList() instead", name)))
	}

	i.arguments[name] = value
}

// Sets an argument array value by name
func (i *abstractInput) SetArgumentList(name string, value []string) {
	if !i.definition.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	arg := i.definition.Argument(name)

	if !arg.IsList() {
		panic(errors.New(fmt.Sprintf("the '%s' argument is not an array, use SetArgument() instead", name)))
	}

	i.argumentArrays[name] = value
}

// Returns all the given options merged with the default values
func (i *abstractInput) Options() map[string]string {
	return i.options
}

// Returns all the given options array merged with the default values
func (i *abstractInput) OptionLists() map[string][]string {
	return i.optionArrays
}

// Returns true if an InputOption object exists by name
func (i *abstractInput) HasOption(name string) bool {
	return i.definition.HasOption(name)
}

// Returns the option value for a given option name
func (i *abstractInput) Option(name string) string {
	if !i.definition.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' option does not exist", name)))
	}

	opt := i.definition.Option(name)

	if opt.IsList() {
		panic(errors.New(fmt.Sprintf("the '%s' option is an array, use OptionList() instead", name)))
	}

	if val, ok := i.options[name]; ok {
		return val
	}

	// TODO find a better way to handle option.None
	if !opt.IsAcceptValue() {
		return option.Undefined
	}

	return opt.Default()
}

// Returns the option array value for a given option name
func (i *abstractInput) OptionList(name string) []string {
	if !i.definition.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' option does not exist", name)))
	}

	opt := i.definition.Option(name)

	if !opt.IsList() {
		panic(errors.New(fmt.Sprintf("the '%s' option is not an array, use Option() instead", name)))
	}

	if val, ok := i.optionArrays[name]; ok {
		return val
	}

	return opt.Defaults()
}

// Sets an option value by name
func (i *abstractInput) SetOption(name string, value string) {
	if !i.definition.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' option does not exist", name)))
	}

	opt := i.definition.Option(name)

	if opt.IsList() {
		panic(errors.New(fmt.Sprintf("the '%s' option is an array, use SetOptionList() instead", name)))
	}

	i.options[name] = value
}

// Sets an option array value by name
func (i *abstractInput) SetOptionList(name string, value []string) {
	if !i.definition.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' option does not exist", name)))
	}

	opt := i.definition.Option(name)

	if !opt.IsList() {
		panic(errors.New(fmt.Sprintf("the '%s' option is not an array, use SetOption() instead", name)))
	}

	i.optionArrays[name] = value
}

// Is this input means interactive?
func (i *abstractInput) IsInteractive() bool {
	return i.interactive
}

// Sets the input interactivity
func (i *abstractInput) SetInteractive(interactive bool) {
	i.interactive = interactive
}

func (i *abstractInput) initialize() {
	i.options = make(map[string]string)
	i.arguments = make(map[string]string)

	i.optionArrays = make(map[string][]string)
	i.argumentArrays = make(map[string][]string)
}

// Binds the current input instance with the given arguments and options
func (i *abstractInput) Bind(def definition.InputDefinition) {
	i.initialize()

	i.definition = def

	i.Parse()
}

// Processes command line arguments
func (i *abstractInput) Parse() {
	i.doParse()
}

// Validates the input
func (i *abstractInput) Validate() {
	i.doValidate()
}
