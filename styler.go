package go_console

import (
	"fmt"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/helper"
	"github.com/DrSmithFr/go-console/input"
	"github.com/DrSmithFr/go-console/output"
	"github.com/DrSmithFr/go-console/verbosity"
	"io"
	"strings"
)

// MaxLineLength terminal max line length
const MaxLineLength = 120

type Styler struct {
	input          input.InputInterface
	output         output.OutputInterface
	bufferedOutput output.BufferedOutput
	maxLineLength  int
}

// Implements io.Writer

var _ io.Writer = (*Styler)(nil)

func (g *Styler) Write(p []byte) (n int, err error) {
	return g.output.Write(p)
}

// Implements StylerInterface

var _ StylerInterface = (*Styler)(nil)

// SetMaxLineLength overwrite terminal max line length
func (g *Styler) SetMaxLineLength(length int) {
	g.maxLineLength = length
}

// MaxLineLength return current max terminal line length
func (g *Styler) MaxLineLength() int {
	return g.maxLineLength
}

// PrintNewLine print n newline(n).
func (g *Styler) PrintNewLine(count int) {
	g.writeList([]string{strings.Repeat("\n", count)}, false)
}

// PrintTitle formats and print a command title.
func (g *Styler) PrintTitle(message string) {
	g.autoPrependBlock()

	messageRealLength := helper.StrlenWithoutDecoration(g.output.Formatter(), message)

	g.writeList(
		[]string{
			fmt.Sprintf("<comment>%s</>", message),
			fmt.Sprintf("<comment>%s</>", strings.Repeat("=", messageRealLength)),
		},
		true,
	)

	g.PrintNewLine(1)
}

// PrintSection formats and print a section title.
func (g *Styler) PrintSection(message string) {
	g.autoPrependBlock()

	messageRealLength := helper.StrlenWithoutDecoration(g.output.Formatter(), message)

	g.writeList(
		[]string{
			fmt.Sprintf("<comment>%s</>", message),
			fmt.Sprintf("<comment>%s</>", strings.Repeat("-", messageRealLength)),
		},
		true,
	)

	g.PrintNewLine(1)
}

// PrintListing formats and print a list.
func (g *Styler) PrintListing(messages []string) {
	g.autoPrependText()

	for _, msg := range messages {
		g.write(fmt.Sprintf(" * %s", msg), true)
	}

	g.PrintNewLine(1)
}

// PrintText formats and print informational text.
func (g *Styler) PrintText(message string) {
	g.autoPrependText()
	g.write(fmt.Sprintf("%s", message), false)
	g.PrintNewLine(1)
}

// PrintTexts formats and print informational text array.
func (g *Styler) PrintTexts(messages []string) {
	g.autoPrependText()

	for _, msg := range messages {
		g.write(fmt.Sprintf("%s", msg), true)
	}

	g.PrintNewLine(1)
}

// PrintComment formats and print a comment bar.
func (g *Styler) PrintComment(message string) {
	g.PrintComments([]string{message})
}

// PrintComments formats and print a comment bar.
func (g *Styler) PrintComments(messages []string) {
	g.blockList(messages, "", "", "<fg=default;bg=default> // </>", false, false)
}

// PrintSuccess formats and print a success result bar.
func (g *Styler) PrintSuccess(message string) {
	g.PrintSuccesses([]string{message})
}

// PrintSuccesses formats and print a success result bar.
func (g *Styler) PrintSuccesses(messages []string) {
	g.blockList(messages, "OK", "fg=black;bg=green", " ", true, false)
}

func (g *Styler) PrintError(message string) {
	g.PrintErrors([]string{message})
}

// PrintErrors formats and print an error result bar.
func (g *Styler) PrintErrors(messages []string) {
	g.blockList(messages, "ERROR", "fg=white;bg=red", " ", true, false)
}

// PrintWarning formats and print an warning result bar.
func (g *Styler) PrintWarning(message string) {
	g.PrintWarnings([]string{message})
}

// PrintWarnings formats and print an warning result bar.
func (g *Styler) PrintWarnings(messages []string) {
	g.blockList(messages, "WARNING", "fg=white;bg=red", " ", true, false)
}

// PrintNote formats and print a note.
func (g *Styler) PrintNote(message string) {
	g.PrintNotes([]string{message})
}

// PrintNotes formats and print a note.
func (g *Styler) PrintNotes(messages []string) {
	g.blockList(messages, "NOTE", "fg=yellow", " ! ", false, false)
}

// PrintCaution formats and print a caution.
func (g *Styler) PrintCaution(message string) {
	g.PrintCautions([]string{message})
}

// PrintCautions formats and print a caution.
func (g *Styler) PrintCautions(messages []string) {
	g.blockList(messages, "CAUTION", "fg=white;bg=red", " ! ", true, false)
}

//
// internal
//

func (g *Styler) write(message string, newLine bool) {
	if newLine {
		g.output.Println(message)
		g.bufferedOutput.Println(message)
	} else {
		g.output.Print(message)
		g.bufferedOutput.Print(message)
	}
}

func (g *Styler) writeList(messages []string, newLine bool) {
	for _, message := range messages {
		g.write(message, newLine)
	}
}

//
// Prepend
//

func (g *Styler) autoPrependBlock() {
	fetched := g.bufferedOutput.Fetch()

	if len(fetched) == 0 {
		g.PrintNewLine(1)
		return
	}

	if len(fetched) == 1 {
		if fetched[1:] == "\n" {
			g.PrintNewLine(1)
		}

		return
	}

	g.PrintNewLine(2 - strings.Count(fetched[2:], "\n"))
}

func (g *Styler) autoPrependText() {
	fetched := g.bufferedOutput.Fetch()

	if len(fetched) != 0 && "\n" == fetched[1:] {
		g.PrintNewLine(1)
	}
}

//
// block internal
//

func (g *Styler) block(message string, title string, style string, prefix string, padding bool, escape bool) {
	g.autoPrependBlock()
	g.writeList(g.createBlock(message, title, style, prefix, padding, escape), false)
	g.PrintNewLine(1)
}

func (g *Styler) createBlock(message string, title string, style string, prefix string, padding bool, escape bool) []string {
	return g.createBlockList([]string{message}, title, style, prefix, padding, escape)
}

func (g *Styler) blockList(message []string, title string, style string, prefix string, padding bool, escape bool) {
	g.autoPrependBlock()
	g.writeList(g.createBlockList(message, title, style, prefix, padding, escape), true)
	g.PrintNewLine(1)
}

func (g *Styler) createBlockList(messages []string, title string, style string, prefix string, padding bool, escape bool) []string {
	indentLength := 0
	prefixLength := helper.StrlenWithoutDecoration(g.output.Formatter(), prefix)

	lineIndentation := ""

	var lines []string

	if "" != title {
		title = fmt.Sprintf("[%s] ", title)
		indentLength = len(title)
		lineIndentation = strings.Repeat(" ", indentLength)
	}

	for key, message := range messages {
		if escape {
			message = formatter.Escape(message)
		}

		lines = append(
			lines,
			strings.Split(
				helper.Wordwrap(message, g.maxLineLength-prefixLength-indentLength, '\n'),
				"\n",
			)...,
		)

		if len(messages) > 1 && key < len(messages)-1 {
			lines = append(lines, "")
		}
	}

	firstLineIndex := 0

	if padding && g.output.IsDecorated() {
		firstLineIndex = 1
		lines = helper.ArrayUnshift(lines, "")
		lines = append(lines, "")
	}

	var formattedLines []string

	for i, line := range lines {
		if "" != title {
			if firstLineIndex == i {
				line = fmt.Sprintf("%s%s", title, line)
			} else {
				line = fmt.Sprintf("%s%s", lineIndentation, line)
			}
		}

		line = fmt.Sprintf("%s%s", prefix, line)
		line = fmt.Sprintf("%s%s", line, strings.Repeat(" ", g.maxLineLength-helper.StrlenWithoutDecoration(g.output.Formatter(), line)))

		if "" != style {
			line = fmt.Sprintf("<%s>%s</>", style, line)
		}

		formattedLines = append(formattedLines, line)
	}

	return formattedLines
}

func (g *Styler) Verbosity() verbosity.Level {
	return g.output.Verbosity()
}

func (g *Styler) IsQuiet() bool {
	return g.output.IsQuiet()
}

func (g *Styler) IsVerbose() bool {
	return g.output.IsVerbose()
}

func (g *Styler) IsVeryVerbose() bool {
	return g.output.IsVeryVerbose()
}

func (g *Styler) IsDebug() bool {
	return g.output.IsDebug()
}
