package style

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/formatter"
	"github.com/DrSmithFr/go-console/pkg/input"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/DrSmithFr/go-console/pkg/output"
	"github.com/DrSmithFr/go-console/pkg/verbosity"
	"os"
	"strings"
)

// Deprecated: NewConsoleStyler is deprecated, use NewConsoleCommand instead.
func NewConsoleStyler() *GoStyler {
	return NewConsoleCommand()
}

// simple constructor
func NewConsoleCommand() *GoStyler {
	out := output.NewConsoleOutput(true, nil)
	in := input.NewArgvInput(nil)

	// manage verbosity
	io := NewCommandStyler(in, out)
	io.
		AddInputOption(
			option.
				New("quiet", option.None).
				SetShortcut("q"),
		).
		AddInputOption(
			option.
				New("verbose", option.None).
				SetShortcut("v"),
		).
		AddInputOption(
			option.
				New("very-verbose", option.None).
				SetShortcut("vv"),
		).
		AddInputOption(
			option.
				New("debug", option.None).
				SetShortcut("vvv"),
		)

	return io
}

// Deprecated: NewGoStyler is deprecated, use NewCommandStyler instead.
func NewGoStyler(in input.InputInterface, out output.OutputInterface) *GoStyler {
	return NewCommandStyler(in, out)
}

// custom constructor
func NewCommandStyler(in input.InputInterface, out output.OutputInterface) *GoStyler {
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

func (g *GoStyler) Build() *GoStyler {
	g.parseInput()
	g.validateInput()
	g.findOutputVerbosity()

	return g
}

func (g *GoStyler) parseInput() *GoStyler {
	if g.alreadyParsed {
		panic(errors.New("argv is already parsed"))
	}

	defer g.handleParsingException()

	g.in.Parse()
	g.alreadyParsed = true

	return g
}

func (g *GoStyler) validateInput() *GoStyler {
	if !g.alreadyParsed {
		panic(errors.New("cannot validate unparsed input"))
	}

	defer g.handleParsingException()

	g.in.Validate()

	return g
}

func (g *GoStyler) findOutputVerbosity() *GoStyler {
	level := verbosity.Normal

	if g.in.GetOption("quiet") == option.Defined {
		level = verbosity.Quiet
	} else if g.in.GetOption("verbose") == option.Defined {
		level = verbosity.Verbose
	} else if g.in.GetOption("very-verbose") == option.Defined {
		level = verbosity.VeryVerbose
	} else if g.in.GetOption("debug") == option.Defined {
		level = verbosity.Debug
	}

	g.out.SetVerbosity(level)

	return g
}

func (g *GoStyler) handleParsingException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	g.Error(fmt.Sprintf("%s", err))

	cmd := os.Args[0]
	synopsis := g.in.GetDefinition().GetSynopsis(false)

	usage := fmt.Sprintf(
		"<info>Usage:</info> <comment>%s %s</comment>",
		cmd,
		formatter.Escape(synopsis),
	)

	g.out.Writeln(usage)

	os.Exit(2)
}

func (g *GoStyler) HandleRuntimeException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	msg := fmt.Sprintf("%s", err)
	full := fmt.Sprintf("%+v", err)

	traces := strings.TrimPrefix(full, msg)
	traces = strings.Replace(traces, "\n\t", "() at ", -1)

	g.Error(msg)

	g.out.Write("<comment>Exception trace:</comment>")
	for _, trace := range strings.Split(traces, "\n") {
		g.out.Writeln(
			fmt.Sprintf(
				" %s",
				formatter.Escape(trace),
			),
		)
	}

	os.Exit(2)
}
