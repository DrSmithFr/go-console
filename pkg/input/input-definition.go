package input

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/option"
)

func NewInputDefinition(arguments []argument.InputArgument, options []option.InputOption) *InputDefinition {
	def := &InputDefinition{
		arguments: map[string]argument.InputArgument{},
		options:   map[string]option.InputOption{},

		requiredCount: -1,

		hasOptional:        false,
		hasAnArrayArgument: false,

		shortcuts: []string{},
	}

	def.AddArguments(arguments)

	return def
}

// A InputDefinition represents a set of valid command line arguments and options.
type InputDefinition struct {
	arguments map[string]argument.InputArgument
	options   map[string]option.InputOption

	requiredCount int

	hasOptional        bool
	hasAnArrayArgument bool

	shortcuts []string
}

// Sets the InputArgument objects.
func (i *InputDefinition) SetArguments(arguments []argument.InputArgument) {
	i.arguments = map[string]argument.InputArgument{}
	i.requiredCount = 0
	i.hasOptional = false
	i.hasAnArrayArgument = false
	i.AddArguments(arguments)
}

// Adds an array of InputArgument objects.
func (i *InputDefinition) AddArguments(arguments []argument.InputArgument) {
	for _, arg := range arguments {
		i.AddArgument(arg)
	}
}

// panic when incorrect argument is given
func (i *InputDefinition) AddArgument(arg argument.InputArgument) {
	if _, ok := i.arguments[arg.GetName()]; ok {
		panic(errors.New(fmt.Sprintf("an argument with name '%s' already exists", arg.GetName())))
	}

	if i.hasAnArrayArgument {
		panic(errors.New("cannot add an argument after an array argument"))
	}

	if arg.IsRequired() && i.hasOptional {
		panic(errors.New("cannot add a required argument after an optional one"))
	}

	if arg.IsArray() {
		i.hasAnArrayArgument = true
	}

	if arg.IsRequired() {
		i.requiredCount += 1
	} else {
		i.hasOptional = true
	}

	i.arguments[arg.GetName()] = arg
}

// Returns true if an InputArgument object exists by name or position.
func (i *InputDefinition) HasArgument(name string) bool {
	_, found := i.arguments[name]
	return found
}

func (i *InputDefinition) GetArgument(name string) *argument.InputArgument {
	if ! i.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	if arg, ok := i.arguments[name]; ok {
		return &arg
	}

	return nil
}

func (i *InputDefinition) GetArguments() []argument.InputArgument {
	var args []argument.InputArgument

	for _, val := range i.arguments {
		args = append(args, val)
	}

	return args
}

// Returns the number of InputArguments.
func (i *InputDefinition) GetArgumentCount() int {
	return len(i.arguments)
}

// Returns the number of required InputArguments.
func (i *InputDefinition) GetArgumentRequiredCount() int {
	return i.requiredCount
}

// Gets the default values.
func (i *InputDefinition) getArgumentDefaults() map[string]string {
	values := make(map[string]string)

	for _, arg := range i.arguments {
		values[arg.GetName()] = arg.GetDefault()
	}

	return values
}

// sets the InputOption objects.
func (i *InputDefinition) SetOptions(options []option.InputOption) {
	i.options = make(map[string]option.InputOption)
	i.shortcuts = []string{}
	i.AddOptions(options)
}

// Adds an array of InputOption objects.
func (i *InputDefinition) AddOptions(options []option.InputOption) {
	for _, opt := range options {
		i.AddOption(opt)
	}
}

// panic when option given already exist
func (i *InputDefinition) AddOption(opt option.InputOption) {
	if _, ok := i.arguments[opt.GetName()]; ok {
		panic(errors.New(fmt.Sprintf("an option named '%s' already exists", opt.GetName())))
	}
}
