package input

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"strings"
)

// constructor
func NewInputDefinition() *InputDefinition {
	def := &InputDefinition{
		arguments: map[string]argument.InputArgument{},
		options:   map[string]option.InputOption{},

		requiredCount: 0,

		hasOptional:        false,
		hasAnArrayArgument: false,

		shortcuts: map[string]string{},
	}

	return def
}

// A InputDefinition represents a set of valid command line arguments and options.
type InputDefinition struct {
	arguments map[string]argument.InputArgument
	options   map[string]option.InputOption

	requiredCount int

	hasOptional        bool
	hasAnArrayArgument bool

	shortcuts map[string]string
}

// Sets the InputArgument objects.
func (i *InputDefinition) SetArguments(arguments []argument.InputArgument) *InputDefinition {
	i.arguments = map[string]argument.InputArgument{}
	i.requiredCount = 0
	i.hasOptional = false
	i.hasAnArrayArgument = false
	i.AddArguments(arguments)
	return i
}

// Adds an array of InputArgument objects.
func (i *InputDefinition) AddArguments(arguments []argument.InputArgument) *InputDefinition {
	for _, arg := range arguments {
		i.AddArgument(arg)
	}

	return i
}

// panic when incorrect argument is given
func (i *InputDefinition) AddArgument(arg argument.InputArgument) *InputDefinition {
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
		i.requiredCount = i.requiredCount + 1
	} else {
		i.hasOptional = true
	}

	i.arguments[arg.GetName()] = arg

	return i
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

func (i *InputDefinition) GetArguments() map[string]argument.InputArgument {
	return i.arguments
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
func (i *InputDefinition) GetArgumentDefaults() map[string][]string {
	values := make(map[string][]string)

	for _, arg := range i.arguments {
		if arg.IsArray() {
			values[arg.GetName()] = arg.GetDefaults()
		} else {
			if "" != arg.GetDefault() {
				values[arg.GetName()] = []string{arg.GetDefault()}
			} else {
				values[arg.GetName()] = []string{}
			}
		}

	}

	return values
}

// sets the InputOption objects.
func (i *InputDefinition) SetOptions(options []option.InputOption) *InputDefinition {
	i.options = make(map[string]option.InputOption)
	i.shortcuts = make(map[string]string)
	i.AddOptions(options)
	return i
}

// Adds an array of InputOption objects.
func (i *InputDefinition) AddOptions(options []option.InputOption) *InputDefinition {
	for _, opt := range options {
		i.AddOption(opt)
	}

	return i
}

// panic when option given already exist
func (i *InputDefinition) AddOption(opt option.InputOption) *InputDefinition {
	if _, ok := i.options[opt.GetName()]; ok {
		panic(errors.New(fmt.Sprintf("an option named '%s' already exists", opt.GetName())))
	}

	if "" != opt.GetShortcut() {
		for _, shortcut := range strings.Split(opt.GetShortcut(), "|") {
			if i.HasShortcut(shortcut) && ! opt.Equals(*i.GetOptionForShortcut(shortcut)) {
				panic(errors.New(fmt.Sprintf("An option with shortcut '%s' already exists", shortcut)))
			}
		}
	}

	i.options[opt.GetName()] = opt

	if "" != opt.GetShortcut() {
		for _, shortcut := range strings.Split(opt.GetShortcut(), "|") {
			i.shortcuts[shortcut] = opt.GetName()
		}
	}

	return i
}

// Returns true if an InputArgument object exists by name or position.
func (i *InputDefinition) HasOption(name string) bool {
	_, found := i.options[name]
	return found
}

func (i *InputDefinition) GetOption(name string) *option.InputOption {
	if ! i.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	if arg, ok := i.options[name]; ok {
		return &arg
	}

	return nil
}

func (i *InputDefinition) GetOptions() map[string]option.InputOption {
	return i.options
}

// returns true if an InputOption object exists by shortcut.
func (i *InputDefinition) HasShortcut(s string) bool {
	_, found := i.shortcuts[s]
	return found
}

// Gets an InputOption by shortcut.
func (i *InputDefinition) GetOptionForShortcut(s string) *option.InputOption {
	return i.GetOption(i.ShortcutToName(s))
}

// Gets the default values.
func (i *InputDefinition) GetOptionDefaults() map[string][]string {
	values := make(map[string][]string)

	for _, opt := range i.options {
		if opt.IsArray() {
			values[opt.GetName()] = opt.GetDefaults()
		} else {
			if "" != opt.GetDefault() {
				values[opt.GetName()] = []string{opt.GetDefault()}
			} else {
				values[opt.GetName()] = []string{}
			}
		}

	}

	return values
}

// Returns the InputOption name given a shortcut.
func (i *InputDefinition) ShortcutToName(s string) string {
	opt, found := i.shortcuts[s]

	if ! found {
		panic(errors.New(fmt.Sprintf("the '-%s' option does not exist", s)))
	}

	return opt
}

// Returns the InputOption name given a shortcut.
func (i *InputDefinition) GetSynopsis(short bool) string {
	var elements []string

	if short && 0 != len(i.GetOptions()) {
		elements = append(elements, "[options]")
	} else if !short {
		for _, opt := range i.options {
			value := ""

			start := ""
			end := ""

			if opt.IsValueOptional() {
				start = "["
				end = "]"
			}

			if opt.AcceptValue() {
				value = fmt.Sprintf(
					" %s%s%s",
					start,
					strings.ToUpper(opt.GetName()),
					end,
				)
			}

			shortcut := ""

			if "" != opt.GetShortcut() {
				shortcut = fmt.Sprintf("-s|", opt.GetShortcut())
			}

			elements = append(
				elements,
				fmt.Sprintf(
					"[%s--%s%s]",
					shortcut,
					opt.GetName(),
					value,
				),
			)
		}
	}

	if 0 != len(elements) && 0 != len(i.GetArguments()) {
		elements = append(elements, "[--]")
	}

	tail := ""

	for _, arg := range i.arguments {
		element := fmt.Sprintf("<%s>", arg.GetName())

		if arg.IsArray() {
			element = fmt.Sprintf("%s...", element)
		}

		if !arg.IsRequired() {
			element = fmt.Sprintf("[%s", element)
			tail = fmt.Sprintf("%s]", tail)
		}

		elements = append(elements, element)
	}

	return fmt.Sprintf(
		"%s%s",
		strings.Join(elements, " "),
		tail,
	)
}
