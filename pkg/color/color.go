package color

func NewColor(set uint, unset uint) Color  {
	color := new(Color)

	color.set = set
	color.unset = unset

	return *color
}

type Color struct {
	set uint
	unset uint
}

func (c *Color) GetValue() uint {
	return c.set
}

func (c *Color) GetUnset() uint {
	return c.unset
}