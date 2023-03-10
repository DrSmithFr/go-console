package go_console

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/input"
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/DrSmithFr/go-console/input/option"
	"github.com/DrSmithFr/go-console/output"
	"github.com/DrSmithFr/go-console/table"
	"github.com/DrSmithFr/go-console/verbosity"
	"os"
	"strings"
)

const (
	ExitSuccess = iota
	ExitError
	ExitInvalid
)

// simple constructor
func NewCli() *Cli {
	out := output.NewCliOutput(true, nil)
	in := input.NewArgvInput(nil)

	// manage verbosity
	cmd := CustomCli(in, out)
	cmd.addDefaultOptions()

	return cmd
}

// custom constructor
func CustomCli(in input.InputInterface, out output.OutputInterface) *Cli {
	cmd := &Cli{
		alreadyParsed:    false,
		definitionParsed: true,
	}

	cmd.addDefaultOptions()

	// clone the formatter to retrieve styles and avoid state change
	format := *out.Formatter()

	cmd.in = in
	cmd.out = out
	cmd.lineLength = MAX_LINE_LENGTH
	cmd.bufferedOutput = *output.NewBufferedOutput(false, &format)

	return cmd
}

// Output decorator helpers for the Style Guide
type Cli struct {
	abstractStyler
	alreadyParsed bool

	caller      string
	description string

	// Short definition
	definitionParsed bool
	DefaultOpts      bool

	Name string
	Desc string

	In  input.InputInterface
	Out output.OutputInterface

	Args []Arg
	Opts []Opt
}

func (cmd *Cli) Description() string {
	return cmd.description
}

func (cmd *Cli) SetDescription(description string) *Cli {
	cmd.description = description
	return cmd
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

	Description string

	DefaultValue  string
	DefaultValues []string
}

func (cmd *Cli) addDefaultOptions() {
	cmd.
		// add help option
		AddInputOption(
			option.
				New("help", option.None).
				SetShortcut("p").
				SetDescription("Display help for the given command."),
		).
		// add help option
		AddInputOption(
			option.
				New("no-interaction", option.None).
				SetShortcut("n").
				SetDescription("Do not ask any interactive question"),
		).
		// add verbosity options
		AddInputOption(
			option.
				New("quiet", option.None).
				SetShortcut("q").
				SetDescription("Do not output any message"),
		).
		AddInputOption(
			option.New("verbose", option.Optional).
				SetShortcut("v|vv|vvv").
				SetDescription("Increase the verbosity of messages: 1 for normal output, 2 for more verbose output and 3 for debug"),
		)
}

// Implements io.Writer
func (cmd *Cli) Write(p []byte) (n int, err error) {
	return cmd.out.Write(p)
}

// (helper) add option to input definition
func (cmd *Cli) AddInputOption(opt *option.InputOption) *Cli {
	if cmd.alreadyParsed {
		panic(errors.New("cannot add option on parsed input"))
	}

	cmd.in.Definition().AddOption(*opt)

	return cmd
}

// (helper) add argument to input definition
func (cmd *Cli) AddInputArgument(arg *argument.InputArgument) *Cli {
	if cmd.alreadyParsed {
		panic(errors.New("cannot add argument on parsed input"))
	}

	cmd.in.Definition().AddArgument(*arg)

	return cmd
}

func (cmd *Cli) Build() *Cli {
	if cmd.definitionParsed == false {
		cmd = cmd.parseDefinition()
		cmd.definitionParsed = true
	}

	cmd.parseInput()
	cmd.validateInput()
	cmd.findOutputVerbosity()

	cmd.handleHelpCall()

	return cmd
}

func (cmd *Cli) parseDefinition() *Cli {
	var in input.InputInterface
	if cmd.In != nil {
		in = cmd.In
	} else {
		in = input.NewArgvInput(nil)
	}

	var out output.OutputInterface
	if cmd.Out != nil {
		out = cmd.Out
	} else {
		out = output.NewCliOutput(true, nil)
	}

	if cmd.Name != "" {
		cmd.caller = cmd.Name
	}

	if cmd.Desc != "" {
		cmd.description = cmd.Desc
	}

	cmd.in = in
	cmd.out = out

	// clone the formatter to retrieve styles and avoid state change
	format := *out.Formatter()

	cmd.alreadyParsed = false

	cmd.lineLength = MAX_LINE_LENGTH
	cmd.bufferedOutput = *output.NewBufferedOutput(false, &format)

	if len(cmd.Args) > 0 {
		for _, arg := range cmd.Args {
			newArg := argument.New(arg.Name, arg.Mode).
				SetDescription(arg.Description)

			if arg.DefaultValue != "" {
				newArg.SetDefault(arg.DefaultValue)
			}

			if len(arg.DefaultValues) > 0 {
				newArg.SetDefaults(arg.DefaultValues)
			}

			cmd.AddInputArgument(newArg)
		}
	}

	if len(cmd.Opts) > 0 {
		for _, opt := range cmd.Opts {
			newOpt := option.New(opt.Name, opt.Mode)

			if opt.Shortcut != "" {
				newOpt.SetShortcut(opt.Shortcut)
			}

			if opt.Description != "" {
				newOpt.SetDescription(opt.Description)
			}

			if opt.DefaultValue != "" {
				newOpt.SetDefault(opt.DefaultValue)
			}

			if len(opt.DefaultValues) > 0 {
				newOpt.SetDefaults(opt.DefaultValues)
			}

			cmd.AddInputOption(newOpt)
		}
	}

	if !cmd.DefaultOpts {
		cmd.addDefaultOptions()
	}

	return cmd
}

func (cmd *Cli) parseInput() *Cli {
	if cmd.alreadyParsed {
		panic(errors.New("argv is already parsed"))
	}

	defer cmd.handleParsingException()

	cmd.in.Parse()
	cmd.alreadyParsed = true

	return cmd
}

func (cmd *Cli) validateInput() *Cli {
	if !cmd.alreadyParsed {
		panic(errors.New("cannot validate unparsed input"))
	}

	defer cmd.handleParsingException()

	cmd.in.Validate()

	return cmd
}

func (cmd *Cli) findOutputVerbosity() *Cli {
	level := verbosity.Normal

	if cmd.in.Option("quiet") == option.Defined {
		level = verbosity.Quiet
	} else if cmd.in.Option("verbose") == option.Defined {
		lvl := cmd.in.Option("verbose")
		if lvl == "vv" {
			level = verbosity.Debug
		} else if lvl == "v" {
			level = verbosity.VeryVerbose
		}
	}

	cmd.out.SetVerbosity(level)

	return cmd
}

func (cmd *Cli) handleParsingException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	cmd.PrintError(fmt.Sprintf("%s", err))

	args := os.Args[0]
	synopsis := cmd.in.Definition().Synopsis(false)

	usage := fmt.Sprintf(
		"<info>Usage:</info> <comment>%s %s</comment>",
		args,
		formatter.Escape(synopsis),
	)

	cmd.out.Println(usage)

	os.Exit(2)
}

func (cmd *Cli) HandleRuntimeException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	msg := fmt.Sprintf("%s", err)
	full := fmt.Sprintf("%+v", err)

	traces := strings.TrimPrefix(full, msg)
	traces = strings.Replace(traces, "\n\t", "() at ", -1)

	cmd.PrintError(msg)

	cmd.out.Print("<comment>Exception trace:</comment>")
	for _, trace := range strings.Split(traces, "\n") {
		cmd.out.Println(
			fmt.Sprintf(
				" %s",
				formatter.Escape(trace),
			),
		)
	}

	os.Exit(2)
}

func (cmd *Cli) handleHelpCall() {
	if cmd.in.Option("help") == option.Undefined {
		return
	}

	if cmd.Description() != "" {
		cmd.PrintText("<comment>Description:</comment>")
		cmd.PrintText(cmd.Description())
		cmd.PrintNewLine(1)
	}

	cmd.PrintText("<comment>Usage:</comment>")

	synopsis := cmd.in.Definition().Synopsis(false)
	synopsis = strings.ReplaceAll(synopsis, "<", "\\<")

	var cmdName string
	if cmd.caller != "" {
		cmdName = cmd.caller
	} else {
		cmdName = os.Args[0]
	}

	cmd.out.SetDecorated(false)
	cmd.PrintText(fmt.Sprintf(" %s <info>%s</info>", cmdName, synopsis))
	cmd.out.SetDecorated(true)

	render := table.
		NewRender(cmd.out).
		SetStyleFromName("compact")

	if len(cmd.in.Definition().Arguments()) > 0 {
		cmd.PrintNewLine(1)
		cmd.PrintText("<comment>Arguments:</comment>")

		render.
			SetContent(cmd.createArgsTable()).
			Render()
	}

	if len(cmd.in.Definition().Options()) > 0 {
		cmd.PrintNewLine(1)
		cmd.PrintText("<comment>Options:</comment>")

		render.
			SetContent(cmd.createOptsTable()).
			Render()
	}

	os.Exit(ExitSuccess)
}

func (cmd *Cli) createArgsTable() *table.Table {
	argTab := table.NewTable()

	for _, argKey := range cmd.in.Definition().ArgumentsOrder() {
		arg := cmd.in.Definition().Argument(argKey)

		name := fmt.Sprintf(
			" <info>%s</info>",
			arg.Name(),
		)

		flags := []string{}
		if arg.IsRequired() {
			flags = append(flags, "required")
		} else {
			flags = append(flags, "optional")
		}

		if arg.IsList() {
			flags = append(flags, "list")
		}

		flagLine := fmt.Sprintf("<comment>[%s]</comment>", strings.Join(flags, "|"))

		desc := arg.Description()

		if !arg.IsList() && arg.Default() != "" {
			desc += fmt.Sprintf(
				" <comment>[default: \"%s\"]</comment>",
				arg.Default(),
			)
		}

		if arg.IsList() && arg.Defaults() != nil {
			desc += fmt.Sprintf(
				" <comment>[defaults: \"%s\"]</comment>",
				desc,
				strings.Join(arg.Defaults(), "\", \""),
			)
		}

		argTab.
			AddRowFromString([]string{
				name, flagLine, desc,
			})
	}

	return argTab
}

func (cmd *Cli) createOptsTable() *table.Table {
	optTab := table.NewTable()

	for _, optKey := range cmd.in.Definition().OptionsOrder() {
		opt := cmd.in.Definition().Option(optKey)
		shortcut := ""

		if opt.Shortcut() != "" {
			shortcut = fmt.Sprintf(
				"<info>-%s,</info>",
				opt.Shortcut(),
			)
		}

		name := fmt.Sprintf(
			" <info>--%s</info>",
			opt.Name(),
		)

		desc := opt.Description()

		if !opt.IsList() && opt.Default() != "" {
			desc += fmt.Sprintf(
				" <comment>[default: \"%s\"]</comment>",
				opt.Default(),
			)
		}

		if opt.IsList() && opt.Defaults() != nil {
			desc += fmt.Sprintf(
				" <comment>[defaults: [[\"%s\"]]</comment>",
				desc,
				strings.Join(opt.Defaults(), "\", \""),
			)
		}

		optTab.
			AddRowFromString([]string{
				shortcut, name, desc,
			})
	}

	return optTab
}
