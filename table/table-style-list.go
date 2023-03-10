package table

import (
	"fmt"
)

var Styles map[string]TableStyleInterface

func initStyles() {
	Styles = make(map[string]TableStyleInterface)

	borderless := NewTableStyle().
		SetHorizontalBorderChars("=").
		SetVerticalBorderChars(" ").
		SetCrossingChar(" ")

	compact := NewTableStyle().
		SetHorizontalBorderChars("").
		SetVerticalBorderChars(" ").
		SetCrossingChar("").
		SetCellRowContentFormat("%s")

	styleGuide := NewTableStyle().
		SetHorizontalBorderChars("-").
		SetVerticalBorderChars(" ").
		SetCrossingChar(" ").
		SetCellHeaderFormat("%s")

	box := NewTableStyle().
		SetHorizontalBorderChars("─").
		SetVerticalBorderChars("│").
		SetCrossingChars("┼", "┌", "┬", "┐", "┤", "┘", "┴", "└", "├", "├", "┼", "┤")

	boxDouble := NewTableStyle().
		SetHorizontalOutsideBorderChar("═").
		SetHorizontalInsideBorderChar("─").
		SetVerticalOutsideBorderChar("║").
		SetVerticalInsideBorderChar("│").
		SetCrossingChars("┼", "╔", "╤", "╗", "╢", "╝", "╧", "╚", "╟", "╠", "╪", "╣")

	Styles["default"] = NewTableStyle()
	Styles["borderless"] = borderless
	Styles["compact"] = compact
	Styles["style-guide"] = styleGuide
	Styles["box"] = box
	Styles["box-double"] = boxDouble
}

func SetStyleDefinition(name string, style TableStyleInterface) {
	if Styles == nil {
		initStyles()
	}

	Styles[name] = style
}

func GetStyleDefinition(name string) TableStyleInterface {
	if Styles == nil {
		initStyles()
	}

	if _, ok := Styles[name]; !ok {
		panic(fmt.Sprintf("Style '%s' is not defined.", name))
	}

	return Styles[name]
}
