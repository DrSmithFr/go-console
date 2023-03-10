package definition

import (
	"github.com/DrSmithFr/go-console/input/argument"
	"github.com/DrSmithFr/go-console/input/definition"
	"github.com/DrSmithFr/go-console/input/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

var arguments = map[string]argument.InputArgument{
	"foo":  *argument.New("foo", argument.Optional),
	"bar":  *argument.New("bar", argument.Optional),
	"foo1": *argument.New("foo", argument.Optional),
	"foo2": *argument.New("foo2", argument.Required),
}

var options = map[string]option.InputOption{
	"foo": *option.New("foo", option.Optional).
		SetShortcut("f"),
	"bar": *option.New("bar", option.Optional).
		SetShortcut("b"),
	"foo1": *option.New("fooBis", option.Optional).
		SetShortcut("f"),
	"foo2": *option.New("foo", option.Optional).
		SetShortcut("p"),
	"multi": *option.New("multi", option.Optional).
		SetShortcut("m|mm|mmm"),
}

func TestConstructorArguments(t *testing.T) {
	def1 := definition.New()
	assert.Equal(t, map[string]argument.InputArgument{}, def1.Arguments())
	assert.Equal(t, map[string]option.InputOption{}, def1.Options())
}

func TestSetArguments(t *testing.T) {
	def := definition.New().
		SetArguments([]argument.InputArgument{
			arguments["foo"],
		})

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"foo": arguments["foo"],
		},
		def.Arguments(),
	)

	def.SetArguments([]argument.InputArgument{
		arguments["bar"],
	})

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"bar": arguments["bar"],
		},
		def.Arguments(),
	)
}

func TestAddArguments(t *testing.T) {
	def := definition.New().
		AddArguments([]argument.InputArgument{
			arguments["foo"],
		})

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"foo": arguments["foo"],
		},
		def.Arguments(),
	)

	def.AddArguments([]argument.InputArgument{
		arguments["bar"],
	})

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"foo": arguments["foo"],
			"bar": arguments["bar"],
		},
		def.Arguments(),
	)
}

func TestAddArgument(t *testing.T) {
	def := definition.New().
		AddArgument(arguments["foo"])

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"foo": arguments["foo"],
		},
		def.Arguments(),
	)

	def.AddArgument(arguments["bar"])

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"foo": arguments["foo"],
			"bar": arguments["bar"],
		},
		def.Arguments(),
	)
}

func TestArgumentsMustHaveDifferentNames(t *testing.T) {
	assert.Panics(t, func() {
		definition.New().
			AddArgument(arguments["foo"]).
			AddArgument(arguments["foo1"])
	})
}

func TestArrayArgumentHasToBeLast(t *testing.T) {
	assert.Panics(t, func() {
		definition.New().
			AddArgument(*argument.New("fooarray", argument.List)).
			AddArgument(*argument.New("anotherbar", argument.Optional))
	})
}

func TestRequiredArgumentCannotFollowAnOptionalOne(t *testing.T) {
	assert.Panics(t, func() {
		definition.New().
			AddArgument(arguments["foo"]).
			AddArgument(arguments["foo2"])
	})
}

func TestGetArgument(t *testing.T) {
	def := definition.New().
		AddArgument(arguments["foo"])

	assert.Equal(t, arguments["foo"], *def.Argument("foo"))
}

func TestGetInvalidArgument(t *testing.T) {
	assert.Panics(t, func() {
		definition.New().
			AddArgument(arguments["foo"]).
			Argument("bar")
	})
}

func TestHasArgument(t *testing.T) {
	def := definition.New().
		AddArgument(arguments["foo"])

	assert.True(t, def.HasArgument("foo"))
	assert.False(t, def.HasArgument("bar"))
}

func TestGetArgumentRequiredCount(t *testing.T) {
	def := definition.New().
		AddArgument(arguments["foo2"])

	assert.Equal(t, 1, def.ArgumentRequiredCount())

	def.AddArgument(arguments["foo"])

	assert.Equal(t, 1, def.ArgumentRequiredCount())
}

func TestGetArgumentCount(t *testing.T) {
	def := definition.New().
		AddArgument(arguments["foo2"])

	assert.Equal(t, 1, def.ArgumentCount())

	def.AddArgument(arguments["foo"])

	assert.Equal(t, 2, def.ArgumentCount())
}

func TestGetArgumentDefaults(t *testing.T) {
	def := definition.New().
		SetArguments([]argument.InputArgument{
			*argument.New("foo1", argument.Optional),

			*argument.New("foo2", argument.Optional).
				SetDefault("default"),

			*argument.New("foo3", argument.Optional|argument.List),
		})

	validation := map[string][]string{
		"foo1": nil,
		"foo2": {"default"},
		"foo3": {},
		"foo4": {"1", "2"},
	}

	assert.Equal(t, validation["foo1"], def.ArgumentDefaults()["foo1"])
	assert.Equal(t, validation["foo2"], def.ArgumentDefaults()["foo2"])
	assert.Equal(t, validation["foo3"], def.ArgumentDefaults()["foo3"])

	def2 := definition.New().
		SetArguments([]argument.InputArgument{
			*argument.New("foo4", argument.Optional|argument.List).
				SetDefaults([]string{"1", "2"}),
		})

	assert.Equal(t, validation["foo4"], def2.ArgumentDefaults()["foo4"])
}

func TestSetOptions(t *testing.T) {
	def := definition.New().
		SetOptions([]option.InputOption{
			options["foo"],
		})

	assert.Equal(
		t,
		map[string]option.InputOption{
			"foo": options["foo"],
		},
		def.Options(),
	)

	def.SetOptions([]option.InputOption{
		options["bar"],
	})

	assert.Equal(
		t,
		map[string]option.InputOption{
			"bar": options["bar"],
		},
		def.Options(),
	)
}

func TestSetOptionsClearsOptions(t *testing.T) {
	assert.Panics(t, func() {
		definition.New().
			SetOptions([]option.InputOption{
				options["bar"],
			}).
			FindOptionForShortcut("f")
	})
}

func TestAddOptions(t *testing.T) {
	def := definition.New().
		AddOptions([]option.InputOption{
			options["foo"],
		})

	assert.Equal(
		t,
		map[string]option.InputOption{
			"foo": options["foo"],
		},
		def.Options(),
	)

	def.AddOptions([]option.InputOption{
		options["bar"],
	})

	assert.Equal(
		t,
		map[string]option.InputOption{
			"foo": options["foo"],
			"bar": options["bar"],
		},
		def.Options(),
	)
}

func TestAddOption(t *testing.T) {
	def := definition.New().
		AddOption(options["foo"])

	assert.Equal(
		t,
		map[string]option.InputOption{
			"foo": options["foo"],
		},
		def.Options(),
	)

	def.AddOption(options["bar"])

	assert.Equal(
		t,
		map[string]option.InputOption{
			"foo": options["foo"],
			"bar": options["bar"],
		},
		def.Options(),
	)
}

func TestAddDuplicateOption(t *testing.T) {
	assert.Panics(t, func() {
		definition.New().
			AddOption(options["foo"]).
			AddOption(options["foo2"])
	})
}

func TestAddDuplicateShortcutOption(t *testing.T) {
	assert.Panics(t, func() {
		definition.New().
			AddOption(options["foo"]).
			AddOption(options["foo1"])
	})
}

func TestGetOption(t *testing.T) {
	def := definition.New().
		AddOption(options["foo"])

	assert.Equal(t, options["foo"], *def.Option("foo"))
}

func TestGetInvalidOption(t *testing.T) {
	assert.Panics(t, func() {
		definition.New().
			AddOption(options["foo"]).
			Option("bar")
	})
}

func TestHasOption(t *testing.T) {
	def := definition.New().
		AddOption(options["foo"])

	assert.True(t, def.HasOption("foo"))
	assert.False(t, def.HasOption("BAR"))
}

func TestHasShortcut(t *testing.T) {
	def := definition.New().
		AddOption(options["foo"])

	assert.True(t, def.HasShortcut("f"))
	assert.False(t, def.HasShortcut("b"))
}

func TestGetOptionForShortcut(t *testing.T) {
	def := definition.New().
		AddOption(options["foo"])

	assert.Equal(t, options["foo"], *def.FindOptionForShortcut("f"))
}

func TestGetOptionForMultiShortcut(t *testing.T) {
	def := definition.New().
		AddOption(options["multi"])

	assert.Equal(t, options["multi"], *def.FindOptionForShortcut("m"))
	assert.Equal(t, options["multi"], *def.FindOptionForShortcut("mm"))
	assert.Equal(t, options["multi"], *def.FindOptionForShortcut("mmm"))
}

func TestGetOptionForInvalidShortcut(t *testing.T) {
	assert.Panics(t, func() {
		definition.New().
			FindOptionForShortcut("l")
	})
}

func TestGetOptionDefaults(t *testing.T) {
	def := definition.New().
		SetOptions([]option.InputOption{
			*option.New("foo1", option.None),

			*option.New("foo2", option.Required),

			*option.New("foo3", option.Required).
				SetDefault("default"),

			*option.New("foo4", option.Optional),

			*option.New("foo5", option.Optional).
				SetDefault("default"),

			*option.New("foo6", option.Optional|option.List),

			*option.New("foo7", option.Optional|option.List).
				SetDefaults([]string{"1", "2"}),
		})

	validation := map[string][]string{
		"foo1": {},
		"foo2": {},
		"foo3": {"default"},
		"foo4": {},
		"foo5": {"default"},
		"foo6": {},
		"foo7": {"1", "2"},
	}

	assert.Equal(t, validation["foo1"], def.OptionDefaults()["foo1"])
	assert.Equal(t, validation["foo2"], def.OptionDefaults()["foo2"])
	assert.Equal(t, validation["foo3"], def.OptionDefaults()["foo3"])
	assert.Equal(t, validation["foo4"], def.OptionDefaults()["foo4"])
	assert.Equal(t, validation["foo5"], def.OptionDefaults()["foo5"])
	assert.Equal(t, validation["foo6"], def.OptionDefaults()["foo6"])
	assert.Equal(t, validation["foo7"], def.OptionDefaults()["foo7"])
}

func TestGetSynopsis(t *testing.T) {
	for _, pattern := range getSynopticPattern() {
		assert.Equalf(t, pattern.synoptic, pattern.definition.Synopsis(false), pattern.message)
	}
}

type synopticPattern struct {
	definition definition.InputDefinition
	synoptic   string
	message    string
}

func getSynopticPattern() []synopticPattern {
	return []synopticPattern{
		// testing options
		{
			definition: *definition.New().
				AddOption(*option.New("foo", option.None)),
			synoptic: "[--foo]",
			message:  "puts optional options in square brackets",
		},
		{
			definition: *definition.New().
				AddOption(
					*option.New("foo", option.None).
						SetShortcut("f"),
				),
			synoptic: "[-f|--foo]",
			message:  "separates shortcut with a pipe",
		},
		{
			definition: *definition.New().
				AddOption(
					*option.New("foo", option.Required).
						SetShortcut("f"),
				),
			synoptic: "[-f|--foo FOO]",
			message:  "uses shortcut as value placeholder",
		},
		{
			definition: *definition.New().
				AddOption(
					*option.New("foo", option.Optional).
						SetShortcut("f"),
				),
			synoptic: "[-f|--foo [FOO]]",
			message:  "puts optional values in square brackets",
		},

		// testing arguments
		{
			definition: *definition.New().
				AddArgument(
					*argument.New("foo", argument.Required),
				),
			synoptic: "<foo>",
			message:  "puts arguments in angle brackets",
		},
		{
			definition: *definition.New().
				AddArgument(
					*argument.New("foo", argument.Optional),
				),
			synoptic: "[<foo>]",
			message:  "puts optional arguments in square brackets",
		},
		{
			definition: *definition.New().
				AddArgument(
					*argument.New("foo", argument.Optional),
				).
				AddArgument(
					*argument.New("bar", argument.Optional),
				),
			synoptic: "[<foo> [<bar>]]",
			message:  "chains optional arguments inside brackets",
		},
		{
			definition: *definition.New().
				AddArgument(
					*argument.New("foo", argument.List),
				),
			synoptic: "[<foo>...]",
			message:  "uses an ellipsis for array arguments",
		},
		{
			definition: *definition.New().
				AddArgument(
					*argument.New("foo", argument.List|argument.Required),
				),
			synoptic: "<foo>...",
			message:  "uses an ellipsis for required array arguments",
		},

		// testing options and arguments
		{
			definition: *definition.New().
				AddOption(*option.New("foo", option.None)).
				AddArgument(
					*argument.New("foo", argument.Required),
				),
			synoptic: "[--foo] [--] <foo>",
			message:  "puts [--] between options and arguments",
		},
	}
}
