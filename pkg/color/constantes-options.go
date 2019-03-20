package color

import "errors"

var options = map[string]Color{
	BOLD:       NewColor(1, 22),
	UNDERSCORE: NewColor(4, 24),
	BLINK:      NewColor(5, 25),
	REVERSE:    NewColor(7, 27),
	CONCEAL:    NewColor(8, 28),
}

func GetOption(name string) Color {
	if option, ok := options[name]; ok {
		return option
	}

	panic(errors.New("invalid foreground color specified"))
}
