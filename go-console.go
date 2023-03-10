package go_console

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/input"
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/DrSmithFr/go-console/input/option"
	"github.com/DrSmithFr/go-console/output"
	"github.com/DrSmithFr/go-console/verbosity"
	"os"
	"strings"
)

// simple constructor
func NewCli() *Cli {
	out := output.NewCliOutput(true, nil)
	in := input.NewArgvInput(nil)

	// manage verbosity
	cmd := CustomCli(in, out)
	cmd.
		AddInputOption(
			option.New("quiet", option.None).
				SetShortcut("q"),
		).
		AddInputOption(
			option.New("verbose", option.None).
				SetShortcut("v"),
		).
		AddInputOption(
			option.New("very-verbose", option.None).
				SetShortcut("vv"),
		).
		AddInputOption(
			option.New("debug", option.None).
				SetShortcut("vvv"),
		)

	return cmd
}

// custom constructor
func CustomCli(in input.InputInterface, out output.OutputInterface) *Cli {
	g := &Cli{
		alreadyParsed:    false,
		definitionParsed: true,
	}

	// clone the formatter to retrieve styles and avoid state change
	format := *out.Formatter()

	g.in = in
	g.out = out
	g.lineLength = MAX_LINE_LENGTH
	g.bufferedOutput = *output.NewBufferedOutput(false, &format)

	return g
}

// Output decorator helpers for the Style Guide
type Cli struct {
	abstractStyler
	alreadyParsed bool

	// Short definition
	definitionParsed bool
	DefaultOpts      bool

	In  input.InputInterface
	Out output.OutputInterface

	Args []Arg
	Opts []Opt
}

type Arg struct {
	Name string
	Mode int

	Description string

	DefaultValue  string
	DefaultValues []string
}

type Opt struct {
	Name     string
	Shortcut string
	Mode     int

	description string

	DefaultValue  string
	DefaultValues []string
}

// Implements io.Writer
func (g *Cli) Write(p []byte) (n int, err error) {
	return g.out.Write(p)
}

// (helper) add option to input definition
func (g *Cli) AddInputOption(opt *option.InputOption) *Cli {
	if g.alreadyParsed {
		panic(errors.New("cannot add option on parsed input"))
	}

	g.in.Definition().AddOption(*opt)

	return g
}

// (helper) add argument to input definition
func (g *Cli) AddInputArgument(arg *argument.InputArgument) *Cli {
	if g.alreadyParsed {
		panic(errors.New("cannot add argument on parsed input"))
	}

	g.in.Definition().AddArgument(*arg)

	return g
}

func (g *Cli) Build() *Cli {
	if g.definitionParsed == false {
		g = g.parseDefinition()
		g.definitionParsed = true
	}

	g.parseInput()
	g.validateInput()
	g.findOutputVerbosity()

	return g
}

func (g *Cli) parseDefinition() *Cli {
	var in input.InputInterface
	if g.In != nil {
		in = g.In
	} else {
		in = input.NewArgvInput(nil)
	}

	var out output.OutputInterface
	if g.Out != nil {
		out = g.Out
	} else {
		out = output.NewCliOutput(true, nil)
	}

	g.in = in
	g.out = out

	// clone the formatter to retrieve styles and avoid state change
	format := *out.Formatter()

	g.alreadyParsed = false

	g.lineLength = MAX_LINE_LENGTH
	g.bufferedOutput = *output.NewBufferedOutput(false, &format)

	if len(g.Args) > 0 {
		for _, arg := range g.Args {
			newArg := argument.New(arg.Name, arg.Mode).
				SetDescription(arg.Description)

			if arg.DefaultValue != "" {
				newArg.SetDefault(arg.DefaultValue)
			}

			if len(arg.DefaultValues) > 0 {
				newArg.SetDefaults(arg.DefaultValues)
			}

			g.AddInputArgument(newArg)
		}
	}

	if len(g.Opts) > 0 {
		for _, opt := range g.Opts {
			newOpt := option.New(opt.Name, opt.Mode)

			if opt.Shortcut != "" {
				newOpt.SetShortcut(opt.Shortcut)
			}

			if opt.description != "" {
				newOpt.SetDescription(opt.description)
			}

			if opt.DefaultValue != "" {
				newOpt.SetDefault(opt.DefaultValue)
			}

			if len(opt.DefaultValues) > 0 {
				newOpt.SetDefaults(opt.DefaultValues)
			}

			g.AddInputOption(newOpt)
		}
	}

	if !g.DefaultOpts {
		g.
			AddInputOption(
				option.New("quiet", option.None).
					SetShortcut("q"),
			).
			AddInputOption(
				option.New("verbose", option.None).
					SetShortcut("v"),
			).
			AddInputOption(
				option.New("very-verbose", option.None).
					SetShortcut("vv"),
			).
			AddInputOption(
				option.New("debug", option.None).
					SetShortcut("vvv"),
			)
	}

	return g
}

func (g *Cli) parseInput() *Cli {
	if g.alreadyParsed {
		panic(errors.New("argv is already parsed"))
	}

	defer g.handleParsingException()

	g.in.Parse()
	g.alreadyParsed = true

	return g
}

func (g *Cli) validateInput() *Cli {
	if !g.alreadyParsed {
		panic(errors.New("cannot validate unparsed input"))
	}

	defer g.handleParsingException()

	g.in.Validate()

	return g
}

func (g *Cli) findOutputVerbosity() *Cli {
	level := verbosity.Normal

	if g.in.Option("quiet") == option.Defined {
		level = verbosity.Quiet
	} else if g.in.Option("verbose") == option.Defined {
		level = verbosity.Verbose
	} else if g.in.Option("very-verbose") == option.Defined {
		level = verbosity.VeryVerbose
	} else if g.in.Option("debug") == option.Defined {
		level = verbosity.Debug
	}

	g.out.SetVerbosity(level)

	return g
}

func (g *Cli) handleParsingException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	g.PrintError(fmt.Sprintf("%s", err))

	cmd := os.Args[0]
	synopsis := g.in.Definition().Synopsis(false)

	usage := fmt.Sprintf(
		"<info>Usage:</info> <comment>%s %s</comment>",
		cmd,
		formatter.Escape(synopsis),
	)

	g.out.Println(usage)

	os.Exit(2)
}

func (g *Cli) HandleRuntimeException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	msg := fmt.Sprintf("%s", err)
	full := fmt.Sprintf("%+v", err)

	traces := strings.TrimPrefix(full, msg)
	traces = strings.Replace(traces, "\n\t", "() at ", -1)

	g.PrintError(msg)

	g.out.Print("<comment>Exception trace:</comment>")
	for _, trace := range strings.Split(traces, "\n") {
		g.out.Println(
			fmt.Sprintf(
				" %s",
				formatter.Escape(trace),
			),
		)
	}

	os.Exit(2)
}
