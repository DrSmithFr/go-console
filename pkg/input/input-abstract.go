package input

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/input/definition"
	"github.com/DrSmithFr/go-console/pkg/input/option"
)

type abstractInput struct {
	definition definition.InputDefinition

	interactive bool

	arguments      map[string]string
	argumentArrays map[string][]string

	options      map[string]string
	optionArrays map[string][]string

	doParse func()
}

func (i *abstractInput) GetDefinition() *definition.InputDefinition {
	return &i.definition
}

func (i *abstractInput) GetArguments() map[string]string {
	return i.arguments
}

func (i *abstractInput) GetArgumentArrays() map[string][]string {
	return i.argumentArrays
}

func (i *abstractInput) HasArgument(name string) bool {
	return i.definition.HasArgument(name)
}

func (i *abstractInput) GetArgument(name string) string {
	if !i.definition.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	arg := i.definition.GetArgument(name)

	if arg.IsArray() {
		panic(errors.New(fmt.Sprintf("the '%s' argument is an array, use GetArgumentArray() instead", name)))
	}

	if val, ok := i.arguments[name]; ok {
		return val
	}

	return arg.GetDefault()
}

func (i *abstractInput) GetArgumentArray(name string) []string {
	if !i.definition.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	arg := i.definition.GetArgument(name)

	if !arg.IsArray() {
		panic(errors.New(fmt.Sprintf("the '%s' argument is not an array, use GetArgument() instead", name)))
	}

	if val, ok := i.argumentArrays[name]; ok {
		return val
	}

	return arg.GetDefaults()
}

func (i *abstractInput) SetArgument(name string, value string) {
	if !i.definition.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	arg := i.definition.GetArgument(name)

	if arg.IsArray() {
		panic(errors.New(fmt.Sprintf("the '%s' argument is an array, use SetArgumentArray() instead", name)))
	}

	i.arguments[name] = value
}

func (i *abstractInput) SetArgumentArray(name string, value []string) {
	if !i.definition.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	arg := i.definition.GetArgument(name)

	if !arg.IsArray() {
		panic(errors.New(fmt.Sprintf("the '%s' argument is not an array, use SetArgument() instead", name)))
	}

	i.argumentArrays[name] = value
}

func (i *abstractInput) GetOptions() map[string]string {
	return i.options
}

func (i *abstractInput) GetOptionArrays() map[string][]string {
	return i.optionArrays
}

func (i *abstractInput) HasOption(name string) bool {
	return i.definition.HasOption(name)
}

func (i *abstractInput) GetOption(name string) string {
	if !i.definition.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' option does not exist", name)))
	}

	opt := i.definition.GetOption(name)

	if opt.IsArray() {
		panic(errors.New(fmt.Sprintf("the '%s' option is an array, use GetOptionArray() instead", name)))
	}

	if val, ok := i.options[name]; ok {
		return val
	}

	// TODO find a better way to handle option.NONE
	if ! opt.AcceptValue() {
		return option.UNDEFINED
	}

	return opt.GetDefault()
}

func (i *abstractInput) GetOptionArray(name string) []string {
	if !i.definition.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' option does not exist", name)))
	}

	opt := i.definition.GetOption(name)

	if !opt.IsArray() {
		panic(errors.New(fmt.Sprintf("the '%s' option is not an array, use GetOption() instead", name)))
	}

	if val, ok := i.optionArrays[name]; ok {
		return val
	}

	return opt.GetDefaults()
}

func (i *abstractInput) SetOption(name string, value string) {
	if !i.definition.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' option does not exist", name)))
	}

	opt := i.definition.GetOption(name)

	if opt.IsArray() {
		panic(errors.New(fmt.Sprintf("the '%s' option is an array, use SetOptionArray() instead", name)))
	}

	i.options[name] = value
}

func (i *abstractInput) SetOptionArray(name string, value []string) {
	if !i.definition.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' option does not exist", name)))
	}

	opt := i.definition.GetOption(name)

	if !opt.IsArray() {
		panic(errors.New(fmt.Sprintf("the '%s' option is not an array, use SetOption() instead", name)))
	}

	i.optionArrays[name] = value
}

func (i *abstractInput) IsInteractive() bool {
	return i.interactive
}

func (i *abstractInput) SetInteractive(interactive bool) {
	i.interactive = interactive
}

func (i *abstractInput) initialize() {
	i.options = make(map[string]string)
	i.arguments = make(map[string]string)

	i.optionArrays = make(map[string][]string)
	i.argumentArrays = make(map[string][]string)
}

func (i *abstractInput) Bind(def definition.InputDefinition) {
	i.initialize()

	i.definition = def

	i.Parse()
}

func (i *abstractInput) Parse() {
	i.doParse()
}

func (i *abstractInput) Validate()  {
	// TODO
}
