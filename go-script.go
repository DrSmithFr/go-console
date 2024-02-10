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
	"path/filepath"
	"strings"
)

type ExitCode int

const (
	ExitSuccess ExitCode = iota
	ExitError   ExitCode = iota
	ExitInvalid ExitCode = iota
)

// NewScript simple console CLi constructor
func NewScript() *Script {
	// manage verbosity
	cmd := NewScriptCustom(
		input.NewArgvInput(nil),
		output.NewCliOutput(true, nil),
		true,
	)

	return cmd
}

// NewScriptCustom create a new Script with custom input/output with or without default options
func NewScriptCustom(in input.InputInterface, out output.OutputInterface, AddDefaultOptions bool) *Script {
	cmd := &Script{
		inputParsed:      false,
		definitionParsed: true,
	}

	// accessors
	cmd.Input = in
	cmd.Output = out

	// clone the formatter to retrieve styles and avoid state change
	format := *out.Formatter()

	// enable style within the script
	cmd.input = in
	cmd.output = out
	cmd.maxLineLength = MaxLineLength
	cmd.bufferedOutput = *output.NewBufferedOutput(false, &format)

	if AddDefaultOptions {
		cmd.addDefaultOptions()
	}

	return cmd
}

// Script is the base class for all scripts.
type Script struct {
	Styler

	// Short definition
	Input  input.InputInterface
	Output output.OutputInterface

	AddDefaultOpts bool

	Name        string
	Description string

	Arguments []Argument
	Options   []Option

	Runner CommandRunner

	// internal
	inputParsed      bool
	definitionParsed bool
	parentScriptName string

	BuildInfo *BuildInfo
}

func (s *Script) SetDescription(description string) *Script {
	s.Description = description
	return s
}

type Argument struct {
	Name  string
	Value int

	Description string

	DefaultValue  string
	DefaultValues []string
}

type Option struct {
	Name     string
	Shortcut string
	Value    int

	Description string

	DefaultValue  string
	DefaultValues []string
}

func (s *Script) addDefaultOptions() {
	s.
		// add help option
		AddInputOption(
			option.
				New("help", option.None).
				SetShortcut("h").
				SetDescription("Display help for the given command."),
		).
		AddInputOption(
			option.
				New("version", option.None).
				SetShortcut("V").
				SetDescription("Display version for the given command."),
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
func (s *Script) Write(p []byte) (n int, err error) {
	return s.output.Write(p)
}

// AddInputOption add option to input definition (fluent)
func (s *Script) AddInputOption(opt *option.InputOption) *Script {
	if s.inputParsed {
		panic(errors.New("cannot add option on parsed input"))
	}

	s.input.Definition().AddOption(*opt)

	return s
}

// AddInputArgument add argument to input definition (fluent)
func (s *Script) AddInputArgument(arg *argument.InputArgument) *Script {
	if s.inputParsed {
		panic(errors.New("cannot add argument on parsed input"))
	}

	s.input.Definition().AddArgument(*arg)

	return s
}

func (s *Script) Build() *Script {
	if !s.definitionParsed {
		s.parseDefinition()
		s.definitionParsed = true
	}

	s.parseInput()
	s.findOutputVerbosity()
	s.handleHelpCall()
	s.handleVersionCall()

	s.validateInput()

	if s.Runner != nil {
		os.Exit(int(s.Runner(s)))
	}

	return s
}

func (s *Script) parseDefinition() *Script {
	var in input.InputInterface
	var out output.OutputInterface

	if s.Input != nil {
		in = s.Input
	} else {
		in = input.NewArgvInput(nil)
	}

	if s.Output != nil {
		out = s.Output
	} else {
		out = output.NewCliOutput(true, nil)
	}

	// clone the formatter to retrieve styles and avoid state change
	format := *out.Formatter()

	s.input = in
	s.output = out
	s.maxLineLength = MaxLineLength
	s.bufferedOutput = *output.NewBufferedOutput(false, &format)

	if len(s.Arguments) > 0 {
		for _, arg := range s.Arguments {
			newArg := argument.New(arg.Name, arg.Value).
				SetDescription(arg.Description)

			if arg.DefaultValue != "" {
				newArg.SetDefault(arg.DefaultValue)
			}

			if len(arg.DefaultValues) > 0 {
				newArg.SetDefaults(arg.DefaultValues)
			}

			s.AddInputArgument(newArg)
		}
	}

	if len(s.Options) > 0 {
		for _, opt := range s.Options {
			newOpt := option.New(opt.Name, opt.Value)

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

			s.AddInputOption(newOpt)
		}
	}

	if !s.AddDefaultOpts {
		s.addDefaultOptions()
	}

	return s
}

func (s *Script) parseInput() *Script {
	if s.inputParsed {
		panic(errors.New("argv is already parsed"))
	}

	defer s.handleParsingException()

	s.input.Parse()
	s.inputParsed = true

	return s
}

func (s *Script) validateInput() *Script {
	if !s.inputParsed {
		panic(errors.New("cannot validate unparsed input"))
	}

	defer s.handleParsingException()

	s.input.Validate()

	return s
}

func (s *Script) findOutputVerbosity() *Script {
	level := verbosity.Normal

	if s.input.Option("quiet") == option.Defined {
		level = verbosity.Quiet
	} else if s.input.Option("verbose") == option.Defined {
		lvl := s.input.Option("verbose")
		if lvl == "vv" {
			level = verbosity.Debug
		} else if lvl == "v" {
			level = verbosity.VeryVerbose
		}
	}

	s.output.SetVerbosity(level)

	return s
}

func (s *Script) handleParsingException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	s.PrintError(fmt.Sprintf("%s", err))

	args := os.Args[0]
	synopsis := s.input.Definition().Synopsis(false)

	usage := fmt.Sprintf(
		"<info>Usage:</info> <comment>%s %s</comment>",
		args,
		formatter.Escape(synopsis),
	)

	s.output.Println(usage)

	os.Exit(2)
}

func (s *Script) HandleRuntimeException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	msg := fmt.Sprintf("%s", err)
	full := fmt.Sprintf("%+v", err)

	traces := strings.TrimPrefix(full, msg)
	traces = strings.Replace(traces, "\n\t", "() at ", -1)

	s.PrintError(msg)

	s.output.Print("<comment>Exception trace:</comment>")
	for _, trace := range strings.Split(traces, "\n") {
		s.output.Println(
			fmt.Sprintf(
				" %s",
				formatter.Escape(trace),
			),
		)
	}

	os.Exit(2)
}

func (s *Script) handleHelpCall() {
	if s.input.Option("help") == option.Undefined {
		return
	}

	if s.Description != "" {
		s.PrintText("<comment>Description:</comment>")
		s.PrintText(s.Description)
		s.PrintNewLine(1)
	}

	s.PrintText("<comment>Usage:</comment>")

	synopsis := s.input.Definition().Synopsis(false)
	synopsis = strings.ReplaceAll(synopsis, "<", "\\<")

	cmdName := filepath.Base(os.Args[0])

	if s.parentScriptName != "" {
		cmdName = s.parentScriptName + " " + cmdName
	}

	s.output.SetDecorated(false)
	s.PrintText(fmt.Sprintf(" %s <info>%s</info>", cmdName, synopsis))
	s.output.SetDecorated(true)

	render := table.
		NewRender(s.output).
		SetStyleFromName("compact")

	if len(s.input.Definition().Arguments()) > 0 {
		s.PrintNewLine(1)
		s.PrintText("<comment>Arguments:</comment>")

		render.
			SetContent(s.createArgsTable()).
			Render()
	}

	if len(s.input.Definition().Options()) > 0 {
		s.PrintNewLine(1)
		s.PrintText("<comment>Options:</comment>")

		render.
			SetContent(s.createOptsTable()).
			Render()
	}

	os.Exit(int(ExitSuccess))
}

func (s *Script) handleVersionCall() {
	// deprecated but still supported and prior to other options
	cmdName := s.Name

	if cmdName == "" && s.BuildInfo != nil {
		cmdName = s.BuildInfo.Name
	}

	if cmdName == "" {
		// fallback to the script name
		cmdName = filepath.Base(os.Args[0])
	}

	if s.parentScriptName != "" {
		cmdName = s.parentScriptName + " " + cmdName
	}

	version := "latest"

	if s.BuildInfo != nil && s.BuildInfo.Version != "" {
		version = s.BuildInfo.Version
	}

	tagLine := fmt.Sprintf("<info>%s</info><comment>@%s</comment>", cmdName, version)

	if s.BuildInfo != nil && s.BuildInfo.BuildFlag != "" {
		tagLine += " " + s.BuildInfo.BuildFlag
	}

	s.PrintText(tagLine)
	os.Exit(int(ExitSuccess))
}

func (s *Script) createArgsTable() *table.Table {
	argTab := table.NewTable()

	for _, argKey := range s.input.Definition().ArgumentsOrder() {
		arg := s.input.Definition().Argument(argKey)

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

		if arg.IsList() && arg.Defaults() != nil && len(arg.Defaults()) > 0 {
			desc += fmt.Sprintf(
				" <comment>[defaults: \"%s\"]</comment>",
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

func (s *Script) createOptsTable() *table.Table {
	optTab := table.NewTable()

	for _, optKey := range s.input.Definition().OptionsOrder() {
		opt := s.input.Definition().Option(optKey)
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

func (s *Script) SetParentScriptName(name string) {
	s.parentScriptName = name
}
