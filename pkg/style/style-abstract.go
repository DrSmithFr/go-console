package style

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/formatter"
	"github.com/DrSmithFr/go-console/pkg/helper"
	"github.com/DrSmithFr/go-console/pkg/input"
	"github.com/DrSmithFr/go-console/pkg/output"
	"strings"
)

const MAX_LINE_LENGTH = 120

type abstractStyler struct {
	lineLength     int
	in             input.InputInterface
	out            output.OutputInterface
	bufferedOutput output.BufferedOutput
}

func (g *abstractStyler) SetMaxLineLength(length int) {
	g.lineLength = length
}

func (g *abstractStyler) GetMaxLineLength() int {
	return g.lineLength
}

func (g *abstractStyler) GetInput() input.InputInterface {
	return g.in
}

func (g *abstractStyler) GetOutput() output.OutputInterface {
	return g.out
}

func (g *abstractStyler) NewLine(count int) {
	g.writeArray([]string{strings.Repeat("\n", count)}, false)
}

func (g *abstractStyler) Title(message string) {
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

func (g *abstractStyler) Section(message string) {
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

func (g *abstractStyler) Listing(messages []string) {
	g.autoPrependText()

	for _, msg := range messages {
		g.write(fmt.Sprintf(" * %s", msg), true)
	}

	g.NewLine(1)
}

func (g *abstractStyler) Text(message string) {
	g.autoPrependText()
	g.write(fmt.Sprintf(" %s", message), false)
	g.NewLine(1)
}

func (g *abstractStyler) TextArray(messages []string) {
	g.autoPrependText()

	for _, msg := range messages {
		g.write(fmt.Sprintf(" %s", msg), true)
	}

	g.NewLine(1)
}

func (g *abstractStyler) Comment(message string) {
	g.CommentArray([]string{message})
}

func (g *abstractStyler) CommentArray(messages []string) {
	g.blockArray(messages, "", "", "<fg=default;bg=default> // </>", false, false)
}

func (g *abstractStyler) Success(message string) {
	g.SuccessArray([]string{message})
}

func (g *abstractStyler) SuccessArray(messages []string) {
	g.blockArray(messages, "OK", "fg=black;bg=green", " ", true, false)
}

func (g *abstractStyler) Error(message string) {
	g.ErrorArray([]string{message})
}

func (g *abstractStyler) ErrorArray(messages []string) {
	g.blockArray(messages, "ERROR", "fg=white;bg=red", " ", true, false)
}

func (g *abstractStyler) Warning(message string) {
	g.WarningArray([]string{message})
}

func (g *abstractStyler) WarningArray(messages []string) {
	g.blockArray(messages, "WARNING", "fg=white;bg=red", " ", true, false)
}

func (g *abstractStyler) Note(message string) {
	g.NoteArray([]string{message})
}

func (g *abstractStyler) NoteArray(messages []string) {
	g.blockArray(messages, "NOTE", "fg=yellow", " ! ", false, false)
}

func (g *abstractStyler) Caution(message string) {
	g.CautionArray([]string{message})
}

func (g *abstractStyler) CautionArray(messages []string) {
	g.blockArray(messages, "CAUTION", "fg=white;bg=red", " ! ", true, false)
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

func (g *abstractStyler) writeArray(messages []string, newLine bool) {
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

func (g *abstractStyler) autoPrependText() {
	fetched := g.bufferedOutput.Fetch()

	if len(fetched) != 0 && "\n" == fetched[1:] {
		g.NewLine(1)
	}
}

//
// block internal
//

func (g *abstractStyler) block(message string, title string, style string, prefix string, padding bool, escape bool) {
	g.autoPrependBlock()
	g.writeArray(g.createBlock(message, title, style, prefix, padding, escape), false)
	g.NewLine(1)
}

func (g *abstractStyler) createBlock(message string, title string, style string, prefix string, padding bool, escape bool) []string {
	return g.createBlockArray([]string{message}, title, style, prefix, padding, escape)
}

func (g *abstractStyler) blockArray(message []string, title string, style string, prefix string, padding bool, escape bool) {
	g.autoPrependBlock()
	g.writeArray(g.createBlockArray(message, title, style, prefix, padding, escape), true)
	g.NewLine(1)
}

func (g *abstractStyler) createBlockArray(messages []string, title string, style string, prefix string, padding bool, escape bool) []string {
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
