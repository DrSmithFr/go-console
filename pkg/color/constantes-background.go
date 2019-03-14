package color

import "errors"

var backgroundColors = map[string]Color{
	BLACK: NewColor(40, 49),
	RED: NewColor(41, 49),
	GREEN: NewColor(42, 49),
	YELLOW: NewColor(43, 49),
	BLUE: NewColor(44, 49),
	MAGENTA: NewColor(45, 49),
	CYAN: NewColor(46, 49),
	WHITE: NewColor(47, 49),
	DEFAULT: NewColor(49, 49),
}

func GetBackgroundColor(name string) Color {
	if color, ok := backgroundColors[name]; ok {
		return color
	}

	panic(errors.New("invalid background color specified"))
}
