package go_console

import (
	"fmt"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/DrSmithFr/go-console/helper"
	"github.com/DrSmithFr/go-console/input"
	"github.com/DrSmithFr/go-console/output"
	"github.com/DrSmithFr/go-console/verbosity"
	"strings"
)

// terminal max line length
const MAX_LINE_LENGTH = 120

type abstractStyler struct {
	lineLength     int
	in             input.InputInterface
	out            output.OutputInterface
	bufferedOutput output.BufferedOutput
}

// set max line length
func (g *abstractStyler) SetMaxLineLength(length int) {
	g.lineLength = length
}

// get max line length
func (g *abstractStyler) MaxLineLength() int {
	return g.lineLength
}

// get current inputInterface Instance
func (g *abstractStyler) Input() input.InputInterface {
	return g.in
}

// get current outputInterface Instance
func (g *abstractStyler) Output() output.OutputInterface {
	return g.out
}

// Add newline(s).
func (g *abstractStyler) PrintNewLine(count int) {
	g.writeList([]string{strings.Repeat("\n", count)}, false)
}

// Formats a command title.
func (g *abstractStyler) PrintTitle(message string) {
	g.autoPrependBlock()

	messageRealLength := helper.StrlenWithoutDecoration(g.out.Formatter(), message)

	g.writeList(
		[]string{
			fmt.Sprintf("<comment>%s</>", message),
			fmt.Sprintf("<comment>%s</>", strings.Repeat("=", messageRealLength)),
		},
		true,
	)

	g.PrintNewLine(1)
}

// Formats a section title.
func (g *abstractStyler) PrintSection(message string) {
	g.autoPrependBlock()

	messageRealLength := helper.StrlenWithoutDecoration(g.out.Formatter(), message)

	g.writeList(
		[]string{
			fmt.Sprintf("<comment>%s</>", message),
			fmt.Sprintf("<comment>%s</>", strings.Repeat("-", messageRealLength)),
		},
		true,
	)

	g.PrintNewLine(1)
}

// Formats a list.
func (g *abstractStyler) PrintListing(messages []string) {
	g.autoPrependText()

	for _, msg := range messages {
		g.write(fmt.Sprintf(" * %s", msg), true)
	}

	g.PrintNewLine(1)
}

// Formats informational text.
func (g *abstractStyler) PrintText(message string) {
	g.autoPrependText()
	g.write(fmt.Sprintf(" %s", message), false)
	g.PrintNewLine(1)
}

// Formats informational text array.
func (g *abstractStyler) PrintTexts(messages []string) {
	g.autoPrependText()

	for _, msg := range messages {
		g.write(fmt.Sprintf(" %s", msg), true)
	}

	g.PrintNewLine(1)
}

// Formats a comment bar.
func (g *abstractStyler) PrintComment(message string) {
	g.PrintComments([]string{message})
}

// Formats a comment bar.
func (g *abstractStyler) PrintComments(messages []string) {
	g.blockList(messages, "", "", "<fg=default;bg=default> // </>", false, false)
}

// Formats a success result bar.
func (g *abstractStyler) PrintSuccess(message string) {
	g.PrintSuccesses([]string{message})
}

// Formats a success result bar.
func (g *abstractStyler) PrintSuccesses(messages []string) {
	g.blockList(messages, "OK", "fg=black;bg=green", " ", true, false)
}

func (g *abstractStyler) PrintError(message string) {
	g.PrintErrors([]string{message})
}

// Formats an error result bar.
func (g *abstractStyler) PrintErrors(messages []string) {
	g.blockList(messages, "ERROR", "fg=white;bg=red", " ", true, false)
}

// Formats an warning result bar.
func (g *abstractStyler) PrintWarning(message string) {
	g.PrintWarnings([]string{message})
}

// Formats an warning result bar.
func (g *abstractStyler) PrintWarnings(messages []string) {
	g.blockList(messages, "WARNING", "fg=white;bg=red", " ", true, false)
}

// Formats a note admonition.
func (g *abstractStyler) PrintNote(message string) {
	g.PrintNotes([]string{message})
}

// Formats a note admonition.
func (g *abstractStyler) PrintNotes(messages []string) {
	g.blockList(messages, "NOTE", "fg=yellow", " ! ", false, false)
}

// Formats a caution admonition.
func (g *abstractStyler) PrintCaution(message string) {
	g.PrintCautions([]string{message})
}

// Formats a caution admonition.
func (g *abstractStyler) PrintCautions(messages []string) {
	g.blockList(messages, "CAUTION", "fg=white;bg=red", " ! ", true, false)
}

//
// internal
//

func (g *abstractStyler) write(message string, newLine bool) {
	if newLine {
		g.out.Writeln(message)
		g.bufferedOutput.Writeln(message)
	} else {
		g.out.Write(message)
		g.bufferedOutput.Write(message)
	}
}

func (g *abstractStyler) writeList(messages []string, newLine bool) {
	for _, message := range messages {
		g.write(message, newLine)
	}
}

//
// Prepend
//

func (g *abstractStyler) autoPrependBlock() {
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

func (g *abstractStyler) autoPrependText() {
	fetched := g.bufferedOutput.Fetch()

	if len(fetched) != 0 && "\n" == fetched[1:] {
		g.PrintNewLine(1)
	}
}

//
// block internal
//

func (g *abstractStyler) block(message string, title string, style string, prefix string, padding bool, escape bool) {
	g.autoPrependBlock()
	g.writeList(g.createBlock(message, title, style, prefix, padding, escape), false)
	g.PrintNewLine(1)
}

func (g *abstractStyler) createBlock(message string, title string, style string, prefix string, padding bool, escape bool) []string {
	return g.createBlockList([]string{message}, title, style, prefix, padding, escape)
}

func (g *abstractStyler) blockList(message []string, title string, style string, prefix string, padding bool, escape bool) {
	g.autoPrependBlock()
	g.writeList(g.createBlockList(message, title, style, prefix, padding, escape), true)
	g.PrintNewLine(1)
}

func (g *abstractStyler) createBlockList(messages []string, title string, style string, prefix string, padding bool, escape bool) []string {
	indentLength := 0
	prefixLength := helper.StrlenWithoutDecoration(g.out.Formatter(), prefix)

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
				helper.Wordwrap(message, g.lineLength-prefixLength-indentLength, '\n'),
				"\n",
			)...,
		)

		if len(messages) > 1 && key < len(messages)-1 {
			lines = append(lines, "")
		}
	}

	firstLineIndex := 0

	if padding && g.out.IsDecorated() {
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
		line = fmt.Sprintf("%s%s", line, strings.Repeat(" ", g.lineLength-helper.StrlenWithoutDecoration(g.out.Formatter(), line)))

		if "" != style {
			line = fmt.Sprintf("<%s>%s</>", style, line)
		}

		formattedLines = append(formattedLines, line)
	}

	return formattedLines
}

func (g *abstractStyler) Verbosity() verbosity.Level {
	return g.out.Verbosity()
}

func (g *abstractStyler) IsQuiet() bool {
	return g.out.IsQuiet()
}

func (g *abstractStyler) IsVerbose() bool {
	return g.out.IsVerbose()
}

func (g *abstractStyler) IsVeryVerbose() bool {
	return g.out.IsVeryVerbose()
}

func (g *abstractStyler) IsDebug() bool {
	return g.out.IsDebug()
}
