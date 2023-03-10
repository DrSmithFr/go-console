package argument

import (
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	arg := argument.New("foo", argument.Optional)

	assert.Equal(t, "foo", arg.Name())
}

func TestModes(t *testing.T) {
	arg1 := argument.New("foo", argument.Optional)
	assert.False(t, arg1.IsRequired())

	arg2 := argument.New("foo", argument.Required)
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

func TestIsList(t *testing.T) {
	arg1 := argument.New("foo", argument.List)
	assert.True(t, arg1.IsList())

	arg2 := argument.New("foo", argument.Optional|argument.List)
	assert.True(t, arg2.IsList())

	arg3 := argument.New("foo", argument.Optional)
	assert.False(t, arg3.IsList())
}

func TestGetDescription(t *testing.T) {
	arg := argument.New("foo", argument.List).
		SetDescription("Some description")

	assert.Equal(t, "Some description", arg.Description())
}

func TestGetDefault(t *testing.T) {
	arg1 := argument.New("foo", argument.Optional).
		SetDefault("default")

	assert.Equal(t, "default", arg1.Default())

	arg2 := argument.New("foo", argument.List).
		SetDefaults([]string{"default"})

	assert.Equal(t, []string{"default"}, arg2.Defaults())
}

func TestSetDefaultsOnNotList(t *testing.T) {
	assert.Panics(t, func() {
		argument.New("foo", argument.Optional).
			SetDefaults([]string{"default"})
	})
}

func TestSetDefaultOnList(t *testing.T) {
	assert.Panics(t, func() {
		argument.New("foo", argument.List).
			SetDefault("default")
	})

	assert.Panics(t, func() {
		argument.New("foo", argument.List|argument.Required).
			SetDefault("default")
	})
}

func TestSetDefault(t *testing.T) {
	arg1 := argument.New("foo", argument.Optional)

	assert.Equal(t, "", arg1.Default())

	arg1.SetDefault("another")

	assert.Equal(t, "another", arg1.Default())

	arg2 := argument.New("foo", argument.List).
		SetDefaults([]string{"1", "2"})

	assert.Equal(t, []string{"1", "2"}, arg2.Defaults())
}

func TestSetDefaultWithRequiredArgument(t *testing.T) {
	assert.Panics(t, func() {
		argument.New("foo", argument.Required).
			SetDefault("default")
	})
}
