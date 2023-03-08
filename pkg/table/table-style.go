package table

import (
	"math"
	"strings"
)

type PaddingType int

const (
	PadDefault  PaddingType = 0
	PadToRight  PaddingType = 1
	PadToLeft   PaddingType = 2
	PadToCenter PaddingType = 3
)

type TableStyleInterface interface {
	GetPaddingChar() string

	GetHorizontalOutsideBorderChar() string
	GetHorizontalInsideBorderChar() string

	GetVerticalOutsideBorderChar() string
	GetVerticalInsideBorderChar() string

	GetCrossingChar() string

	GetCrossingTopRightChar() string
	GetCrossingTopMidChar() string
	GetCrossingTopLeftChar() string

	GetCrossingBottomRightChar() string
	GetCrossingBottomMidChar() string
	GetCrossingBottomLeftChar() string

	GetCrossingMidRightChar() string
	GetCrossingMidLeftChar() string

	GetCrossingTopLeftBottomChar() string
	GetCrossingTopMidBottomChar() string
	GetCrossingTopRightBottomChar() string

	GetHeaderTitleFormat() string
	GetFooterTitleFormat() string

	GetCellHeaderFormat() string
	GetCellRowFormat() string
	GetCellRowContentFormat() string
	GetBorderFormat() string
	GetPadType() PaddingType
	Pad(content string, length int, pad string, direction PaddingType) string
}

type TableStyle struct {
	paddingChar string

	horizontalOutsideBorderChar string
	horizontalInsideBorderChar  string

	verticalOutsideBorderChar string
	verticalInsideBorderChar  string

	crossingChar string

	crossingTopRightChar string
	crossingTopMidChar   string
	crossingTopLeftChar  string

	crossingBottomRightChar string
	crossingBottomMidChar   string
	crossingBottomLeftChar  string

	crossingMidRightChar string
	crossingMidLeftChar  string

	crossingTopLeftBottomChar  string
	crossingTopMidBottomChar   string
	crossingTopRightBottomChar string

	headerTitleFormat string
	footerTitleFormat string

	cellHeaderFormat     string
	cellRowFormat        string
	cellRowContentFormat string
	borderFormat         string

	padType PaddingType
}

func NewTableStyle() *TableStyle {
	s := new(TableStyle)

	s.paddingChar = " "

	s.horizontalOutsideBorderChar = "-"
	s.horizontalInsideBorderChar = "-"

	s.verticalOutsideBorderChar = "|"
	s.verticalInsideBorderChar = "|"

	s.crossingChar = "+"

	s.crossingTopRightChar = "+"
	s.crossingTopMidChar = "+"
	s.crossingTopLeftChar = "+"

	s.crossingMidRightChar = "+"
	s.crossingMidLeftChar = "+"

	s.crossingBottomRightChar = "+"
	s.crossingBottomMidChar = "+"
	s.crossingBottomLeftChar = "+"

	s.crossingTopLeftBottomChar = "+"
	s.crossingTopMidBottomChar = "+"
	s.crossingTopRightBottomChar = "+"

	s.headerTitleFormat = "<fg=black;bg=white;options=bold> %s </>"
	s.footerTitleFormat = "<fg=black;bg=white;options=bold> %s </>"

	s.cellHeaderFormat = "<info>%s</info>"
	s.cellRowFormat = "%s"
	s.cellRowContentFormat = " %s "
	s.borderFormat = "%s"
	s.padType = PadToLeft

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
	case PadToRight:
		return strings.Repeat(pad, numPads) + content
	case PadToLeft:
		return content + strings.Repeat(pad, numPads)
	case PadToCenter:
		quotient := numPads / 2
		remainder := numPads % 2
		return strings.Repeat(pad, quotient) + content + strings.Repeat(pad, quotient+remainder)
	}

	panic("Invalid padding type")
}

// Implement TableStyleInterface

var _ TableStyleInterface = (*TableStyle)(nil)

func (t *TableStyle) GetPaddingChar() string {
	return t.paddingChar
}

func (t *TableStyle) GetHorizontalOutsideBorderChar() string {
	return t.horizontalOutsideBorderChar
}

func (t *TableStyle) GetHorizontalInsideBorderChar() string {
	return t.horizontalInsideBorderChar
}

func (t *TableStyle) GetVerticalOutsideBorderChar() string {
	return t.verticalOutsideBorderChar
}

func (t *TableStyle) GetVerticalInsideBorderChar() string {
	return t.verticalInsideBorderChar
}

func (t *TableStyle) GetCrossingChar() string {
	return t.crossingChar
}

func (t *TableStyle) GetCrossingTopRightChar() string {
	return t.crossingTopRightChar
}

func (t *TableStyle) GetCrossingTopMidChar() string {
	return t.crossingTopMidChar
}

func (t *TableStyle) GetCrossingTopLeftChar() string {
	return t.crossingTopLeftChar
}

func (t *TableStyle) GetCrossingBottomRightChar() string {
	return t.crossingBottomRightChar
}

func (t *TableStyle) GetCrossingBottomMidChar() string {
	return t.crossingBottomMidChar
}

func (t *TableStyle) GetCrossingBottomLeftChar() string {
	return t.crossingBottomLeftChar
}

func (t *TableStyle) GetCrossingMidRightChar() string {
	return t.crossingMidRightChar
}

func (t *TableStyle) GetCrossingMidLeftChar() string {
	return t.crossingMidLeftChar
}

func (t *TableStyle) GetCrossingTopLeftBottomChar() string {
	return t.crossingTopLeftBottomChar
}

func (t *TableStyle) GetCrossingTopMidBottomChar() string {
	return t.crossingTopMidBottomChar
}

func (t *TableStyle) GetCrossingTopRightBottomChar() string {
	return t.crossingTopRightBottomChar
}

func (t TableStyle) GetHeaderTitleFormat() string {
	return t.headerTitleFormat
}

func (t TableStyle) GetFooterTitleFormat() string {
	return t.footerTitleFormat
}

func (t *TableStyle) GetCellHeaderFormat() string {
	return t.cellHeaderFormat
}

func (t *TableStyle) GetCellRowFormat() string {
	return t.cellRowFormat
}

func (t *TableStyle) GetCellRowContentFormat() string {
	return t.cellRowContentFormat
}

func (t *TableStyle) GetBorderFormat() string {
	return t.borderFormat
}

func (t *TableStyle) GetPadType() PaddingType {
	return t.padType
}

// Implement TableStyle fluent setters

func (t *TableStyle) SetPaddingChar(paddingChar string) *TableStyle {
	t.paddingChar = paddingChar
	return t
}

func (t *TableStyle) SetHorizontalBorderChar(borderChar string) *TableStyle {
	t.horizontalOutsideBorderChar = borderChar
	t.horizontalInsideBorderChar = borderChar
	return t
}

func (t *TableStyle) SetHorizontalOutsideBorderChar(borderChar string) *TableStyle {
	t.horizontalOutsideBorderChar = borderChar
	return t
}

func (t *TableStyle) SetHorizontalInsideBorderChar(borderChar string) *TableStyle {
	t.horizontalInsideBorderChar = borderChar
	return t
}

func (t *TableStyle) SetVerticalBorderChar(borderChar string) *TableStyle {
	t.verticalOutsideBorderChar = borderChar
	t.verticalInsideBorderChar = borderChar
	return t
}

func (t *TableStyle) SetVerticalOutsideBorderChar(borderChar string) *TableStyle {
	t.verticalOutsideBorderChar = borderChar
	return t
}

func (t *TableStyle) SetVerticalInsideBorderChar(borderChar string) *TableStyle {
	t.verticalInsideBorderChar = borderChar
	return t
}

func (t *TableStyle) SetCrossingChar(crossingChar string) *TableStyle {
	t.crossingChar = crossingChar

	t.crossingBottomLeftChar = crossingChar
	t.crossingBottomMidChar = crossingChar
	t.crossingBottomRightChar = crossingChar

	t.crossingMidLeftChar = crossingChar
	t.crossingMidRightChar = crossingChar

	t.crossingTopLeftChar = crossingChar
	t.crossingTopMidChar = crossingChar
	t.crossingTopRightChar = crossingChar

	t.crossingTopLeftBottomChar = crossingChar
	t.crossingTopMidBottomChar = crossingChar
	t.crossingTopRightBottomChar = crossingChar
	return t
}

func (t *TableStyle) SetCrossingChars(
	cross string,
	topLeft string,
	topMid string,
	topRight string,
	midRight string,
	bottomRight string,
	bottomMid string,
	bottomLeft string,
	midLeft string,
	topLeftBottom string,
	topMidBottom string,
	topRightBottom string,

) *TableStyle {
	t.crossingChar = cross

	t.crossingBottomLeftChar = bottomLeft
	t.crossingBottomMidChar = bottomMid
	t.crossingBottomRightChar = bottomRight

	t.crossingMidLeftChar = midLeft
	t.crossingMidRightChar = midRight

	t.crossingTopLeftChar = topLeft
	t.crossingTopMidChar = topMid
	t.crossingTopRightChar = topRight

	t.crossingTopLeftBottomChar = topLeftBottom
	t.crossingTopMidBottomChar = topMidBottom
	t.crossingTopRightBottomChar = topRightBottom

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
