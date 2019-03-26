package test_helper

import (
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/definition"
	"github.com/DrSmithFr/go-console/pkg/input/option"
)

func NewParserPattern(argv []string) *ParserPattern {
	pattern := &ParserPattern{
		argv:       argv,
		definition: *definition.New(),

		arguments:      map[string]string{},
		argumentArrays: map[string][]string{},

		options:      map[string]string{},
		optionArrays: map[string][]string{},

		message: "",
	}

	return pattern
}

type ParserPattern struct {
	argv       []string
	definition definition.InputDefinition

	arguments      map[string]string
	argumentArrays map[string][]string

	options      map[string]string
	optionArrays map[string][]string

	message string
}

func (p *ParserPattern) Message() string {
	return p.message
}

func (p *ParserPattern) SetMessage(message string) *ParserPattern {
	p.message = message
	return p
}

func (p *ParserPattern) Argv() []string {
	return p.argv
}

func (p *ParserPattern) Definition() definition.InputDefinition {
	return p.definition
}

func (p *ParserPattern) AddArgument(arg argument.InputArgument) *ParserPattern {
	p.definition.AddArgument(arg)
	return p
}

func (p *ParserPattern) AddOption(opt option.InputOption) *ParserPattern {
	p.definition.AddOption(opt)
	return p
}

func (p *ParserPattern) Arguments() map[string]string {
	return p.arguments
}

func (p *ParserPattern) SetArguments(arguments map[string]string) *ParserPattern {
	p.arguments = arguments
	return p
}

func (p *ParserPattern) ArgumentArrays() map[string][]string {
	return p.argumentArrays
}

func (p *ParserPattern) SetArgumentArrays(argumentArrays map[string][]string) *ParserPattern {
	p.argumentArrays = argumentArrays
	return p
}

func (p *ParserPattern) Options() map[string]string {
	return p.options
}

func (p *ParserPattern) SetOptions(options map[string]string) *ParserPattern {
	p.options = options
	return p
}

func (p *ParserPattern) OptionArrays() map[string][]string {
	return p.optionArrays
}

func (p *ParserPattern) SetOptionArrays(optionArrays map[string][]string)  *ParserPattern {
	p.optionArrays = optionArrays
	return p
}
