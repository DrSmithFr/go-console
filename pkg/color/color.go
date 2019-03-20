package color

func NewColor(set int, unset int) Color {
	color := new(Color)

	color.set = set
	color.unset = unset

	return *color
}

type Color struct {
	set   int
	unset int
}

func (c *Color) GetValue() int {
	return c.set
}

func (c *Color) GetUnset() int {
	return c.unset
}
