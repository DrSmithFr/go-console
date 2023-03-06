package table

import (
	"math"
	"strings"
)

type PaddingType int

const (
	STR_PAD_LEFT  PaddingType = iota
	STR_PAD_RIGHT PaddingType = iota
	STR_PAD_BOTH  PaddingType = iota
)

type TableStyleInterface interface {
	GetPaddingChar() string
	GetHorizontalBorderChar() string
	GetVerticalBorderChar() string
	GetCrossingChar() string
	GetCellHeaderFormat() string
	GetCellRowFormat() string
	GetCellRowContentFormat() string
	GetBorderFormat() string
	GetPadType() PaddingType
	Pad(content string, length int, pad string, direction PaddingType) string
}

type TableStyle struct {
	paddingChar          string
	horizontalBorderChar string
	verticalBorderChar   string
	crossingChar         string
	cellHeaderFormat     string
	cellRowFormat        string
	cellRowContentFormat string
	borderFormat         string
	padType              PaddingType
}

func NewTableStyle() *TableStyle {
	s := new(TableStyle)

	s.paddingChar = " "
	s.horizontalBorderChar = "-"
	s.verticalBorderChar = "|"
	s.crossingChar = "+"
	s.cellHeaderFormat = "<info>%s</info>"
	s.cellRowFormat = "%s"
	s.cellRowContentFormat = " %s "
	s.borderFormat = "%s"
	s.padType = STR_PAD_RIGHT

	return s
}

// custom methods
func (t TableStyle) Pad(content string, length int, pad string, direction PaddingType) string {
	contentLen := len(content)
	if contentLen >= length {
		return content
	}

	numPads := int(math.Ceil(float64(length-contentLen) / float64(len(pad))))
	switch direction {
	case STR_PAD_LEFT:
		return strings.Repeat(pad, numPads) + content
	case STR_PAD_RIGHT:
		return content + strings.Repeat(pad, numPads)
	case STR_PAD_BOTH:
		quotient := numPads / 2
		remainder := numPads % 2
		return strings.Repeat(pad, quotient) + content + strings.Repeat(pad, quotient+remainder)
	}

	panic("Invalid padding type")
}

// Implement TableStyleInterface

var _ TableStyleInterface = (*TableStyle)(nil)

func (t TableStyle) GetPaddingChar() string {
	return t.paddingChar
}

func (t TableStyle) GetHorizontalBorderChar() string {
	return t.horizontalBorderChar
}

func (t TableStyle) GetVerticalBorderChar() string {
	return t.verticalBorderChar
}

func (t TableStyle) GetCrossingChar() string {
	return t.crossingChar
}

func (t TableStyle) GetCellHeaderFormat() string {
	return t.cellHeaderFormat
}

func (t TableStyle) GetCellRowFormat() string {
	return t.cellRowFormat
}

func (t TableStyle) GetCellRowContentFormat() string {
	return t.cellRowContentFormat
}

func (t TableStyle) GetBorderFormat() string {
	return t.borderFormat
}

func (t TableStyle) GetPadType() PaddingType {
	return t.padType
}

// Implement TableStyle fluent setters

func (t *TableStyle) SetPaddingChar(paddingChar string) *TableStyle {
	t.paddingChar = paddingChar
	return t
}

func (t *TableStyle) SetHorizontalBorderChar(horizontalBorderChar string) *TableStyle {
	t.horizontalBorderChar = horizontalBorderChar
	return t
}

func (t *TableStyle) SetVerticalBorderChar(verticalBorderChar string) *TableStyle {
	t.verticalBorderChar = verticalBorderChar
	return t
}

func (t *TableStyle) SetCrossingChar(crossingChar string) *TableStyle {
	t.crossingChar = crossingChar
	return t
}

func (t *TableStyle) SetCellHeaderFormat(cellHeaderFormat string) *TableStyle {
	t.cellHeaderFormat = cellHeaderFormat
	return t
}

func (t *TableStyle) SetCellRowFormat(cellRowFormat string) *TableStyle {
	t.cellRowFormat = cellRowFormat
	return t
}

func (t *TableStyle) SetCellRowContentFormat(cellRowContentFormat string) *TableStyle {
	t.cellRowContentFormat = cellRowContentFormat
	return t
}

func (t *TableStyle) SetBorderFormat(borderFormat string) *TableStyle {
	t.borderFormat = borderFormat
	return t
}

func (t *TableStyle) SetPadType(padType PaddingType) *TableStyle {
	t.padType = padType
	return t
}
