package argument

import (
	"DrSmithFr/go-console/pkg/input/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	arg := argument.New("foo", argument.OPTIONAL)

	assert.Equal(t, "foo", arg.GetName())
}

func TestModes(t *testing.T) {
	arg1 := argument.New("foo", argument.OPTIONAL)
	assert.False(t, arg1.IsRequired())

	arg2 := argument.New("foo", argument.REQUIRED)
	assert.True(t, arg2.IsRequired())
}

func TestInvalidModes(t *testing.T) {
	assert.Panics(
		t,
		func() {
			argument.New("foo", -1)
		},
	)
}

func TestIsArray(t *testing.T) {
	arg1 := argument.New("foo", argument.IS_ARRAY)
	assert.True(t, arg1.IsArray())

	arg2 := argument.New("foo", argument.OPTIONAL|argument.IS_ARRAY)
	assert.True(t, arg2.IsArray())

	arg3 := argument.New("foo", argument.OPTIONAL)
	assert.False(t, arg3.IsArray())
}

func TestGetDescription(t *testing.T) {
	arg := argument.
		New("foo", argument.IS_ARRAY).
		SetDescription("Some description")

	assert.Equal(t, "Some description", arg.GetDescription())
}

func TestGetDefault(t *testing.T) {
	arg1 := argument.
		New("foo", argument.OPTIONAL).
		SetDefault("default")

	assert.Equal(t, "default", arg1.GetDefault())

	arg2 := argument.
		New("foo", argument.IS_ARRAY).
		SetDefaults([]string{"default"})

	assert.Equal(t, []string{"default"}, arg2.GetDefaults())
}

func TestSetDefaultsOnNotArray(t *testing.T) {
	assert.Panics(t, func() {
		argument.
			New("foo", argument.OPTIONAL).
			SetDefaults([]string{"default"})
	})
}

func TestSetDefaultOnArray(t *testing.T) {
	assert.Panics(t, func() {
		argument.
			New("foo", argument.IS_ARRAY).
			SetDefault("default")
	})

	assert.Panics(t, func() {
		argument.
			New("foo", argument.IS_ARRAY|argument.REQUIRED).
			SetDefault("default")
	})
}

func TestSetDefault(t *testing.T) {
	arg1 := argument.New("foo", argument.OPTIONAL)

	assert.Equal(t, "", arg1.GetDefault())

	arg1.SetDefault("another")

	assert.Equal(t, "another", arg1.GetDefault())

	arg2 := argument.
		New("foo", argument.IS_ARRAY).
		SetDefaults([]string{"1", "2"})

	assert.Equal(t, []string{"1", "2"}, arg2.GetDefaults())
}

func TestSetDefaultWithRequiredArgument(t *testing.T) {
	assert.Panics(t, func() {
		argument.
			New("foo", argument.REQUIRED).
			SetDefault("default")
	})
}
