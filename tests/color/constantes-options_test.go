package color

import (
	"github.com/DrSmithFr/go-console/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptions(t *testing.T) {
	assert.Equal(t, color.NewColor(1, 22), color.GetOption(color.Bold))
	assert.Equal(t, color.NewColor(4, 24), color.GetOption(color.Underscore))
	assert.Equal(t, color.NewColor(5, 25), color.GetOption(color.Blink))
	assert.Equal(t, color.NewColor(7, 27), color.GetOption(color.Reverse))
	assert.Equal(t, color.NewColor(8, 28), color.GetOption(color.Conceal))

	assert.Panics(t, func() {
		color.GetOption("undefined-option")
	})
}
