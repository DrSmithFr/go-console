package color

// color constructor
func NewColor(set int, unset int) Color {
	color := new(Color)

	color.set = set
	color.unset = unset

	return *color
}

// color struct
type Color struct {
	set   int
	unset int
}

// Get the setter color value
func (c *Color) Value() int {
	return c.set
}

// Get the default color value
func (c *Color) Unset() int {
	return c.unset
}
