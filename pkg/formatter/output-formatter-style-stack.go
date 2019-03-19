package formatter

import (
	"errors"
	"github.com/MrSmith777/go-console/pkg/color"
)

func NewOutputFormatterStyleStack(style *OutputFormatterStyle) *OutputFormatterStyleStack {
	stack := new(OutputFormatterStyleStack)

	if nil != style {
		stack.defaultStyle = style
	} else {
		stack.defaultStyle = NewOutputFormatterStyle(color.DEFAULT, color.DEFAULT, nil)
	}

	stack.styles = []*OutputFormatterStyle{}

	return stack
}

type OutputFormatterStyleStack struct {
	styles []*OutputFormatterStyle
	defaultStyle *OutputFormatterStyle
}

// Resets stack (ie. empty internal arrays).
func (stack *OutputFormatterStyleStack) Reset() {
	stack.styles = []*OutputFormatterStyle{}
}

// Pushes a style in the stack.
func (stack *OutputFormatterStyleStack) Push(style *OutputFormatterStyle) {
	stack.styles = append(stack.styles, style)
}

// Pops a style from the stack, pop the first if style equal nil.
func (stack *OutputFormatterStyleStack) Pop(style *OutputFormatterStyle) OutputFormatterStyle {
	if 0 == len(stack.styles) {
		return *stack.defaultStyle
	}

	if nil == style {
		first := stack.styles[len(stack.styles) - 1]
		newStack := stack.styles[:len(stack.styles) - 1]
		stack.styles = newStack
		return *first
	}

	for index := len(stack.styles) - 1; index >= 0; index-- {
		stackedStyle := stack.styles[index]

		currentStyleResult := (*style).Apply("")
		stackedStyleResult := stackedStyle.Apply("")

		if currentStyleResult == stackedStyleResult {
			stack.styles = stack.styles[:index]
			return *stackedStyle
		}
	}

	panic(errors.New("incorrectly nested style tag found"))
}

func (stack *OutputFormatterStyleStack) GetCurrent() *OutputFormatterStyle {
	if 0 == len(stack.styles) {
		return stack.defaultStyle
	}

	return stack.styles[len(stack.styles) - 1]
}

func (stack *OutputFormatterStyleStack) GetDefaultStyle() *OutputFormatterStyle {
	return stack.defaultStyle
}

func (stack *OutputFormatterStyleStack) SetDefaultStyle(style OutputFormatterStyle) {
	stack.defaultStyle = &style
}