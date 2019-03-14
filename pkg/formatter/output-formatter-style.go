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

	style.options = &map[string]color.Color{}

	return style
}

type OutputFormatterStyle struct {
	foreground *color.Color
	background *color.Color
	options    *map[string]color.Color
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

// Sets multiple style options at once.
func (style *OutputFormatterStyle) SetOptions(options []string) {
	for _, name := range options {
		style.SetOption(name)
	}
}

// Sets some specific style option.
func (style *OutputFormatterStyle) SetOption(name string) {
	(*style.options)[name] = color.GetOption(name)
}

// Unsets some specific style option.
func (style *OutputFormatterStyle) UnsetOption(name string) {
	if _, ok := (*style.options)[name]; ok {
		delete(*style.options, name)
	}
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

	if 0 != len(*style.options) {
		for _, option := range *style.options {
			setCode = append(setCode, option.GetValue())
			unsetCode = append(unsetCode, option.GetUnset())
		}
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
