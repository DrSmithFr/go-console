package color

import (
	"github.com/DrSmithFr/go-console/pkg/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBackground(t *testing.T) {
	assert.Equal(t, color.NewColor(40, 49), color.GetBackgroundColor(color.BLACK))
	assert.Equal(t, color.NewColor(41, 49), color.GetBackgroundColor(color.RED))
	assert.Equal(t, color.NewColor(42, 49), color.GetBackgroundColor(color.GREEN))
	assert.Equal(t, color.NewColor(43, 49), color.GetBackgroundColor(color.YELLOW))
	assert.Equal(t, color.NewColor(44, 49), color.GetBackgroundColor(color.BLUE))
	assert.Equal(t, color.NewColor(45, 49), color.GetBackgroundColor(color.MAGENTA))
	assert.Equal(t, color.NewColor(46, 49), color.GetBackgroundColor(color.CYAN))
	assert.Equal(t, color.NewColor(47, 49), color.GetBackgroundColor(color.WHITE))
	assert.Equal(t, color.NewColor(49, 49), color.GetBackgroundColor(color.DEFAULT))

	assert.Panics(t, func() {
		color.GetBackgroundColor("undefined-color")
	})
}
