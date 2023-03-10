package color

import (
	"github.com/DrSmithFr/go-console/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := color.NewColor(1, 2)

	assert.Equal(t, 1, c.Value())
	assert.Equal(t, 2, c.Unset())
}
