package option

import (
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	opt1 := option.New("foo", option.OPTIONAL)
	assert.Equal(t, "foo", opt1.GetName())

	opt2 := option.New("--foo", option.OPTIONAL)
	assert.Equal(t, "foo", opt2.GetName())
}

func TestArrayModeWithoutValue(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.IS_ARRAY).
			SetShortcut("f")
	})
}

func TestShortcut(t *testing.T) {
	opt1 := option.
		New("foo", option.OPTIONAL).
		SetShortcut("f")

	assert.Equal(t, "f", opt1.GetShortcut())

	opt2 := option.
		New("foo", option.OPTIONAL).
		SetShortcut("-f|-ff|fff")

	assert.Equal(t, "f|ff|fff", opt2.GetShortcut())

	opt3 := option.New("foo", option.OPTIONAL)
	assert.Equal(t, "", opt3.GetShortcut())
}

func TestModes(t *testing.T) {
	opt2 := option.New("foo", option.NONE)

	assert.False(t, opt2.AcceptValue())
	assert.False(t, opt2.IsValueRequired())
	assert.False(t, opt2.IsValueOptional())

	opt3 := option.New("foo", option.REQUIRED)

	assert.True(t, opt3.AcceptValue())
	assert.True(t, opt3.IsValueRequired())
	assert.False(t, opt3.IsValueOptional())

	opt4 := option.New("foo", option.OPTIONAL)

	assert.True(t, opt4.AcceptValue())
	assert.False(t, opt4.IsValueRequired())
	assert.True(t, opt4.IsValueOptional())
}

func TestInvalidModes(t *testing.T) {
	assert.Panics(t, func() {
		option.New("foo", -1)
	})
}

func TestEmptyNameIsInvalid(t *testing.T) {
	assert.Panics(t, func() {
		option.New("", option.NONE)
	})
}

func TestDoubleDashNameIsInvalid(t *testing.T) {
	assert.Panics(t, func() {
		option.New("--", option.NONE)
	})
}

func TestSingleDashOptionShortcutIsInvalid(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.NONE).
			SetShortcut("-")
	})
}

func TestIsArray(t *testing.T) {
	opt1 := option.New("foo", option.OPTIONAL|option.IS_ARRAY)
	assert.True(t, opt1.IsArray())

	opt2 := option.New("foo", option.NONE)
	assert.False(t, opt2.IsArray())
}

func TestGetDescription(t *testing.T) {
	opt1 := option.
		New("foo", option.OPTIONAL|option.IS_ARRAY).
		SetDescription("Some description")
	assert.Equal(t, "Some description", opt1.GetDescription())

}

func TestGetDefault(t *testing.T) {
	opt1 := option.
		New("foo", option.OPTIONAL).
		SetDefault("default")

	assert.Equal(t, "default", opt1.GetDefault())

	opt2 := option.
		New("foo", option.REQUIRED).
		SetDefault("default")

	assert.Equal(t, "default", opt2.GetDefault())

	opt3 := option.
		New("foo", option.REQUIRED)
	assert.Equal(t, "", opt3.GetDefault())

	opt4 := option.
		New("foo", option.NONE)
	assert.Equal(t, "", opt4.GetDefault())
}

func TestSetDefaultsOnNoneOption(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.NONE).
			SetDefault("default")
	})

	assert.Panics(t, func() {
		option.
			New("foo", option.NONE).
			SetDefaults([]string{"default"})
	})
}

func TestSetDefaultsOnNotArray(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.OPTIONAL).
			SetDefaults([]string{"default"})
	})
}

func TestSetDefaultOnArray(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.IS_ARRAY).
			SetDefault("default")
	})

	assert.Panics(t, func() {
		option.
			New("foo", option.IS_ARRAY|option.REQUIRED).
			SetDefault("default")
	})
}

func TestEquals(t *testing.T) {
	opt1 := option.
		New("foo", option.NONE).
		SetShortcut("f").
		SetDescription("Some description")

	opt2 := option.
		New("foo", option.NONE).
		SetShortcut("f").
		SetDescription("Alternative description")

	assert.True(t, opt1.Equals(*opt2))

	opt3 := option.
		New("foo", option.OPTIONAL).
		SetShortcut("f")

	opt4 := option.
		New("foo", option.OPTIONAL).
		SetShortcut("f").
		SetDefault("default")

	assert.False(t, opt3.Equals(*opt4))

	opt5 := option.
		New("foo", option.OPTIONAL).
		SetShortcut("f").
		SetDescription("Some description")

	opt6 := option.
		New("bar", option.OPTIONAL).
		SetShortcut("f").
		SetDescription("Some description")

	assert.False(t, opt5.Equals(*opt6))

	opt7 := option.
		New("foo", option.NONE).
		SetShortcut("f")

	opt8 := option.
		New("foo", option.OPTIONAL).
		SetShortcut("f")

	assert.False(t, opt7.Equals(*opt8))
}
