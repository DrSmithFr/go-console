package formatter

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/color"
	"strings"
)

func NewOutputFormatterStyle(foreground string, background string) *OutputFormatterStyle {
	style := new(OutputFormatterStyle)

	style.SetForeground(&foreground)
	style.SetBackground(&background)

	return style
}

type OutputFormatterStyle struct {
	foreground *color.Color
	background *color.Color
}

// Sets style foreground color.
func (style *OutputFormatterStyle) SetForeground(name *string) {
	if nil == name {
		style.foreground = nil
		return
	}

	foreground := color.GetForegroundColor(*name)
	style.foreground = &foreground
}

// Sets style background color.
func (style *OutputFormatterStyle) SetBackground(name *string) {
	if nil == name {
		style.background = nil
		return
	}

	background := color.GetBackgroundColor(*name)
	style.background = &background
}

// Applies the style to a given text.
func (style *OutputFormatterStyle) Apply(text string) string {
	var setCode, unsetCode []int

	if nil != style.foreground {
		setCode = append(setCode, style.foreground.GetValue())
		unsetCode = append(unsetCode, style.foreground.GetUnset())
	}

	if nil != style.background {
		setCode = append(setCode, style.background.GetValue())
		unsetCode = append(unsetCode, style.background.GetUnset())
	}

	return fmt.Sprintf(
		"\033[%sm%s\033[%sm",
		arrayToString(setCode, ";"),
		text,
		arrayToString(unsetCode, ";"),
	)
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}