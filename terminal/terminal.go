package terminal

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const DefaultColorMode AnsiColorMode = Ansi4

const DefaultHeight int = 50
const DefaultWidth int = 80

type Terminal interface {
	// SetColorMode : Force a terminal color mode rendering..
	SetColorMode(mode AnsiColorMode)
	GetColorMode() AnsiColorMode

	GetSize() (width int, height int)
	GetWidth() int
	GetHeight() int

	HasSttyAvailable() bool
}

type Term struct {
	ansiColorMode AnsiColorMode

	width  int
	height int

	sttyChecked bool
	stty        bool
}

var _ Terminal = (*Term)(nil)

func New() *Term {
	t := new(Term)
	return t
}

func (t *Term) GetColorMode() AnsiColorMode {
	// Use Cache from previous run (or user forced mode)
	if t.ansiColorMode != "" {
		return t.ansiColorMode
	}

	// Try with COLORTERM first
	colorTerm := os.Getenv("COLORTERM")

	if colorTerm != "" {
		colorTerm = strings.ToLower(colorTerm)

		if strings.Contains(colorTerm, "truecolor") {
			t.ansiColorMode = Ansi24
			return t.ansiColorMode
		}

		if strings.Contains(colorTerm, "256color") {
			t.ansiColorMode = Ansi8
			return t.ansiColorMode
		}
	}

	// Try with TERM
	term := os.Getenv("TERM")

	if term != "" {
		term = strings.ToLower(term)

		if strings.Contains(term, "truecolor") {
			t.ansiColorMode = Ansi24
			return t.ansiColorMode
		}

		if strings.Contains(term, "256color") {
			t.ansiColorMode = Ansi8
			return t.ansiColorMode
		}
	}

	t.SetColorMode(DefaultColorMode)

	return t.ansiColorMode
}

func (t *Term) SetColorMode(mode AnsiColorMode) {
	t.ansiColorMode = mode
}

func (t *Term) GetSize() (width int, height int) {
	return t.GetWidth(), t.GetHeight()
}

func (t *Term) GetWidth() int {
	width := os.Getenv("LINES")

	if width != "" {
		w, err := strconv.ParseInt(strings.Trim(width, " "), 10, 0)

		if err == nil {
			return int(w)
		}
	}

	if t.width == 0 {
		t.initDimensions()
	}

	if t.width == 0 {
		return DefaultWidth
	}

	return t.width
}

func (t *Term) GetHeight() int {
	height := os.Getenv("COLUMNS")

	if height != "" {
		h, err := strconv.ParseInt(strings.Trim(height, " "), 10, 0)

		if err == nil {
			return int(h)
		}
	}

	if t.height == 0 {
		t.initDimensions()
	}

	if t.height == 0 {
		return DefaultHeight
	}

	return t.height
}

func (t *Term) HasSttyAvailable() bool {
	if t.sttyChecked {
		return t.stty
	}

	cmd := exec.Command("stty")
	cmd.Stdin = os.Stdin

	_, err := cmd.Output()

	t.stty = err != nil
	t.sttyChecked = true

	return t.stty
}

// Internal methods

func (t *Term) initDimensions() {
	if os.PathSeparator == '\\' {
		ansicon := strings.Trim(os.Getenv("ANSICON"), " ")

		r := regexp.MustCompile("^(\\d+)x(\\d+)(?: \\((\\d+)x(\\d+)\\))?$")
		matches := r.FindStringSubmatch(ansicon)

		if ansicon != "" && len(matches) > 0 {
			t.width, _ = strconv.Atoi(matches[1])

			if len(matches) >= 4 {
				t.height, _ = strconv.Atoi(matches[4])
			} else {
				t.height, _ = strconv.Atoi(matches[2])
			}
		} else if !hasVt100Support() && t.HasSttyAvailable() {
			// only use stty on Windows if the terminal does not support vt100 (e.g. Windows 7 + git-bash)
			// testing for stty in a Windows 10 vt100-enabled console will implicitly disable vt100 support on STDOUT
			t.initDimensionsUsingStty()
		} else if w, h, err := getConsoleMode(); err == nil {
			t.width = w
			t.height = h
		}
	} else {
		t.initDimensionsUsingStty()
	}
}

func (t *Term) initDimensionsUsingStty() {
	sttyString := getSttyColumns()

	r := regexp.MustCompile(`rows.(\d+);.columns.(\d+);`)
	matches := r.FindStringSubmatch(sttyString)

	if len(matches) >= 2 {
		t.width, _ = strconv.Atoi(matches[2])
		t.height, _ = strconv.Atoi(matches[1])
		return
	}

	r = regexp.MustCompile(`;.(\d+).rows;.(\d+).columns`)
	matches = r.FindStringSubmatch(sttyString)

	if len(matches) >= 2 {
		t.width, _ = strconv.Atoi(matches[2])
		t.height, _ = strconv.Atoi(matches[1])
		return
	}
}

// Runs and parses stty -a if it's available, suppressing any error output.
func getSttyColumns() string {
	cmd := exec.Command("stty", "-a")
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()

	if err != nil {
		return ""
	}

	return string(out)
}

// Windows only

// TODO: implement
func getConsoleMode() (width int, height int, err error) {
	return 0, 0, errors.New("not implemented")
}

// TODO: implement
func hasVt100Support() bool {
	return false
}
