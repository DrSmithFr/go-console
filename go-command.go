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
	"regexp"
	"sort"
	"strings"
)

type CommandRunner func(cmd *Script) ExitCode

// NewCommand create a new console script
func NewCommand() *Command {
	argv := os.Args

	if len(argv) > 2 {
		argv = argv[0:2]
	}

	script := newCommandCustom(
		input.NewArgvInput(argv),
		output.NewCliOutput(true, nil),
		true,
	)

	return script
}

// NewScriptCustom create a new script with custom input/output with or without default options
func newCommandCustom(in input.InputInterface, out output.OutputInterface, AddDefaultOptions bool) *Command {
	script := &Command{
		inputParsed:      false,
		definitionParsed: true,

		Description: "Script to run a commands",
		Scripts:     []*Script{},

		registeredScripts: make(map[string]*Script),
		runners:           make(map[string]CommandRunner),
	}

	// clone the formatter to retrieve styles and avoid state change
	format := *out.Formatter()

	// accessors
	script.Input = in
	script.Output = out

	// enable style within the script
	script.input = in
	script.output = out
	script.bufferedOutput = *output.NewBufferedOutput(false, &format)
	script.maxLineLength = MaxLineLength

	if AddDefaultOptions {
		script.addDefaultOptions()
	}

	return script
}

type Command struct {
	Styler

	UseNamespace bool
	Description  string

	Output output.OutputInterface
	Input  input.InputInterface

	Scripts           []*Script
	registeredScripts map[string]*Script
	runners           map[string]CommandRunner

	inputParsed      bool
	definitionParsed bool

	Info *Info
}

type Info struct {
	Name      string
	Version   string
	BuildFlag string
}

// (helper) add default options
func (c *Command) addDefaultOptions() {
	c.
		// add command argument
		addInputArgument(
			argument.
				New("command", argument.Optional).
				SetDescription("The command to execute"),
		).
		// add help option
		addInputOption(
			option.
				New("help", option.None).
				SetShortcut("h").
				SetDescription("Display help for the given command."),
		).
		// add help option
		addInputOption(
			option.
				New("no-interaction", option.None).
				SetShortcut("n").
				SetDescription("Do not ask any interactive question"),
		).
		// add verbosity options
		addInputOption(
			option.
				New("quiet", option.None).
				SetShortcut("q").
				SetDescription("Do not output any message"),
		).
		addInputOption(
			option.New("verbose", option.Optional).
				SetShortcut("v|vv|vvv").
				SetDescription("Increase the verbosity of messages: 1 for normal output, 2 for more verbose output and 3 for debug"),
		)

	if c.Info != nil {
		c.addInputOption(
			option.
				New("version", option.None).
				SetShortcut("V").
				SetDescription("Display this application version."),
		)
	}
}

// addInputOption add option to input definition (fluent)
func (c *Command) addInputOption(opt *option.InputOption) *Command {
	if c.inputParsed {
		panic(errors.New("cannot add option on parsed input"))
	}

	c.input.Definition().AddOption(*opt)

	return c
}

// addInputArgument add argument to input definition (fluent)
func (c *Command) addInputArgument(arg *argument.InputArgument) *Command {
	if c.inputParsed {
		panic(errors.New("cannot add argument on parsed input"))
	}

	c.input.Definition().AddArgument(*arg)

	return c
}

// AddScript add a command to the script (fluent)
func (c *Command) AddScript(cmd *Script, run CommandRunner) *Command {
	if c.registeredScripts[cmd.Name] != nil {
		panic(errors.New(fmt.Sprintf("Script '%s' already exists", cmd.Name)))
	}

	c.registeredScripts[cmd.Name] = cmd
	c.runners[cmd.Name] = run
	return c
}

// Script return a script by name
func (c *Command) Script(name string) *Script {
	return c.registeredScripts[name]
}

// Runner return a command runner by command name
func (c *Command) Runner(name string) CommandRunner {
	return c.runners[name]
}

// ScriptOrderByName return a list of command name sorted by name
func (c *Command) ScriptOrderByName() []string {
	names := []string{}

	for key := range c.registeredScripts {
		names = append(names, key)
	}

	sort.Strings(names)

	return names
}

func (c *Command) FindScriptOrderByName(search string) []string {
	namespace := strings.Split(search, ":")

	var pattern string
	if len(namespace) == 1 {
		pattern = "^" + search + ".*"
	} else {
		pattern = "^"
		for key, part := range namespace {
			pattern += part + "[^:]*"
			if key < len(namespace)-1 {
				pattern += ":"
			}
		}
	}

	regex := regexp.MustCompile(pattern)
	names := []string{}

	for key := range c.registeredScripts {
		if regex.MatchString(key) {
			names = append(names, key)
		}
	}

	sort.Strings(names)

	return names
}

func (c *Command) Run() {
	c.build()

	if c.Info != nil && option.Defined == c.input.Option("version") {
		c.showVersion()

		os.Exit(int(ExitSuccess))
	}

	command := c.input.Argument("command")

	if command == "" {
		c.showHelp()

		if option.Defined == c.input.Option("help") {
			os.Exit(int(ExitSuccess))
		}

		os.Exit(int(ExitInvalid))
	}

	script := c.Script(command)

	if script == nil && !c.UseNamespace {
		c.PrintError(fmt.Sprintf("Command '%s' is not defined.", command))
		os.Exit(int(ExitInvalid))
	}

	if script == nil && c.UseNamespace {
		scripts := c.FindScriptOrderByName(command)

		if len(scripts) == 0 {
			c.PrintError(fmt.Sprintf("Command '%s' is not defined.", command))
			os.Exit(int(ExitInvalid))
		}

		if len(scripts) > 1 {
			// show possible commands
			c.showAutocompletionHelp(command, scripts)
			os.Exit(int(ExitInvalid))
		} else {
			// autocompleted command
			command = scripts[0]
			script = c.Script(command)
		}
	}

	run := c.Runner(command)

	if run == nil {
		_, err := fmt.Fprintf(c.output, "<error>Script '%s' must have runner to work within script.</error>", command)

		if err != nil {
			panic(err)
		}

		os.Exit(int(ExitError))
	}

	argv := os.Args

	// setup script
	script.Input = input.NewArgvInput(argv[1:])
	script.SetParentScriptName(argv[0])
	script.Output = c.output

	script.Build()
	os.Exit(int(run(script)))
}

// Run parse Definition and input and handle all the script logic
func (c *Command) build() {
	if !c.definitionParsed {
		c.parseDefinition()
		c.definitionParsed = true
	}

	c.parseInput()
	c.validateInput()
	c.findOutputVerbosity()
	c.registerCommands()
}

func (c *Command) registerCommands() {
	for _, cmd := range c.Scripts {
		if cmd.Runner == nil {
			panic(errors.New(fmt.Sprintf("Script '%s' has no runner", cmd.Name)))
		}

		c.AddScript(cmd, cmd.Runner)
	}
}

func (c *Command) parseDefinition() {
	argv := os.Args

	if len(argv) > 2 {
		argv = argv[0:2]
	}

	var in input.InputInterface
	var out output.OutputInterface

	if c.Input == nil {
		in = input.NewArgvInput(argv)
	} else {
		in = c.Input
	}

	if c.Output == nil {
		out = output.NewCliOutput(true, nil)
	} else {
		out = c.Output
	}

	// clone the formatter to retrieve styles and avoid state change
	format := *out.Formatter()

	// accessors
	c.Input = in
	c.Output = out

	// enable style within the script
	c.input = in
	c.output = out
	c.bufferedOutput = *output.NewBufferedOutput(false, &format)
	c.maxLineLength = MaxLineLength

	c.addDefaultOptions()
	c.inputParsed = false

	c.registeredScripts = make(map[string]*Script)
	c.runners = make(map[string]CommandRunner)

}

func (c *Command) parseInput() *Command {
	if c.inputParsed {
		panic(errors.New("argv is already parsed"))
	}

	defer c.handleParsingException()

	c.input.Parse()
	c.inputParsed = true

	return c
}

func (c *Command) validateInput() *Command {
	if !c.inputParsed {
		panic(errors.New("cannot validate unparsed input"))
	}

	defer c.handleParsingException()
	c.input.Validate()

	return c
}

func (c *Command) findOutputVerbosity() *Command {
	level := verbosity.Normal

	if c.input.Option("quiet") == option.Defined {
		level = verbosity.Quiet
	} else if c.input.Option("verbose") == option.Defined {
		lvl := c.input.Option("verbose")
		if lvl == "vv" {
			level = verbosity.Debug
		} else if lvl == "v" {
			level = verbosity.VeryVerbose
		}
	}

	c.output.SetVerbosity(level)

	return c
}

func (c *Command) handleParsingException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	_, err1 := fmt.Fprintf(c.output, "<error>%s</error>", err)

	if err1 != nil {
		panic(err1)
	}

	args := os.Args[0]
	synopsis := c.input.Definition().Synopsis(false)

	usage := fmt.Sprintf(
		"<info>Usage:</info> <comment>%s %s</comment>",
		args,
		formatter.Escape(synopsis),
	)

	c.output.Println(usage)

	os.Exit(2)
}

func (c *Command) HandleRuntimeException() {
	err := recover()

	if err == nil {
		// nothing append, continue
		return
	}

	msg := fmt.Sprintf("%s", err)
	full := fmt.Sprintf("%+v", err)

	traces := strings.TrimPrefix(full, msg)
	traces = strings.Replace(traces, "\n\t", "() at ", -1)

	_, err1 := fmt.Fprintf(c.output, "<error>%s</error>", msg)

	if err1 != nil {
		panic(err1)
	}

	c.output.Print("<comment>Exception trace:</comment>")
	for _, trace := range strings.Split(traces, "\n") {
		c.output.Println(
			fmt.Sprintf(
				" %s",
				formatter.Escape(trace),
			),
		)
	}

	os.Exit(2)
}

func (c *Command) showHelp() {
	c.displayHelpIntro()

	render := table.
		NewRender(c.output).
		SetStyleFromName("compact")

	if len(c.input.Definition().Options()) > 0 {
		c.displayOptsHelper(*render)
	}

	if len(c.Scripts) > 0 {
		if c.UseNamespace {
			c.displayAllScriptsByNamespacesTable(*render)
		} else {
			c.displayAllScriptsTable(*render)
		}
	}
}

func (c *Command) showVersion() {
	appName := filepath.Base(os.Args[0])

	if c.Info == nil {
		c.PrintText(fmt.Sprintf(
			"<info>%s</info>@latest",
			appName,
		))
	}

	if c.Info.Name != "" {
		appName = c.Info.Name
	}

	tagLine := fmt.Sprintf(
		"<info>%s</info><comment>@%s</comment>",
		appName,
		c.Info.Version,
	)

	if c.Info.BuildFlag != "" {
		tagLine += " " + c.Info.BuildFlag
	}

	c.PrintText(tagLine)
}

func (c *Command) showAutocompletionHelp(command string, scripts []string) {
	c.displayHelpIntro()

	render := table.
		NewRender(c.output).
		SetStyleFromName("compact")

	if len(c.input.Definition().Options()) > 0 {
		c.displayOptsHelper(*render)
	}

	if len(scripts) > 0 {
		c.displayScriptsTable(command, scripts, *render)
	}
}

func (c *Command) displayHelpIntro() {
	if c.Description != "" {
		c.PrintText("<comment>Description:</comment>")
		c.PrintText(c.Description)
		c.PrintNewLine(1)
	}

	c.PrintText("<comment>Usage:</comment>")
	c.PrintText(" command [options] [arguments]")
}

func (c *Command) displayOptsHelper(render table.TableRender) {
	c.PrintNewLine(1)
	c.PrintText("<comment>Options:</comment>")

	optTab := table.NewTable()

	for _, optKey := range c.input.Definition().OptionsOrder() {
		opt := c.input.Definition().Option(optKey)
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

	render.
		SetContent(optTab).
		Render()
}

func (c *Command) displayAllScriptsTable(render table.TableRender) {
	c.PrintNewLine(1)
	c.PrintText("<comment>Available commands:</comment>")

	argTab := table.NewTable()

	for _, key := range c.ScriptOrderByName() {
		cmd := c.Script(key)

		name := fmt.Sprintf(
			" <info>%s</info>",
			cmd.Name,
		)

		argTab.
			AddRowFromString([]string{
				name, cmd.Description,
			})
	}

	render.
		SetContent(argTab).
		Render()
}

func (c *Command) displayAllScriptsByNamespacesTable(render table.TableRender) {
	c.PrintNewLine(1)
	c.PrintText("<comment>Available commands:</comment>")

	argTab := table.NewTable()
	oldNamespace := ""
	for _, key := range c.ScriptOrderByName() {
		cmd := c.Script(key)

		namespace := strings.Split(cmd.Name, ":")[0]

		name := fmt.Sprintf(
			" <info>%s</info>",
			cmd.Name,
		)

		if namespace != oldNamespace {
			argTab.
				AddRow(
					&table.TableRow{
						Columns: map[int]table.TableColumnInterface{
							0: &table.TableColumn{
								Cell: &table.TableCell{
									Value:   fmt.Sprintf("<comment>%s</comment>", namespace),
									Colspan: 2,
								},
							},
						},
					},
				)

			oldNamespace = namespace
		}

		argTab.
			AddRowFromString([]string{
				name, cmd.Description,
			})
	}

	render.
		SetContent(argTab).
		Render()
}

func (c *Command) displayScriptsTable(command string, script []string, render table.TableRender) {
	c.PrintNewLine(1)
	c.PrintText(fmt.Sprintf("<comment>Available commands for the pattern '%s':</comment>", command))

	argTab := table.NewTable()

	for _, key := range script {
		cmd := c.Script(key)

		name := fmt.Sprintf(
			" <info>%s</info>",
			cmd.Name,
		)

		argTab.
			AddRowFromString([]string{
				name, cmd.Description,
			})
	}

	render.
		SetContent(argTab).
		Render()
}
