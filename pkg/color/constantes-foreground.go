package color

import "errors"

var foregroundColors = map[string]Color{
	BLACK:   NewColor(30, 39),
	RED:     NewColor(31, 39),
	GREEN:   NewColor(32, 39),
	YELLOW:  NewColor(33, 39),
	BLUE:    NewColor(34, 39),
	MAGENTA: NewColor(35, 39),
	CYAN:    NewColor(36, 39),
	WHITE:   NewColor(37, 39),
	DEFAULT: NewColor(39, 39),
}

func GetForegroundColor(name string) Color {
	if color, ok := foregroundColors[name]; ok {
		return color
	}

	panic(errors.New("invalid foreground color specified"))
}
