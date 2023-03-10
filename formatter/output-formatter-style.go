package formatter

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/color"
	"sort"
	"strings"
)

// Formatter style constructor
func NewOutputFormatterStyle(
	foreground string,
	background string,
	options []string,
) *OutputFormatterStyle {
	style := &OutputFormatterStyle{
		options: new(map[string]color.Color),
	}

	style.SetForeground(foreground)
	style.SetBackground(background)

	if nil != options {
		style.SetOptions(options)
	}

	return style
}

// Formatter style class for defining styles
type OutputFormatterStyle struct {
	foreground *color.Color
	background *color.Color
	options    *map[string]color.Color
}

// Sets style foreground color.
func (style *OutputFormatterStyle) SetForeground(name string) {
	if color.Null == name {
		style.foreground = nil
		return
	}

	foreground := color.ForegroundColor(name)
	style.foreground = &foreground
}

// Sets style background color.
func (style *OutputFormatterStyle) SetBackground(name string) {
	if color.Null == name {
		style.background = nil
		return
	}

	background := color.BackgroundColor(name)
	style.background = &background
}

// Sets multiple style options at once.
func (style *OutputFormatterStyle) SetOptions(options []string) {
	style.options = &map[string]color.Color{}

	for _, name := range options {
		style.SetOption(name)
	}
}

// Sets some specific style option.
func (style *OutputFormatterStyle) SetOption(name string) {
	(*style.options)[name] = color.Option(name)
}

// Unsets some specific style option.
func (style *OutputFormatterStyle) UnsetOption(name string) {
	if _, ok := (*style.options)[name]; ok {
		delete(*style.options, name)
		return
	}

	panic(errors.New("cannot unset undefined options"))
}

// Applies the style to a given text.
func (style *OutputFormatterStyle) Apply(text string) string {
	var setCode, unsetCode []int

	if nil != style.foreground {
		setCode = append(setCode, style.foreground.Value())
		unsetCode = append(unsetCode, style.foreground.Unset())
	}

	if nil != style.background {
		setCode = append(setCode, style.background.Value())
		unsetCode = append(unsetCode, style.background.Unset())
	}

	if 0 != len(*style.options) {
		sortedOptions := sortOptionsMapByStringKey(*style.options)

		for _, option := range sortedOptions {
			setCode = append(setCode, option.Value())
			unsetCode = append(unsetCode, option.Unset())
		}
	}

	if 0 == len(setCode) {
		fmt.Printf("")
		return text
	}

	setCodeString := arrayToString(setCode, ";")
	unsetCodeString := arrayToString(unsetCode, ";")

	result := fmt.Sprintf("\033[%sm%s\033[%sm", setCodeString, text, unsetCodeString)

	return result
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func sortOptionsMapByStringKey(m map[string]color.Color) []color.Color {
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var sorted []color.Color

	for _, key := range keys {
		sorted = append(sorted, m[key])
	}

	return sorted
}
