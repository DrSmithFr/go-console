package definition

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/DrSmithFr/go-console/input/option"
	"strings"
)

// constructor
func New() *InputDefinition {
	def := &InputDefinition{
		arguments: map[string]argument.InputArgument{},
		options:   map[string]option.InputOption{},

		argumentKeysOrdered: []string{},
		optionKeysOrdered:   []string{},

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

	argumentKeysOrdered []string
	optionKeysOrdered   []string

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
	if _, ok := i.arguments[arg.Name()]; ok {
		panic(errors.New(fmt.Sprintf("an argument with name '%s' already exists", arg.Name())))
	}

	if i.hasAnArrayArgument {
		panic(errors.New("cannot add an argument after an array argument"))
	}

	if arg.IsRequired() && i.hasOptional {
		panic(errors.New("cannot add a required argument after an optional one"))
	}

	if arg.IsList() {
		i.hasAnArrayArgument = true
	}

	if arg.IsRequired() {
		i.requiredCount = i.requiredCount + 1
	} else {
		i.hasOptional = true
	}

	i.arguments[arg.Name()] = arg
	i.argumentKeysOrdered = append(i.argumentKeysOrdered, arg.Name())

	return i
}

// Returns true if an InputArgument object exists by name or position.
func (i *InputDefinition) HasArgument(name string) bool {
	_, found := i.arguments[name]
	return found
}

// Returns an InputArgument by name or by position
func (i *InputDefinition) Argument(name string) *argument.InputArgument {
	if !i.HasArgument(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	if arg, ok := i.arguments[name]; ok {
		return &arg
	}

	return nil
}

// Gets the array of InputArgument objects
func (i *InputDefinition) Arguments() map[string]argument.InputArgument {
	return i.arguments
}

// Gets the array of InputArgument keys ordered
func (i *InputDefinition) ArgumentsOrder() []string {
	return i.argumentKeysOrdered
}

// Returns the number of InputArguments.
func (i *InputDefinition) ArgumentCount() int {
	return len(i.arguments)
}

// Returns the number of required InputArguments.
func (i *InputDefinition) ArgumentRequiredCount() int {
	return i.requiredCount
}

// Gets the default values.
func (i *InputDefinition) ArgumentDefaults() map[string][]string {
	values := make(map[string][]string)

	for _, key := range i.argumentKeysOrdered {
		arg := i.Argument(key)

		if arg.IsList() {
			values[arg.Name()] = arg.Defaults()
		} else {
			if "" != arg.Default() {
				values[arg.Name()] = []string{arg.Default()}
			} else {
				values[arg.Name()] = nil
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
	if _, ok := i.options[opt.Name()]; ok {
		panic(errors.New(fmt.Sprintf("an option named '%s' already exists", opt.Name())))
	}

	if "" != opt.Shortcut() {
		for _, shortcut := range strings.Split(opt.Shortcut(), "|") {
			if i.HasShortcut(shortcut) && !opt.Equals(*i.FindOptionForShortcut(shortcut)) {
				panic(errors.New(fmt.Sprintf("An option with shortcut '%s' already exists", shortcut)))
			}
		}
	}

	i.options[opt.Name()] = opt
	i.optionKeysOrdered = append(i.optionKeysOrdered, opt.Name())

	if "" != opt.Shortcut() {
		for _, shortcut := range strings.Split(opt.Shortcut(), "|") {
			i.shortcuts[shortcut] = opt.Name()
		}
	}

	return i
}

// Returns true if an InputArgument object exists by name or position.
func (i *InputDefinition) HasOption(name string) bool {
	_, found := i.options[name]
	return found
}

// Returns an InputOption by name
func (i *InputDefinition) Option(name string) *option.InputOption {
	if !i.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '%s' argument does not exist", name)))
	}

	if arg, ok := i.options[name]; ok {
		return &arg
	}

	return nil
}

// Gets the array of InputOption objects
func (i *InputDefinition) Options() map[string]option.InputOption {
	return i.options
}

// Gets the array of InputOption keys ordered
func (i *InputDefinition) OptionsOrder() []string {
	return i.optionKeysOrdered
}

// returns true if an InputOption object exists by shortcut.
func (i *InputDefinition) HasShortcut(s string) bool {
	_, found := i.shortcuts[s]
	return found
}

// Gets an InputOption by shortcut.
func (i *InputDefinition) FindOptionForShortcut(s string) *option.InputOption {
	return i.Option(i.ShortcutToName(s))
}

// Gets the default values.
func (i *InputDefinition) OptionDefaults() map[string][]string {
	values := make(map[string][]string)

	for _, key := range i.optionKeysOrdered {
		opt := i.Option(key)

		if opt.IsList() {
			values[opt.Name()] = opt.Defaults()
		} else {
			if "" != opt.Default() {
				values[opt.Name()] = []string{opt.Default()}
			} else {
				values[opt.Name()] = []string{}
			}
		}
	}

	return values
}

// Returns the InputOption name given a shortcut.
func (i *InputDefinition) ShortcutToName(s string) string {
	opt, found := i.shortcuts[s]

	if !found {
		panic(errors.New(fmt.Sprintf("the '-%s' option does not exist", s)))
	}

	return opt
}

// Returns the InputOption name given a shortcut.
func (i *InputDefinition) Synopsis(short bool) string {
	var elements []string

	if short && 0 != len(i.Options()) {
		elements = append(elements, "[options]")
	} else if !short {
		for _, key := range i.optionKeysOrdered {
			opt := i.Option(key)
			value := ""
			start := ""
			end := ""

			if opt.IsValueOptional() {
				start = "["
				end = "]"
			}

			if opt.IsAcceptValue() {
				value = fmt.Sprintf(
					" %s%s%s",
					start,
					strings.ToUpper(opt.Name()),
					end,
				)
			}

			shortcut := ""

			if "" != opt.Shortcut() {
				shortcut = fmt.Sprintf("-%s|", opt.Shortcut())
			}

			elements = append(
				elements,
				fmt.Sprintf(
					"[%s--%s%s]",
					shortcut,
					opt.Name(),
					value,
				),
			)
		}
	}

	if 0 != len(elements) && 0 != len(i.Arguments()) {
		elements = append(elements, "[--]")
	}

	tail := ""

	for _, key := range i.argumentKeysOrdered {
		arg := i.Argument(key)
		element := fmt.Sprintf("<%s>", arg.Name())

		if arg.IsList() {
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
