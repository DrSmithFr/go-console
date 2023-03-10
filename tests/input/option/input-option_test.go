package option

import (
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstructor(t *testing.T) {
	opt1 := option.New("foo", option.Optional)
	assert.Equal(t, "foo", opt1.GetName())

	opt2 := option.New("--foo", option.Optional)
	assert.Equal(t, "foo", opt2.GetName())
}

func TestArrayModeWithoutValue(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.List).
			SetShortcut("f")
	})
}

func TestShortcut(t *testing.T) {
	opt1 := option.
		New("foo", option.Optional).
		SetShortcut("f")

	assert.Equal(t, "f", opt1.GetShortcut())

	opt2 := option.
		New("foo", option.Optional).
		SetShortcut("-f|-ff|fff")

	assert.Equal(t, "f|ff|fff", opt2.GetShortcut())

	opt3 := option.New("foo", option.Optional)
	assert.Equal(t, "", opt3.GetShortcut())
}

func TestModes(t *testing.T) {
	opt2 := option.New("foo", option.None)

	assert.False(t, opt2.AcceptValue())
	assert.False(t, opt2.IsValueRequired())
	assert.False(t, opt2.IsValueOptional())

	opt3 := option.New("foo", option.Required)

	assert.True(t, opt3.AcceptValue())
	assert.True(t, opt3.IsValueRequired())
	assert.False(t, opt3.IsValueOptional())

	opt4 := option.New("foo", option.Optional)

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
		option.New("", option.None)
	})
}

func TestDoubleDashNameIsInvalid(t *testing.T) {
	assert.Panics(t, func() {
		option.New("--", option.None)
	})
}

func TestSingleDashOptionShortcutIsInvalid(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.None).
			SetShortcut("-")
	})
}

func TestIsList(t *testing.T) {
	opt1 := option.New("foo", option.Optional|option.List)
	assert.True(t, opt1.IsList())

	opt2 := option.New("foo", option.None)
	assert.False(t, opt2.IsList())
}

func TestGetDescription(t *testing.T) {
	opt1 := option.
		New("foo", option.Optional|option.List).
		SetDescription("Some description")
	assert.Equal(t, "Some description", opt1.GetDescription())

}

func TestGetDefault(t *testing.T) {
	opt1 := option.
		New("foo", option.Optional).
		SetDefault("default")

	assert.Equal(t, "default", opt1.GetDefault())

	opt2 := option.
		New("foo", option.Required).
		SetDefault("default")

	assert.Equal(t, "default", opt2.GetDefault())

	opt3 := option.
		New("foo", option.Required)
	assert.Equal(t, "", opt3.GetDefault())

	opt4 := option.
		New("foo", option.None)
	assert.Equal(t, "", opt4.GetDefault())
}

func TestSetDefaultsOnNoneOption(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.None).
			SetDefault("default")
	})

	assert.Panics(t, func() {
		option.
			New("foo", option.None).
			SetDefaults([]string{"default"})
	})
}

func TestSetDefaultsOnNotList(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.Optional).
			SetDefaults([]string{"default"})
	})
}

func TestSetDefaultOnList(t *testing.T) {
	assert.Panics(t, func() {
		option.
			New("foo", option.List).
			SetDefault("default")
	})

	assert.Panics(t, func() {
		option.
			New("foo", option.List|option.Required).
			SetDefault("default")
	})
}

func TestEquals(t *testing.T) {
	opt1 := option.
		New("foo", option.None).
		SetShortcut("f").
		SetDescription("Some description")

	opt2 := option.
		New("foo", option.None).
		SetShortcut("f").
		SetDescription("Alternative description")

	assert.True(t, opt1.Equals(*opt2))

	opt3 := option.
		New("foo", option.Optional).
		SetShortcut("f")

	opt4 := option.
		New("foo", option.Optional).
		SetShortcut("f").
		SetDefault("default")

	assert.False(t, opt3.Equals(*opt4))

	opt5 := option.
		New("foo", option.Optional).
		SetShortcut("f").
		SetDescription("Some description")

	opt6 := option.
		New("bar", option.Optional).
		SetShortcut("f").
		SetDescription("Some description")

	assert.False(t, opt5.Equals(*opt6))

	opt7 := option.
		New("foo", option.None).
		SetShortcut("f")

	opt8 := option.
		New("foo", option.Optional).
		SetShortcut("f")

	assert.False(t, opt7.Equals(*opt8))
}
