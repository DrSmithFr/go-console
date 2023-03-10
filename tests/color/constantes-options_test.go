package color

import (
	"github.com/DrSmithFr/go-console/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptions(t *testing.T) {
	assert.Equal(t, color.NewColor(1, 22), color.Option(color.Bold))
	assert.Equal(t, color.NewColor(4, 24), color.Option(color.Underscore))
	assert.Equal(t, color.NewColor(5, 25), color.Option(color.Blink))
	assert.Equal(t, color.NewColor(7, 27), color.Option(color.Reverse))
	assert.Equal(t, color.NewColor(8, 28), color.Option(color.Conceal))

	assert.Panics(t, func() {
		color.Option("undefined-option")
	})
}
