package color

import (
	"github.com/DrSmithFr/go-console/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBackground(t *testing.T) {
	assert.Equal(t, color.NewColor(40, 49), color.BackgroundColor(color.Black))
	assert.Equal(t, color.NewColor(41, 49), color.BackgroundColor(color.Red))
	assert.Equal(t, color.NewColor(42, 49), color.BackgroundColor(color.Green))
	assert.Equal(t, color.NewColor(43, 49), color.BackgroundColor(color.Yellow))
	assert.Equal(t, color.NewColor(44, 49), color.BackgroundColor(color.Blue))
	assert.Equal(t, color.NewColor(45, 49), color.BackgroundColor(color.Magenta))
	assert.Equal(t, color.NewColor(46, 49), color.BackgroundColor(color.Cyan))
	assert.Equal(t, color.NewColor(47, 49), color.BackgroundColor(color.White))
	assert.Equal(t, color.NewColor(49, 49), color.BackgroundColor(color.Default))

	assert.Panics(t, func() {
		color.BackgroundColor("undefined-color")
	})
}
