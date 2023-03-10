package color

import "errors"

var options = map[string]Color{
	Bold:       NewColor(1, 22),
	Underscore: NewColor(4, 24),
	Blink:      NewColor(5, 25),
	Reverse:    NewColor(7, 27),
	Conceal:    NewColor(8, 28),
}

// get color from option const
func Option(name string) Color {
	if option, ok := options[name]; ok {
		return option
	}

	panic(errors.New("invalid foreground color specified"))
}
