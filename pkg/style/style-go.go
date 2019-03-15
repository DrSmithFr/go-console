package style

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/output"
	"strings"
	"unicode/utf8"
)

const MAX_LINE_LENGTH = 120

func NewGoStyler(out output.OutputInterface) *GoStyler {
	g := new(GoStyler)

	// clone the formatter to retrieve styles and avoid state change
	formatter := *out.GetFormatter()

	g.out = out
	g.bufferedOutput = *output.NewBufferedOutput(false, &formatter)

	return g
}

type GoStyler struct {
	out            output.OutputInterface
	bufferedOutput output.BufferedOutput
}

func (g *GoStyler) NewLine(count int) {
	g.autoPrependBlock()

	g.writeArray([]string{strings.Repeat("\n", count)}, false)
}

func (g *GoStyler) Title(message string) {
	g.autoPrependBlock()

	g.writeArray(
		[]string{
			fmt.Sprintf("<comment>%s</>", message),
			fmt.Sprintf("<comment>%s</>", strings.Repeat("=", utf8.RuneCountInString(message))),
		},
		true,
	)

	g.NewLine(1)
}

//
// internal
//

func (g *GoStyler) writeArray(messages []string, newLine bool) {
	for _, message := range messages {
		if newLine {
			g.out.Writeln(message)
			g.bufferedOutput.Writeln(message)
		} else {
			g.out.Write(message)
			g.bufferedOutput.Write(message)
		}
	}
}

func (g *GoStyler) autoPrependBlock() {
	// TODO
}

func (g *GoStyler) block(message string, title string, style string, prefix byte, padding bool, escape bool) {
	// TODO
}
