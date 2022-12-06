package color

import (
	"DrSmithFr/go-console/pkg/color"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := color.NewColor(1, 2)

	assert.Equal(t, 1, c.GetValue())
	assert.Equal(t, 2, c.GetUnset())
}
