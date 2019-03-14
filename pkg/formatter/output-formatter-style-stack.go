package formatter

import (
	"errors"
)

type OutputFormatterStyleStack struct {
	styles []OutputFormatterStyle
	defaultStyle OutputFormatterStyle
}

// Resets stack (ie. empty internal arrays).
func (stack *OutputFormatterStyleStack) Reset() {
	stack.styles = []OutputFormatterStyle{}
}

// Pushes a style in the stack.
func (stack *OutputFormatterStyleStack) Push(style OutputFormatterStyle) {
	stack.styles = append(stack.styles, style)
}

// Pops a style from the stack, pop the first if style equal nil.
func (stack *OutputFormatterStyleStack) Pop(style *OutputFormatterStyle) OutputFormatterStyle {
	if 0 == len(stack.styles) {
		return stack.defaultStyle
	}

	if nil == style {
		first := stack.styles[0]
		stack.styles = stack.styles[:1]
		return first
	}

	for index := len(stack.styles) - 1; index >= 0; index-- {
		stackedStyle := stack.styles[index]

		if (*style).Apply("") == stackedStyle.Apply("") {
			stack.styles = stack.styles[:index]
			return stackedStyle
		}
	}

	panic(errors.New("incorrectly nested style tag found"))
}

func (stack *OutputFormatterStyleStack) GetCurrent() OutputFormatterStyle {
	if 0 == len(stack.styles) {
		return stack.defaultStyle
	}

	return stack.styles[len(stack.styles) - 1]
}

func (stack *OutputFormatterStyleStack) GetDefaultStyle() OutputFormatterStyle {
	return stack.defaultStyle
}

func (stack *OutputFormatterStyleStack) SetDefaultStyle(style OutputFormatterStyle) {
	stack.defaultStyle = style
}