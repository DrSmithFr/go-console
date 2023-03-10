package input

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/helper"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/definition"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"os"
	"regexp"
	"strings"
)

// constructor
func NewArgvInput(argv []string) *ArgvInput {
	input := new(ArgvInput)

	if nil == argv {
		input.tokens = os.Args[1:]
	} else {
		input.tokens = argv[1:]
	}

	input.doParse = input.ParseArgv
	input.initialize()
	input.definition = *definition.New()

	return input
}

// ArgvInput represents an input coming from the CLI arguments
type ArgvInput struct {
	abstractInput
	tokens []string
	parsed []string
}

// Returns the first argument from the raw parameters (not parsed)
func (i *ArgvInput) GetFirstArgument() string {
	for _, token := range i.tokens {
		if "" != token && '-' == token[0] {
			continue
		}

		return token
	}

	panic(errors.New("first argument not found"))
}

// Returns true if the raw parameters (not parsed) contain a value
func (i *ArgvInput) HasParameterOption(values []string, onlyParams bool) bool {
	panic("implement me")
}

// Returns the value of a raw option (not parsed).
func (i *ArgvInput) GetParameterOption(values []string, defaultValue string, onlyParams bool) {
	panic("implement me")
}

// parse cli argv
func (i *ArgvInput) ParseArgv() {
	parseOptions := true
	i.parsed = i.tokens

	for {
		if 0 == len(i.parsed) {
			break
		}

		token := i.parsed[0]
		i.parsed = i.parsed[1:]

		if parseOptions && "" == token {
			i.parseArgument(token)
		} else if parseOptions && "--" == token {
			parseOptions = false
		} else if parseOptions && regexp.MustCompile("^--").MatchString(token) {
			i.parseLongOption(token)
		} else if parseOptions && '-' == token[0] && "-" != token {
			i.parseShortOption(token)
		} else {
			i.parseArgument(token)
		}

	}
}

//
// internal
//

func (i *ArgvInput) parseShortOption(token string) {
	name := token[1:]

	if len(name) > 1 {
		// allow long shortcut with None value
		if i.definition.HasShortcut(name) && i.definition.GetOptionForShortcut(name).IsValueNone() {
			i.addShortOption(name, "")
			return
		}

		shortcut := name[0:1]

		if i.definition.HasShortcut(shortcut) && i.definition.GetOptionForShortcut(shortcut).AcceptValue() {
			// an option with a value (with no space)
			i.addShortOption(shortcut, name[1:])
		} else {
			i.parseShortOptionSet(name)
		}

		return
	}

	i.addShortOption(name, "")
}

func (i *ArgvInput) parseShortOptionSet(name string) {
	length := len(name)

	for index := 0; index < length; index++ {
		shortcut := name[index : index+1]

		if !i.definition.HasShortcut(shortcut) {
			panic(errors.New(fmt.Sprintf("the '-%s' option does not exist", shortcut)))
		}

		opt := i.definition.GetOptionForShortcut(shortcut)

		if opt.AcceptValue() {
			if index == length-1 {
				i.addLongOption(opt.GetName(), "")
			} else {
				i.addLongOption(opt.GetName(), name[index+1:])
			}

			break
		} else {
			i.addLongOption(opt.GetName(), "")
		}
	}
}

func (i *ArgvInput) parseLongOption(token string) {
	name := token[2:]
	pos := strings.Index(name, "=")

	if pos != -1 {
		value := name[pos+1:]

		if 0 == len(value) {
			i.parsed = append([]string{value}, i.parsed...)
		}

		i.addLongOption(name[0:pos], value)
	} else {
		i.addLongOption(name, "")
	}
}

func (i *ArgvInput) parseArgument(token string) {
	keys := i.definition.GetArgumentsOrder()

	nbArgs := i.countArguments()

	// if input is expecting another argument, add it
	if nbArgs < len(keys) && i.definition.HasArgument(keys[nbArgs]) {
		arg := i.definition.GetArgument(keys[nbArgs])

		if arg.IsArray() {
			i.argumentArrays[arg.GetName()] = []string{token}
		} else {
			i.arguments[arg.GetName()] = token
		}

		// if last argument isArray(), append token to last argument
	} else if nbArgs-1 <= len(keys) &&
		i.definition.HasArgument(keys[nbArgs-1]) &&
		i.definition.GetArgument(keys[nbArgs-1]).IsArray() {
		arg := i.definition.GetArgument(keys[nbArgs-1])
		i.argumentArrays[arg.GetName()] = append(i.argumentArrays[arg.GetName()], token)

		// unexpected argument
	} else {
		all := i.GetDefinition().GetArguments()

		if len(all) != 0 {
			panic(
				errors.New(
					fmt.Sprintf(
						"too many arguments, expected arguments '%s'",
						helper.Implode(" ", getArgumentsMapKeys(all)),
					),
				),
			)
		}

		panic(errors.New(fmt.Sprintf("no arguments expected, got '%s'", token)))
	}
}

func (i *ArgvInput) addShortOption(shortcut string, value string) {
	if !i.definition.HasShortcut(shortcut) {
		panic(errors.New(fmt.Sprintf("the '-%s' option does not exist", shortcut)))
	}

	opt := i.definition.GetOptionForShortcut(shortcut)

	i.addLongOption(opt.GetName(), value)
}

func (i *ArgvInput) addLongOption(name string, value string) {
	if !i.definition.HasOption(name) {
		panic(errors.New(fmt.Sprintf("the '--%s' option does not exist", name)))
	}

	opt := i.definition.GetOption(name)

	if "" != value && !opt.AcceptValue() {
		panic(errors.New(fmt.Sprintf("the '--%s' option does not accept a value", name)))
	} else if !opt.AcceptValue() {
		// TODO find a better way to handle option.None
		value = option.Defined
	}

	if "" == value && opt.AcceptValue() && len(i.parsed) > 0 {
		// if option accepts an optional or mandatory argument
		// let's see if there is one provided
		next := i.parsed[0]
		i.parsed = i.parsed[1:]

		if len(next) > 0 && next[0] != '-' || "" == next {
			value = next
		} else {
			i.parsed = append([]string{next}, i.parsed...)
		}
	}

	if "" == value {
		if opt.IsValueRequired() {
			panic(errors.New(fmt.Sprintf("the '--%s' option requires a value", name)))
		}

		if !opt.IsArray() && !opt.IsValueOptional() {
			value = option.Defined
		}
	}

	if opt.IsArray() {
		i.optionArrays[name] = append(i.optionArrays[name], value)
	} else {
		i.options[name] = value
	}
}

func (i *ArgvInput) countArguments() int {
	return len(i.arguments) + len(i.argumentArrays)
}

func getArgumentsMapKeys(inputs map[string]argument.InputArgument) []string {
	var keys []string

	for k := range inputs {
		keys = append(keys, k)
	}

	return keys
}
