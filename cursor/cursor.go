package cursor

import (
	"fmt"
	"github.com/DrSmithFr/go-console/output"
	"github.com/DrSmithFr/go-console/terminal"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Cursor struct {
	output output.OutputInterface
	input  *io.ReadWriter
	term   *terminal.Term
}

func NewCursor(output output.OutputInterface, input io.ReadWriter) *Cursor {
	if input == nil {
		input = os.Stdin
	}

	return &Cursor{
		output: output,
		input:  &input,
		term:   terminal.New(),
	}
}

func (c *Cursor) write(str string) {
	_, err := c.output.Write([]byte(str))

	if err != nil {
		panic(err)
	}
}

func (c *Cursor) MoveUp(lines int) *Cursor {
	c.write(fmt.Sprintf("\x1b[%dA", lines))
	return c
}

func (c *Cursor) MoveDown(lines int) *Cursor {
	c.write(fmt.Sprintf("\x1b[%dB", lines))
	return c
}

func (c *Cursor) MoveRight(columns int) *Cursor {
	c.write(fmt.Sprintf("\x1b[%dC", columns))
	return c
}

func (c *Cursor) MoveLeft(columns int) *Cursor {
	c.write(fmt.Sprintf("\x1b[%dD", columns))
	return c
}

func (c *Cursor) MoveToColumn(column int) *Cursor {
	c.write(fmt.Sprintf("\x1b[%dG", column))
	return c
}

func (c *Cursor) MoveToPosition(line, column int) *Cursor {
	c.write(fmt.Sprintf("\x1b[%d;%dH", line, column))
	return c
}

func (c *Cursor) SavePosition() *Cursor {
	c.write("\x1b7")
	return c
}

func (c *Cursor) RestorePosition() *Cursor {
	c.write("\x1b8")
	return c
}

func (c *Cursor) Hide() *Cursor {
	c.write("\x1b[?25l")
	return c
}

func (c *Cursor) Show() *Cursor {
	c.write("\x1b[?25h\x1b[?0c")
	return c
}

func (c *Cursor) ClearLine() *Cursor {
	c.write("\x1b[2K")
	return c
}

func (c *Cursor) ClearLineAfter() *Cursor {
	c.write("\x1b[K")
	return c
}

func (c *Cursor) ClearOutput() *Cursor {
	c.write("\x1b[0J")
	return c
}

func (c *Cursor) ClearScreen() *Cursor {
	c.write("\x1b[2J")
	return c
}

func (c *Cursor) GetCurrentPosition() (line, column int) {
	stty := c.term.HasSttyAvailable()
	if !stty {
		return 1, 1
	}

	cmd := exec.Command("stty", "-g")
	cmd.Stdin = os.Stdin

	sttyMode, _ := cmd.Output()

	cmd = exec.Command("stty", "-icanon", "-echo")
	cmd.Stdin = os.Stdin
	_ = cmd.Run()

	in := *(c.input)
	_, _ = in.Write([]byte("\033[6n"))

	rawCode := make([]byte, 1024)
	_, _ = in.Read(rawCode)
	code := strings.TrimSpace(string(rawCode))

	cmd = exec.Command("stty", string(sttyMode))
	cmd.Stdin = os.Stdin
	_ = cmd.Run()

	_, _ = fmt.Sscanf(code, "\033[%d;%dR", &line, &column)
	return line, column
}
