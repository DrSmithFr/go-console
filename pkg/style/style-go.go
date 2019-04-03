package style

import (
	"errors"
	"github.com/DrSmithFr/go-console/pkg/input"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/DrSmithFr/go-console/pkg/output"
)

// simple constructor
func NewConsoleStyler() *GoStyler {
	out := output.NewConsoleOutput(true, nil)
	in := input.NewArgvInput(nil)

	return NewGoStyler(in, out)
}

// custom constructor
func NewGoStyler(in input.InputInterface, out output.OutputInterface) *GoStyler {
	g := &GoStyler{
		alreadyParsed: false,
	}

	// clone the formatter to retrieve styles and avoid state change
	format := *out.GetFormatter()

	g.in = in
	g.out = out
	g.lineLength = MAX_LINE_LENGTH
	g.bufferedOutput = *output.NewBufferedOutput(false, &format)

	return g
}

// Output decorator helpers for the Style Guide
type GoStyler struct {
	abstractStyler
	alreadyParsed bool
}

// (helper) add option to input definition
func (g *GoStyler) AddInputOption(opt *option.InputOption) *GoStyler {
	if g.alreadyParsed {
		panic(errors.New("cannot add option on parsed input"))
	}

	g.in.GetDefinition().AddOption(*opt)

	return g
}

// (helper) add argument to input definition
func (g *GoStyler) AddInputArgument(arg *argument.InputArgument) *GoStyler {
	if g.alreadyParsed {
		panic(errors.New("cannot add argument on parsed input"))
	}

	g.in.GetDefinition().AddArgument(*arg)

	return g
}

// (helper) parse input argv
func (g *GoStyler) ParseInput() *GoStyler {
	if g.alreadyParsed {
		panic(errors.New("argv is already parsed"))
	}

	g.in.Parse()
	g.alreadyParsed = true

	return g
}

// (helper) validate input argv
func (g *GoStyler) ValidateInput() *GoStyler {
	if !g.alreadyParsed {
		panic(errors.New("cannot validate unparsed input"))
	}

	g.in.Validate()

	return g
}