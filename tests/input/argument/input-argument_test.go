package argument

import (
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	arg := argument.NewInputArgument("foo", argument.OPTIONAL, "", "")

	assert.Equal(t, "foo", arg.GetName())
}

func TestModes(t *testing.T) {
	arg1 := argument.NewInputArgument("foo", argument.OPTIONAL, "", "")
	assert.False(t, arg1.IsRequired())

	arg2 := argument.NewInputArgument("foo", argument.REQUIRED, "", "")
	assert.True(t, arg2.IsRequired())
}

func TestInvalidModes(t *testing.T) {
	assert.Panics(
		t,
		func() {
			argument.NewInputArgument("foo", -1, "", "")
		},
	)
}

func TestIsArray(t *testing.T) {
	arg1 := argument.NewInputArgument("foo", argument.IS_ARRAY, "", "")
	assert.True(t, arg1.IsArray())

	arg2 := argument.NewInputArgument("foo", argument.OPTIONAL|argument.IS_ARRAY, "", "")
	assert.True(t, arg2.IsArray())

	arg3 := argument.NewInputArgument("foo", argument.OPTIONAL, "", "")
	assert.False(t, arg3.IsArray())
}

func TestGetDescription(t *testing.T) {
	arg := argument.NewInputArgument("foo", argument.IS_ARRAY, "Some description", "")
	assert.Equal(t, "Some description", arg.GetDescription())
}

func TestGetDefault(t *testing.T) {
	arg := argument.NewInputArgument("foo", argument.IS_ARRAY, "", "default")
	assert.Equal(t, "default", arg.GetDefault())
}

func TestSetDefault(t *testing.T) {
	// TODO
}
