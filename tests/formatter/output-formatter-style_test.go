package formatter

import (
	"DrSmithFr/go-console/pkg/color"
	"DrSmithFr/go-console/pkg/formatter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	s1 := formatter.NewOutputFormatterStyle(color.GREEN, color.BLACK, []string{color.BOLD, color.UNDERSCORE})
	assert.Equal(t, "\033[32;40;1;4mfoo\033[39;49;22;24m", s1.Apply("foo"))

	s2 := formatter.NewOutputFormatterStyle(color.RED, color.NULL, []string{color.BLINK})
	assert.Equal(t, "\033[31;5mfoo\033[39;25m", s2.Apply("foo"))

	s3 := formatter.NewOutputFormatterStyle(color.NULL, color.WHITE, nil)
	assert.Equal(t, "\033[47mfoo\033[49m", s3.Apply("foo"))
}

func TestForeground(t *testing.T) {
	s := formatter.NewOutputFormatterStyle(color.NULL, color.NULL, nil)

	s.SetForeground(color.BLACK)
	assert.Equal(t, "\033[30mfoo\033[39m", s.Apply("foo"))

	s.SetForeground(color.BLUE)
	assert.Equal(t, "\033[34mfoo\033[39m", s.Apply("foo"))

	s.SetForeground(color.DEFAULT)
	assert.Equal(t, "\033[39mfoo\033[39m", s.Apply("foo"))

	assert.Panics(t, func() {
		s.SetForeground("undefined-color")
	})
}

func TestBackground(t *testing.T) {
	s := formatter.NewOutputFormatterStyle(color.NULL, color.NULL, nil)

	s.SetBackground(color.BLACK)
	assert.Equal(t, "\033[40mfoo\033[49m", s.Apply("foo"))

	s.SetBackground(color.YELLOW)
	assert.Equal(t, "\033[43mfoo\033[49m", s.Apply("foo"))

	s.SetBackground(color.DEFAULT)
	assert.Equal(t, "\033[49mfoo\033[49m", s.Apply("foo"))

	assert.Panics(t, func() {
		s.SetBackground("undefined-color")
	})
}

func TestOptions(t *testing.T) {
	s := formatter.NewOutputFormatterStyle(color.NULL, color.NULL, nil)

	s.SetOptions([]string{color.REVERSE, color.CONCEAL})
	assert.Equal(t, "\033[8;7mfoo\033[28;27m", s.Apply("foo"))

	s.SetOption(color.BOLD)
	assert.Equal(t, "\033[1;8;7mfoo\033[22;28;27m", s.Apply("foo"))

	s.UnsetOption(color.REVERSE)
	assert.Equal(t, "\033[1;8mfoo\033[22;28m", s.Apply("foo"))

	s.SetOption(color.BOLD)
	assert.Equal(t, "\033[1;8mfoo\033[22;28m", s.Apply("foo"))

	s.SetOptions([]string{color.BOLD})
	assert.Equal(t, "\033[1mfoo\033[22m", s.Apply("foo"))

	assert.Panics(t, func() {
		s.SetOption("undefined-option")
	})

	assert.Panics(t, func() {
		s.UnsetOption("undefined-option")
	})
}
