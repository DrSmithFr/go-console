package style

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/formatter"
	"github.com/DrSmithFr/go-console/pkg/helper"
	"github.com/DrSmithFr/go-console/pkg/output"
	"strings"
)

const MAX_LINE_LENGTH = 120

func NewConsoleGoStyler() *GoStyler {
	out := output.NewConsoleOutput(true, nil)
	return NewGoStyler(out)
}

func NewGoStyler(out output.OutputInterface) *GoStyler {
	g := new(GoStyler)

	// clone the formatter to retrieve styles and avoid state change
	format := *out.GetFormatter()

	g.lineLength = MAX_LINE_LENGTH
	g.out = out
	g.bufferedOutput = *output.NewBufferedOutput(false, &format)

	return g
}

type GoStyler struct {
	lineLength     int
	out            output.OutputInterface
	bufferedOutput output.BufferedOutput
}

func (g *GoStyler) SetMaxLineLength(length int) {
	g.lineLength = length
}

func (g *GoStyler) GetMaxLineLength() int {
	return g.lineLength
}

func (g *GoStyler) GetOutput() output.OutputInterface {
	return g.out
}

func (g *GoStyler) NewLine(count int) {
	g.writeArray([]string{strings.Repeat("\n", count)}, false)
}

func (g *GoStyler) Title(message string) {
	g.autoPrependBlock()

	messageRealLength := helper.StrlenWithoutDecoration(g.out.GetFormatter(), message)

	g.writeArray(
		[]string{
			fmt.Sprintf("<comment>%s</>", message),
			fmt.Sprintf("<comment>%s</>", strings.Repeat("=", messageRealLength)),
		},
		true,
	)

	g.NewLine(1)
}

func (g *GoStyler) Section(message string) {
	g.autoPrependBlock()

	messageRealLength := helper.StrlenWithoutDecoration(g.out.GetFormatter(), message)

	g.writeArray(
		[]string{
			fmt.Sprintf("<comment>%s</>", message),
			fmt.Sprintf("<comment>%s</>", strings.Repeat("-", messageRealLength)),
		},
		true,
	)

	g.NewLine(1)
}

func (g *GoStyler) Listing(messages []string) {
	g.autoPrependText()

	for _, msg := range messages {
		g.write(fmt.Sprintf(" * %s", msg), true)
	}

	g.NewLine(1)
}

func (g *GoStyler) Text(message string) {
	g.autoPrependText()
	g.write(fmt.Sprintf(" %s", message), false)
	g.NewLine(1)
}

func (g *GoStyler) TextArray(messages []string) {
	g.autoPrependText()

	for _, msg := range messages {
		g.write(fmt.Sprintf(" %s", msg), true)
	}

	g.NewLine(1)
}

func (g *GoStyler) Comment(message string) {
	g.CommentArray([]string{message})
}

func (g *GoStyler) CommentArray(messages []string) {
	g.blockArray(messages, "", "", "<fg=default;bg=default> // </>", false, false)
}

func (g *GoStyler) Success(message string) {
	g.SuccessArray([]string{message})
}

func (g *GoStyler) SuccessArray(messages []string) {
	g.blockArray(messages, "OK", "fg=black;bg=green", " ", true, false)
}

func (g *GoStyler) Error(message string) {
	g.ErrorArray([]string{message})
}

func (g *GoStyler) ErrorArray(messages []string) {
	g.blockArray(messages, "ERROR", "fg=white;bg=red", " ", true, false)
}

func (g *GoStyler) Warning(message string) {
	g.WarningArray([]string{message})
}

func (g *GoStyler) WarningArray(messages []string) {
	g.blockArray(messages, "WARNING", "fg=white;bg=red", " ", true, false)
}

func (g *GoStyler) Note(message string) {
	g.NoteArray([]string{message})
}

func (g *GoStyler) NoteArray(messages []string) {
	g.blockArray(messages, "NOTE", "fg=yellow", " ! ", false, false)
}

func (g *GoStyler) Caution(message string) {
	g.CautionArray([]string{message})
}

func (g *GoStyler) CautionArray(messages []string) {
	g.blockArray(messages, "CAUTION", "fg=white;bg=red", " ! ", true, false)
}

//
// internal
//

func (g *GoStyler) write(message string, newLine bool) {
	if newLine {
		g.out.Writeln(message)
		g.bufferedOutput.Writeln(message)
	} else {
		g.out.Write(message)
		g.bufferedOutput.Write(message)
	}
}

func (g *GoStyler) writeArray(messages []string, newLine bool) {
	for _, message := range messages {
		g.write(message, newLine)
	}
}

//
// Prepend
//

func (g *GoStyler) autoPrependBlock() {
	fetched := g.bufferedOutput.Fetch()

	if len(fetched) == 0 {
		g.NewLine(1)
		return
	}

	if len(fetched) == 1 {
		if fetched[1:] == "\n" {
			g.NewLine(1)
		}

		return
	}

	g.NewLine(2 - strings.Count(fetched[2:], "\n"))
}

func (g *GoStyler) autoPrependText() {
	fetched := g.bufferedOutput.Fetch()

	if len(fetched) != 0 && "\n" == fetched[1:] {
		g.NewLine(1)
	}
}

//
// block internal
//

func (g *GoStyler) block(message string, title string, style string, prefix string, padding bool, escape bool) {
	g.autoPrependBlock()
	g.writeArray(g.createBlock(message, title, style, prefix, padding, escape), false)
	g.NewLine(1)
}

func (g *GoStyler) createBlock(message string, title string, style string, prefix string, padding bool, escape bool) []string {
	return g.createBlockArray([]string{message}, title, style, prefix, padding, escape)
}

func (g *GoStyler) blockArray(message []string, title string, style string, prefix string, padding bool, escape bool) {
	g.autoPrependBlock()
	g.writeArray(g.createBlockArray(message, title, style, prefix, padding, escape), true)
	g.NewLine(1)
}

func (g *GoStyler) createBlockArray(messages []string, title string, style string, prefix string, padding bool, escape bool) []string {
	indentLength := 0
	prefixLength := helper.StrlenWithoutDecoration(g.out.GetFormatter(), prefix)

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
		line = fmt.Sprintf("%s%s", line, strings.Repeat(" ", g.lineLength-helper.StrlenWithoutDecoration(g.out.GetFormatter(), line)))

		if "" != style {
			line = fmt.Sprintf("<%s>%s</>", style, line)
		}

		formattedLines = append(formattedLines, line)
	}

	return formattedLines
}
