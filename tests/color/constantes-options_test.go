package color

import (
	"github.com/MrSmith777/go-console/pkg/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptions(t *testing.T) {
	assert.Equal(t, color.NewColor(1, 22), color.GetOption(color.BOLD))
	assert.Equal(t, color.NewColor(4, 24), color.GetOption(color.UNDERSCORE))
	assert.Equal(t, color.NewColor(5, 25), color.GetOption(color.BLINK))
	assert.Equal(t, color.NewColor(7, 27), color.GetOption(color.REVERSE))
	assert.Equal(t, color.NewColor(8, 28), color.GetOption(color.CONCEAL))

	assert.Panics(t, func() {
		color.GetOption("undefined-option")
	})
}
