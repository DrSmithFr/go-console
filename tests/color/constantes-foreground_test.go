package color

import (
	"github.com/DrSmithFr/go-console/pkg/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForeground(t *testing.T) {
	assert.Equal(t, color.NewColor(30, 39), color.GetForegroundColor(color.Black))
	assert.Equal(t, color.NewColor(31, 39), color.GetForegroundColor(color.Red))
	assert.Equal(t, color.NewColor(32, 39), color.GetForegroundColor(color.Green))
	assert.Equal(t, color.NewColor(33, 39), color.GetForegroundColor(color.Yellow))
	assert.Equal(t, color.NewColor(34, 39), color.GetForegroundColor(color.Blue))
	assert.Equal(t, color.NewColor(35, 39), color.GetForegroundColor(color.Magenta))
	assert.Equal(t, color.NewColor(36, 39), color.GetForegroundColor(color.Cyan))
	assert.Equal(t, color.NewColor(37, 39), color.GetForegroundColor(color.White))
	assert.Equal(t, color.NewColor(39, 39), color.GetForegroundColor(color.Default))

	assert.Panics(t, func() {
		color.GetForegroundColor("undefined-color")
	})
}
