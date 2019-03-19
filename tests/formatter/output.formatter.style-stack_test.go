package formatter

import (
	"github.com/MrSmith777/go-console/pkg/color"
	"github.com/MrSmith777/go-console/pkg/formatter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	stack := formatter.NewOutputFormatterStyleStack(nil)

	s1 := formatter.NewOutputFormatterStyle(color.WHITE, color.BLACK, nil)
	s2 := formatter.NewOutputFormatterStyle(color.YELLOW, color.BLUE, nil)

	stack.Push(s1)
	stack.Push(s2)

	assert.ObjectsAreEqual(s2, stack.GetCurrent())

	s3 := formatter.NewOutputFormatterStyle(color.GREEN, color.RED, nil)

	assert.ObjectsAreEqual(s3, stack.GetCurrent())
}

func TestPop(t *testing.T) {
	stack := formatter.NewOutputFormatterStyleStack(nil)

	s1 := formatter.NewOutputFormatterStyle(color.WHITE, color.BLACK, nil)
	s2 := formatter.NewOutputFormatterStyle(color.YELLOW, color.BLUE, nil)

	stack.Push(s1)
	stack.Push(s2)

	assert.ObjectsAreEqual(s2, stack.Pop(nil))
	assert.ObjectsAreEqual(s1, stack.Pop(nil))
}

func TestPopEmpty(t *testing.T) {
	s := formatter.NewOutputFormatterStyle(color.NULL, color.NULL, nil)
	stack := formatter.NewOutputFormatterStyleStack(s)

	assert.ObjectsAreEqual(s, stack.Pop(nil))
}

func TestPopNotLast(t *testing.T) {
	stack := formatter.NewOutputFormatterStyleStack(nil)

	s1 := formatter.NewOutputFormatterStyle(color.WHITE, color.BLACK, nil)
	s2 := formatter.NewOutputFormatterStyle(color.YELLOW, color.BLUE, nil)
	s3 := formatter.NewOutputFormatterStyle(color.GREEN, color.RED, nil)

	stack.Push(s1)
	stack.Push(s2)
	stack.Push(s3)

	assert.ObjectsAreEqual(s2, stack.Pop(s2))
	assert.ObjectsAreEqual(s1, stack.Pop(nil))
}

func TestInvalidPop(t *testing.T) {
	stack := formatter.NewOutputFormatterStyleStack(nil)

	s1 := formatter.NewOutputFormatterStyle(color.WHITE, color.BLACK, nil)
	s2 := formatter.NewOutputFormatterStyle(color.YELLOW, color.BLUE, nil)

	stack.Push(s1)

	assert.Panics(t, func() { stack.Pop(s2) })
}