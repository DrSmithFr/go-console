package color

import (
	"github.com/MrSmith777/go-console/pkg/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForeground(t *testing.T) {
	assert.Equal(t, color.NewColor(30, 39), color.GetForegroundColor(color.BLACK))
	assert.Equal(t, color.NewColor(31, 39), color.GetForegroundColor(color.RED))
	assert.Equal(t, color.NewColor(32, 39), color.GetForegroundColor(color.GREEN))
	assert.Equal(t, color.NewColor(33, 39), color.GetForegroundColor(color.YELLOW))
	assert.Equal(t, color.NewColor(34, 39), color.GetForegroundColor(color.BLUE))
	assert.Equal(t, color.NewColor(35, 39), color.GetForegroundColor(color.MAGENTA))
	assert.Equal(t, color.NewColor(36, 39), color.GetForegroundColor(color.CYAN))
	assert.Equal(t, color.NewColor(37, 39), color.GetForegroundColor(color.WHITE))
	assert.Equal(t, color.NewColor(39, 39), color.GetForegroundColor(color.DEFAULT))

	assert.Panics(t, func() {
		color.GetForegroundColor("undefined-color")
	})
}
