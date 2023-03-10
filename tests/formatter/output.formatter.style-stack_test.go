package formatter

import (
	"github.com/DrSmithFr/go-console/color"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	stack := formatter.NewOutputFormatterStyleStack(nil)

	s1 := formatter.NewOutputFormatterStyle(color.White, color.Black, nil)
	s2 := formatter.NewOutputFormatterStyle(color.Yellow, color.Blue, nil)

	stack.Push(s1)
	stack.Push(s2)

	assert.EqualValues(t, s2, stack.GetCurrent())

	s3 := formatter.NewOutputFormatterStyle(color.Green, color.Red, nil)
	stack.Push(s3)

	assert.EqualValues(t, s3, stack.GetCurrent())
}

func TestPop(t *testing.T) {
	stack := formatter.NewOutputFormatterStyleStack(nil)

	s1 := formatter.NewOutputFormatterStyle(color.White, color.Black, nil)
	s2 := formatter.NewOutputFormatterStyle(color.Yellow, color.Blue, nil)

	stack.Push(s1)
	stack.Push(s2)

	assert.EqualValues(t, s2, stack.Pop(nil))
	assert.EqualValues(t, s1, stack.Pop(nil))
}

func TestPopEmpty(t *testing.T) {
	stack := formatter.NewOutputFormatterStyleStack(nil)
	s := formatter.NewOutputFormatterStyle(color.Null, color.Null, nil)

	assert.EqualValues(t, s, stack.Pop(nil))
}

func TestPopNotLast(t *testing.T) {
	stack := formatter.NewOutputFormatterStyleStack(nil)

	s1 := formatter.NewOutputFormatterStyle(color.White, color.Black, nil)
	s2 := formatter.NewOutputFormatterStyle(color.Yellow, color.Blue, nil)
	s3 := formatter.NewOutputFormatterStyle(color.Green, color.Red, nil)

	stack.Push(s1)
	stack.Push(s2)
	stack.Push(s3)

	assert.EqualValues(t, s2, stack.Pop(s2))
	assert.EqualValues(t, s1, stack.Pop(nil))
}

func TestInvalidPop(t *testing.T) {
	stack := formatter.NewOutputFormatterStyleStack(nil)

	s1 := formatter.NewOutputFormatterStyle(color.White, color.Black, nil)
	s2 := formatter.NewOutputFormatterStyle(color.Yellow, color.Blue, nil)

	stack.Push(s1)

	assert.Panics(t, func() { stack.Pop(s2) })
}
