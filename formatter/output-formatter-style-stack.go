package formatter

import (
	"errors"
	"github.com/DrSmithFr/go-console/color"
)

// output formatter style constructor
func NewOutputFormatterStyleStack(style *OutputFormatterStyle) *OutputFormatterStyleStack {
	stack := &OutputFormatterStyleStack{
		styles: []*OutputFormatterStyle{},
	}

	if nil != style {
		stack.defaultStyle = style
	} else {
		stack.defaultStyle = NewOutputFormatterStyle(color.Null, color.Null, nil)
	}

	return stack
}

// Formatter style stack class for defining styles
type OutputFormatterStyleStack struct {
	styles       []*OutputFormatterStyle
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
func (stack *OutputFormatterStyleStack) Pop(style *OutputFormatterStyle) *OutputFormatterStyle {
	if 0 == len(stack.styles) {
		return stack.defaultStyle
	}

	if nil == style {
		first := stack.styles[len(stack.styles)-1]
		newStack := stack.styles[:len(stack.styles)-1]
		stack.styles = newStack
		return first
	}

	for index := len(stack.styles) - 1; index >= 0; index-- {
		stackedStyle := stack.styles[index]

		currentStyleResult := (*style).Apply("")
		stackedStyleResult := stackedStyle.Apply("")

		if currentStyleResult == stackedStyleResult {
			stack.styles = stack.styles[:index]
			return stackedStyle
		}
	}

	panic(errors.New("incorrectly nested style tag found"))
}

// Computes current style with stacks top codes
func (stack *OutputFormatterStyleStack) GetCurrent() *OutputFormatterStyle {
	if 0 == len(stack.styles) {
		return stack.defaultStyle
	}

	return stack.styles[len(stack.styles)-1]
}

// get default style
func (stack *OutputFormatterStyleStack) GetDefaultStyle() *OutputFormatterStyle {
	return stack.defaultStyle
}

// set default style
func (stack *OutputFormatterStyleStack) SetDefaultStyle(style OutputFormatterStyle) {
	stack.defaultStyle = &style
}
