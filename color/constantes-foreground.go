package color

import "errors"

var foregroundColors = map[string]Color{
	Black:   NewColor(30, 39),
	Red:     NewColor(31, 39),
	Green:   NewColor(32, 39),
	Yellow:  NewColor(33, 39),
	Blue:    NewColor(34, 39),
	Magenta: NewColor(35, 39),
	Cyan:    NewColor(36, 39),
	White:   NewColor(37, 39),
	Default: NewColor(39, 39),
}

// get color from foreground const
func ForegroundColor(name string) Color {
	if color, ok := foregroundColors[name]; ok {
		return color
	}

	panic(errors.New("invalid foreground color specified"))
}
