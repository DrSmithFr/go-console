package input

import (
	"github.com/DrSmithFr/go-console/pkg/input"
	"github.com/DrSmithFr/go-console/pkg/input/argument"
	"github.com/DrSmithFr/go-console/pkg/input/option"
	"github.com/DrSmithFr/go-console/tests/test-helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArguments(t *testing.T) {
	in := input.NewArgvInput([]string{"cli.php", "foo"})

	in.Bind(
		*input.NewInputDefinition().
			AddArgument(*argument.NewInputArgument("name", argument.OPTIONAL)),
	)

	assert.Equal(t, map[string]string{"name": "foo"}, in.GetArguments())

	// check if stateless
	in.Bind(
		*input.NewInputDefinition().
			AddArgument(*argument.NewInputArgument("name", argument.OPTIONAL)),
	)

	assert.Equal(t, map[string]string{"name": "foo"}, in.GetArguments())
}

func TestParsePatterns(t *testing.T) {
	for _, pattern := range provideParserPatterns() {
		in := input.NewArgvInput(pattern.Argv())
		in.Bind(pattern.Definition())

		assert.Equal(t, pattern.Arguments(), in.GetArguments())
		assert.Equal(t, pattern.ArgumentArrays(), in.GetArgumentArrays())

		assert.Equal(t, pattern.Options(), in.GetOptions())
		assert.Equal(t, pattern.OptionArrays(), in.GetOptionArrays())
	}
}

func provideParserPatterns() []test_helper.ParserPattern {
	return []test_helper.ParserPattern{
		*test_helper.
			NewParserPattern([]string{"cli.php", "--foo"}).
			AddOption(*option.NewInputOption("foo", option.NONE)).
			SetOptions(map[string]string{"foo": ""}),
	}
}
