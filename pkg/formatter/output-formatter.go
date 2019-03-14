package formatter

import (
	"fmt"
)

type OutputFormatter struct {
	decorated bool
	styleStack OutputFormatterStyleStack
	stylesCache map[string]OutputFormatterStyle
}

// Sets the decorated flag.
func (o *OutputFormatter) SetDecorated(decorated bool) {
	o.decorated = decorated
}

// Gets the decorated flag.
func (o *OutputFormatter) IsDecorated() bool {
	return o.decorated
}

// Sets a new style to cache.
func (o *OutputFormatter) SetStyle(name string, style OutputFormatterStyle) {
	o.stylesCache[name] = style
}

// Gets style from cache with specified name.
func (o *OutputFormatter) GetStyle(name string) *OutputFormatterStyle {
	if style, ok := o.stylesCache[name]; ok {
		return &style
	}

	return nil
}

// Gets style stack
func (o *OutputFormatter) GetStyleStack() OutputFormatterStyleStack {
	return o.styleStack
}

// Checks if output formatter has style in cache with specified name.
func (o *OutputFormatter) HasStyle(name string) bool {
	style := o.GetStyle(name)
	return nil != style
}

// Formats a message according to the given styles.
func (o *OutputFormatter) Format(message string) string {
	return o.formatAndWrap(message, 0)
}

// Formats a message according to the given styles, wrapping at `$width` (0 means no wrapping).
func (o *OutputFormatter) formatAndWrap(message string, width int) string {
	offset := 0
	currentLineLenght := 0

	output := ""

	output = fmt.Sprintf(
		"%s%s",
		output,
		o.applyCurrentStyle(output[:offset], output, width, currentLineLenght),
	)

	return output
}

// Applies current style from stack to text, if must be applied.
func (o *OutputFormatter) applyCurrentStyle(text string, current string, width int, currentLineLength int) string {
	if "" == text {
		return ""
	}

	if 0 == width {
		if o.IsDecorated() {
			stack := o.GetStyleStack()
			style := stack.GetCurrent()
			return style.Apply(text)
		}

		return text
	}

	// TODO boxing with currentLineLength

	return text
}
