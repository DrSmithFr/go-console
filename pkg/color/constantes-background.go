package color

import "errors"

var backgroundColors = map[string]Color{
	Black:   NewColor(40, 49),
	Red:     NewColor(41, 49),
	Green:   NewColor(42, 49),
	Yellow:  NewColor(43, 49),
	Blue:    NewColor(44, 49),
	Magenta: NewColor(45, 49),
	Cyan:    NewColor(46, 49),
	White:   NewColor(47, 49),
	Default: NewColor(49, 49),
}

// get color from background const
func GetBackgroundColor(name string) Color {
	if color, ok := backgroundColors[name]; ok {
		return color
	}

	panic(errors.New("invalid background color specified"))
}
