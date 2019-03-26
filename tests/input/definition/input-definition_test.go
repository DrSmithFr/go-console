package definition

import (
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/definition"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

var arguments = map[string]argument.InputArgument{
	"foo":  * argument.New("foo", argument.OPTIONAL),
	"bar":  * argument.New("bar", argument.OPTIONAL),
	"foo1": * argument.New("foo", argument.OPTIONAL),
	"foo2": * argument.New("foo2", argument.REQUIRED),
}

var options = map[string]option.InputOption{
	"foo": * option.
		New("foo", option.OPTIONAL).
		SetShortcut("f"),
	"bar": * option.
		New("bar", option.OPTIONAL).
		SetShortcut("b"),
	"foo1": * option.
		New("fooBis", option.OPTIONAL).
		SetShortcut("f"),
	"foo2": * option.
		New("foo", option.OPTIONAL).
		SetShortcut("p"),
	"multi": * option.
		New("multi", option.OPTIONAL).
		SetShortcut("m|mm|mmm"),
}

func TestConstructorArguments(t *testing.T) {
	def1 := definition.New()
	assert.Equal(t, map[string]argument.InputArgument{}, def1.GetArguments())
	assert.Equal(t, map[string]option.InputOption{}, def1.GetOptions())
}

func TestSetArguments(t *testing.T) {
	def := definition.
		New().
		SetArguments([]argument.InputArgument{
			arguments["foo"],
		})

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"foo": arguments["foo"],
		},
		def.GetArguments(),
	)

	def.SetArguments([]argument.InputArgument{
		arguments["bar"],
	})

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"bar": arguments["bar"],
		},
		def.GetArguments(),
	)
}

func TestAddArguments(t *testing.T) {
	def := definition.
		New().
		AddArguments([]argument.InputArgument{
			arguments["foo"],
		})

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"foo": arguments["foo"],
		},
		def.GetArguments(),
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
		def.GetArguments(),
	)
}

func TestAddArgument(t *testing.T) {
	def := definition.
		New().
		AddArgument(arguments["foo"])

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"foo": arguments["foo"],
		},
		def.GetArguments(),
	)

	def.AddArgument(arguments["bar"])

	assert.Equal(
		t,
		map[string]argument.InputArgument{
			"foo": arguments["foo"],
			"bar": arguments["bar"],
		},
		def.GetArguments(),
	)
}

func TestArgumentsMustHaveDifferentNames(t *testing.T) {
	assert.Panics(t, func() {
		definition.
			New().
			AddArgument(arguments["foo"]).
			AddArgument(arguments["foo1"])
	})
}

func TestArrayArgumentHasToBeLast(t *testing.T) {
	assert.Panics(t, func() {
		definition.
			New().
			AddArgument(*argument.New("fooarray", argument.IS_ARRAY)).
			AddArgument(*argument.New("anotherbar", argument.OPTIONAL))
	})
}

func TestRequiredArgumentCannotFollowAnOptionalOne(t *testing.T) {
	assert.Panics(t, func() {
		definition.
			New().
			AddArgument(arguments["foo"]).
			AddArgument(arguments["foo2"])
	})
}

func TestGetArgument(t *testing.T) {
	def := definition.
		New().
		AddArgument(arguments["foo"])

	assert.Equal(t, arguments["foo"], *def.GetArgument("foo"))
}

func TestGetInvalidArgument(t *testing.T) {
	assert.Panics(t, func() {
		definition.
			New().
			AddArgument(arguments["foo"]).
			GetArgument("bar")
	})
}

func TestHasArgument(t *testing.T) {
	def := definition.
		New().
		AddArgument(arguments["foo"])

	assert.True(t, def.HasArgument("foo"))
	assert.False(t, def.HasArgument("bar"))
}

func TestGetArgumentRequiredCount(t *testing.T) {
	def := definition.
		New().
		AddArgument(arguments["foo2"])

	assert.Equal(t, 1, def.GetArgumentRequiredCount())

	def.AddArgument(arguments["foo"])

	assert.Equal(t, 1, def.GetArgumentRequiredCount())
}

func TestGetArgumentCount(t *testing.T) {
	def := definition.
		New().
		AddArgument(arguments["foo2"])

	assert.Equal(t, 1, def.GetArgumentCount())

	def.AddArgument(arguments["foo"])

	assert.Equal(t, 2, def.GetArgumentCount())
}

func TestGetArgumentDefaults(t *testing.T) {
	def := definition.
		New().
		SetArguments([]argument.InputArgument{
			*argument.
				New("foo1", argument.OPTIONAL),

			*argument.
				New("foo2", argument.OPTIONAL).
				SetDefault("default"),

			*argument.
				New("foo3", argument.OPTIONAL|argument.IS_ARRAY),
		})

	validation := map[string][]string{
		"foo1": nil,
		"foo2": {"default"},
		"foo3": {},
		"foo4": {"1", "2"},
	}

	assert.Equal(t, validation["foo1"], def.GetArgumentDefaults()["foo1"])
	assert.Equal(t, validation["foo2"], def.GetArgumentDefaults()["foo2"])
	assert.Equal(t, validation["foo3"], def.GetArgumentDefaults()["foo3"])

	def2 := definition.
		New().
		SetArguments([]argument.InputArgument{
			*argument.
				New("foo4", argument.OPTIONAL|argument.IS_ARRAY).
				SetDefaults([]string{"1", "2"}),
		})

	assert.Equal(t, validation["foo4"], def2.GetArgumentDefaults()["foo4"])
}

func TestSetOptions(t *testing.T) {
	def := definition.
		New().
		SetOptions([]option.InputOption{
			options["foo"],
		})

	assert.Equal(
		t,
		map[string]option.InputOption{
			"foo": options["foo"],
		},
		def.GetOptions(),
	)

	def.SetOptions([]option.InputOption{
		options["bar"],
	})

	assert.Equal(
		t,
		map[string]option.InputOption{
			"bar": options["bar"],
		},
		def.GetOptions(),
	)
}

func TestSetOptionsClearsOptions(t *testing.T) {
	assert.Panics(t, func() {
		definition.
			New().
			SetOptions([]option.InputOption{
				options["bar"],
			}).
			GetOptionForShortcut("f")
	})
}

func TestAddOptions(t *testing.T) {
	def := definition.
		New().
		AddOptions([]option.InputOption{
			options["foo"],
		})

	assert.Equal(
		t,
		map[string]option.InputOption{
			"foo": options["foo"],
		},
		def.GetOptions(),
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
		def.GetOptions(),
	)
}

func TestAddOption(t *testing.T) {
	def := definition.
		New().
		AddOption(options["foo"])

	assert.Equal(
		t,
		map[string]option.InputOption{
			"foo": options["foo"],
		},
		def.GetOptions(),
	)

	def.AddOption(options["bar"])

	assert.Equal(
		t,
		map[string]option.InputOption{
			"foo": options["foo"],
			"bar": options["bar"],
		},
		def.GetOptions(),
	)
}

func TestAddDuplicateOption(t *testing.T) {
	assert.Panics(t, func() {
		definition.
			New().
			AddOption(options["foo"]).
			AddOption(options["foo2"])
	})
}

func TestAddDuplicateShortcutOption(t *testing.T) {
	assert.Panics(t, func() {
		definition.
			New().
			AddOption(options["foo"]).
			AddOption(options["foo1"])
	})
}

func TestGetOption(t *testing.T) {
	def := definition.
		New().
		AddOption(options["foo"])

	assert.Equal(t, options["foo"], *def.GetOption("foo"))
}

func TestGetInvalidOption(t *testing.T) {
	assert.Panics(t, func() {
		definition.
			New().
			AddOption(options["foo"]).
			GetOption("bar")
	})
}

func TestHasOption(t *testing.T) {
	def := definition.
		New().
		AddOption(options["foo"])

	assert.True(t, def.HasOption("foo"))
	assert.False(t, def.HasOption("BAR"))
}

func TestHasShortcut(t *testing.T) {
	def := definition.
		New().
		AddOption(options["foo"])

	assert.True(t, def.HasShortcut("f"))
	assert.False(t, def.HasShortcut("b"))
}

func TestGetOptionForShortcut(t *testing.T) {
	def := definition.
		New().
		AddOption(options["foo"])

	assert.Equal(t, options["foo"], *def.GetOptionForShortcut("f"))
}

func TestGetOptionForMultiShortcut(t *testing.T) {
	def := definition.
		New().
		AddOption(options["multi"])

	assert.Equal(t, options["multi"], *def.GetOptionForShortcut("m"))
	assert.Equal(t, options["multi"], *def.GetOptionForShortcut("mm"))
	assert.Equal(t, options["multi"], *def.GetOptionForShortcut("mmm"))
}

func TestGetOptionForInvalidShortcut(t *testing.T) {
	assert.Panics(t, func() {
		definition.
			New().
			GetOptionForShortcut("l")
	})
}

func TestGetOptionDefaults(t *testing.T) {
	def := definition.
		New().
		SetOptions([]option.InputOption{
			*option.
				New("foo1", option.NONE),

			*option.
				New("foo2", option.REQUIRED),

			*option.
				New("foo3", option.REQUIRED).
				SetDefault("default"),

			*option.
				New("foo4", option.OPTIONAL),

			*option.
				New("foo5", option.OPTIONAL).
				SetDefault("default"),

			*option.
				New("foo6", option.OPTIONAL|option.IS_ARRAY),

			*option.
				New("foo7", option.OPTIONAL|option.IS_ARRAY).
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

	assert.Equal(t, validation["foo1"], def.GetOptionDefaults()["foo1"])
	assert.Equal(t, validation["foo2"], def.GetOptionDefaults()["foo2"])
	assert.Equal(t, validation["foo3"], def.GetOptionDefaults()["foo3"])
	assert.Equal(t, validation["foo4"], def.GetOptionDefaults()["foo4"])
	assert.Equal(t, validation["foo5"], def.GetOptionDefaults()["foo5"])
	assert.Equal(t, validation["foo6"], def.GetOptionDefaults()["foo6"])
	assert.Equal(t, validation["foo7"], def.GetOptionDefaults()["foo7"])
}

func TestGetSynopsis(t *testing.T) {
	for _, pattern := range getSynopticPattern() {
		assert.Equalf(t, pattern.synoptic, pattern.definition.GetSynopsis(false), pattern.message)
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
			definition: *definition.
				New().
				AddOption(*option.New("foo", option.NONE)),
			synoptic: "[--foo]",
			message:  "puts optional options in square brackets",
		},
		{
			definition: *definition.
				New().
				AddOption(
					*option.
						New("foo", option.NONE).
						SetShortcut("f"),
				),
			synoptic: "[-f|--foo]",
			message:  "separates shortcut with a pipe",
		},
		{
			definition: *definition.
				New().
				AddOption(
					*option.
						New("foo", option.REQUIRED).
						SetShortcut("f"),
				),
			synoptic: "[-f|--foo FOO]",
			message:  "uses shortcut as value placeholder",
		},
		{
			definition: *definition.
				New().
				AddOption(
					*option.
						New("foo", option.OPTIONAL).
						SetShortcut("f"),
				),
			synoptic: "[-f|--foo [FOO]]",
			message:  "puts optional values in square brackets",
		},

		// testing arguments
		{
			definition: *definition.
				New().
				AddArgument(
					*argument.
						New("foo", argument.REQUIRED),
				),
			synoptic: "<foo>",
			message:  "puts arguments in angle brackets",
		},
		{
			definition: *definition.
				New().
				AddArgument(
					*argument.
						New("foo", argument.OPTIONAL),
				),
			synoptic: "[<foo>]",
			message:  "puts optional arguments in square brackets",
		},
		{
			definition: *definition.
				New().
				AddArgument(
					*argument.
						New("foo", argument.OPTIONAL),
				).
				AddArgument(
					*argument.
						New("bar", argument.OPTIONAL),
				),
			synoptic: "[<foo> [<bar>]]",
			message:  "chains optional arguments inside brackets",
		},
		{
			definition: *definition.
				New().
				AddArgument(
					*argument.
						New("foo", argument.IS_ARRAY),
				),
			synoptic: "[<foo>...]",
			message:  "uses an ellipsis for array arguments",
		},
		{
			definition: *definition.
				New().
				AddArgument(
					*argument.
						New("foo", argument.IS_ARRAY|argument.REQUIRED),
				),
			synoptic: "<foo>...",
			message:  "uses an ellipsis for required array arguments",
		},

		// testing options and arguments
		{
			definition: *definition.
				New().
				AddOption(*option.New("foo", option.NONE)).
				AddArgument(
					*argument.
						New("foo", argument.REQUIRED),
				),
			synoptic: "[--foo] [--] <foo>",
			message:  "puts [--] between options and arguments",
		},
	}
}
