package input

import (
	"github.com/DrSmithFr/go-console/pkg/input"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/definition"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/DrSmithFr/go-console/tests/test-helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArguments(t *testing.T) {
	in := input.NewArgvInput([]string{"cli.php", "foo"})

	in.Bind(
		*definition.New().
			AddArgument(*argument.New("name", argument.OPTIONAL)),
	)

	assert.Equal(t, map[string]string{"name": "foo"}, in.GetArguments())

	// check if stateless
	in.Bind(
		*definition.New().
			AddArgument(*argument.New("name", argument.OPTIONAL)),
	)

	assert.Equal(t, map[string]string{"name": "foo"}, in.GetArguments())
}

func TestParsePatterns(t *testing.T) {
	patterns := provideOptionsPatterns()

	// lunch tests in reverse order for easy debugging
	for i := len(patterns) - 1; i > -1; i-- {
		pattern := patterns[i]
		in := input.NewArgvInput(pattern.Argv())
		in.Bind(pattern.Definition())

		assert.Equalf(t, pattern.Options(), in.GetOptions(), pattern.Message())
		assert.Equalf(t, pattern.OptionArrays(), in.GetOptionArrays(), pattern.Message())
	}
}

func provideOptionsPatterns() []test_helper.ParserPattern {
	return []test_helper.ParserPattern{
		*test_helper.
			NewParserPattern([]string{"cli.php", "--foo"}).
			SetMessage("->parse() parses long options without a value").
			AddOption(*option.New("foo", option.NONE)).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "--foo=bar"}).
			SetMessage("->parse() parses long options with a required value (with a = separator)").
			AddOption(
				*option.
					New("foo", option.REQUIRED).
					SetShortcut("f"),
			).
			SetOptions(map[string]string{"foo": "bar"}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "--foo", "bar"}).
			SetMessage("->parse() parses long options with a required value (with a space separator)").
			AddOption(
				*option.
					New("foo", option.REQUIRED).
					SetShortcut("f"),
			).
			SetOptions(map[string]string{"foo": "bar"}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "--foo="}).
			SetMessage("->parse() parses long options with optional value which is empty (with a = separator) as empty string").
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "--foo=", "bar"}).
			SetMessage("->parse() parses long options with optional value without value specified or an empty string (with a = separator) followed by an argument as empty string").
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			AddArgument(*argument.New("name", argument.REQUIRED)).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "bar", "--foo"}).
			SetMessage("->parse() parses long options with optional value which is empty (with a = separator) preceded by an argument").
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			AddArgument(*argument.New("name", argument.REQUIRED)).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "bar", "--foo"}).
			SetMessage("->parse() parses long options with optional value which is empty (with a = separator) preceded by an argument").
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			AddArgument(*argument.New("name", argument.REQUIRED)).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "--foo", "", "bar"}).
			SetMessage("->parse() parses long options with optional value which is empty as empty string even followed by an argument").
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			AddArgument(*argument.New("name", argument.REQUIRED)).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "--foo"}).
			SetMessage("->parse() parses long options with optional value specified with no separator and no value as null").
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			AddArgument(*argument.New("name", argument.REQUIRED)).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-f"}).
			SetMessage("->parse() parses short options without a value").
			AddOption(
				*option.
					New("foo", option.NONE).
					SetShortcut("f"),
			).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-fbar"}).
			SetMessage("->parse() parses short options with a required value (with no separator)").
			AddOption(
				*option.
					New("foo", option.REQUIRED).
					SetShortcut("f"),
			).
			SetOptions(map[string]string{"foo": "bar"}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-f", "bar"}).
			SetMessage("->parse() parses short options with a required value (with a space separator)").
			AddOption(
				*option.
					New("foo", option.REQUIRED).
					SetShortcut("f"),
			).
			SetOptions(map[string]string{"foo": "bar"}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-f", ""}).
			SetMessage("->parse() parses short options with an optional empty value").
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-f", "", "foo"}).
			SetMessage("->parse() parses short options with an optional empty value followed by an argument").
			AddArgument(*argument.New("name", argument.OPTIONAL)).
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			SetOptions(map[string]string{"foo": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-f", "", "-b"}).
			SetMessage("->parse() parses short options with an optional empty value followed by an argument").
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			AddOption(
				*option.
					New("bar", option.OPTIONAL).
					SetShortcut("b"),
			).
			SetOptions(map[string]string{"foo": "", "bar": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-f", "-b", "foo"}).
			SetMessage("->parse() parses short options with an optional value which is not present").
			AddArgument(*argument.New("name", argument.OPTIONAL)).
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			AddOption(
				*option.
					New("bar", option.NONE).
					SetShortcut("b"),
			).
			SetOptions(map[string]string{"foo": "", "bar": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-fb"}).
			SetMessage("->parse() parses short options when they are aggregated as a single one").
			AddOption(
				*option.
					New("foo", option.NONE).
					SetShortcut("f"),
			).
			AddOption(
				*option.
					New("bar", option.NONE).
					SetShortcut("b"),
			).
			SetOptions(map[string]string{"foo": "", "bar": ""}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-fb", "bar"}).
			SetMessage("->parse() parses short options when they are aggregated as a single one and the last one has a required value").
			AddOption(
				*option.
					New("foo", option.NONE).
					SetShortcut("f"),
			).
			AddOption(
				*option.
					New("bar", option.REQUIRED).
					SetShortcut("b"),
			).
			SetOptions(map[string]string{"foo": "", "bar": "bar"}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-fb", "bar"}).
			SetMessage("->parse() parses short options when they are aggregated as a single one and the last one has an optional value").
			AddOption(
				*option.
					New("foo", option.NONE).
					SetShortcut("f"),
			).
			AddOption(
				*option.
					New("bar", option.OPTIONAL).
					SetShortcut("b"),
			).
			SetOptions(map[string]string{"foo": "", "bar": "bar"}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-fbbar"}).
			SetMessage("->parse() parses short options when they are aggregated as a single one and the last one has an optional value with no separator").
			AddOption(
				*option.
					New("foo", option.NONE).
					SetShortcut("f"),
			).
			AddOption(
				*option.
					New("bar", option.OPTIONAL).
					SetShortcut("b"),
			).
			SetOptions(map[string]string{"foo": "", "bar": "bar"}),

		*test_helper.
			NewParserPattern([]string{"cli.php", "-fbbar"}).
			SetMessage("->parse() parses short options when they are aggregated as a single one and one of them takes a value").
			AddOption(
				*option.
					New("foo", option.OPTIONAL).
					SetShortcut("f"),
			).
			AddOption(
				*option.
					New("bar", option.OPTIONAL).
					SetShortcut("b"),
			).
			SetOptions(map[string]string{"foo": "bbar"}),
	}
}