package input

import (
	"errors"
	"os"
	"regexp"
)

func NewArgvInput(argv []string) *ArgvInput {
	input := new(ArgvInput)

	if nil == argv {
		input.tokens = os.Args[1:]
	} else {
		input.tokens = argv[1:]
	}

	input.doParse = input.Parse

	return input
}

type ArgvInput struct {
	abstractInput
	tokens []string
	parsed []string
}

func (i *ArgvInput) GetFirstArgument() string {
	for _, token := range i.tokens {
		if "" != token && '-' == token[0] {
			continue
		}

		return token
	}

	panic(errors.New("first argument not found"))
}

func (i *ArgvInput) HasParameterOption(values []string, onlyParams bool) bool {
	panic("implement me")
}

func (i *ArgvInput) GetParameterOption(values []string, defaultValue string, onlyParams bool) {
	panic("implement me")
}

func (i *ArgvInput) Parse() {
	parseOptions := true
	i.parsed = i.tokens

	longOptionRegex := regexp.MustCompile("^--")

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
		} else if parseOptions && longOptionRegex.MatchString(token) {
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

func (i *ArgvInput) parseShortOption(s string) {

}

func (i *ArgvInput) parseShortOptionSet(s string) {

}

func (i *ArgvInput) parseLongOption(s string) {

}

func (i *ArgvInput) parseArgument(token string) {
	keys := i.definition.GetArgumentsOrder()

	if len(keys) > 0 && i.definition.HasArgument(keys[len(keys)-1]) {
		arg := i.definition.GetArgument(keys[len(keys)-1])

		if arg.IsArray() {
			i.argumentArrays[arg.GetName()] = []string{token}
		} else {
			i.arguments[arg.GetName()] = token
		}
	} else if len(keys) > 1 &&
		i.definition.HasArgument(keys[len(keys)-2]) &&
		i.definition.GetArgument(keys[len(keys)-2]).IsArray() {
		arg := i.definition.GetArgument(keys[len(keys)-2])
		i.argumentArrays[arg.GetName()] = append(i.argumentArrays[arg.GetName()], token)
	}
}

