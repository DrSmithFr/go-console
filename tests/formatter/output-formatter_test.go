package formatter

import (
	"fmt"
	"github.com/DrSmithFr/go-console/color"
	"github.com/DrSmithFr/go-console/formatter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyTag(t *testing.T) {
	format := formatter.NewOutputFormatter()

	assert.Equal(t, "foo<>bar", format.Format("foo<>bar"))
}

func TestLGCharEscaping(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	assert.Equal(t, "foo<>bar", format.Format("foo<>bar"))
	assert.Equal(t, "foo << bar", format.Format("foo << bar"))
	assert.Equal(t, "foo << bar \\", format.Format("foo << bar \\"))

	assert.Equal(
		t,
		"foo << \033[32mbar \\ baz\033[39m \\",
		format.Format("foo << <info>bar \\ baz</info> \\"),
	)

	assert.Equal(t, "<info>some info</info>", format.Format("\\<info>some info\\</info>"))

	assert.Equal(
		t,
		"\\<info>some info\\</info>",
		formatter.Escape("<info>some info</info>"),
	)
	assert.Equal(
		t,
		"\033[33mSymfony\\Component\\Console does work very well!\033[39m",
		format.Format("<comment>Symfony\\Component\\Console does work very well!</comment>"),
	)
}

func TestBundledStyles(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	assert.True(t, format.HasStyle("error"))
	assert.True(t, format.HasStyle("info"))
	assert.True(t, format.HasStyle("comment"))
	assert.True(t, format.HasStyle("question"))
	assert.True(t, format.HasStyle("b"))
	assert.True(t, format.HasStyle("u"))

	assert.Equal(
		t,
		"\033[37;41msome error\033[39;49m",
		format.Format("<error>some error</error>"),
	)
	assert.Equal(
		t,
		"\033[32msome info\033[39m",
		format.Format("<info>some info</info>"),
	)
	assert.Equal(
		t,
		"\033[33msome comment\033[39m",
		format.Format("<comment>some comment</comment>"),
	)
	assert.Equal(
		t,
		"\033[30;46msome question\033[39;49m",
		format.Format("<question>some question</question>"),
	)
	assert.Equal(
		t,
		"\033[1msome bold\033[22m",
		format.Format("<b>some bold</b>"),
	)
	assert.Equal(
		t,
		"\033[4msome underscored\033[24m",
		format.Format("<u>some underscored</u>"),
	)
}

func TestNestedStyles(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	assert.Equal(
		t,
		"\033[37;41msome \033[39;49m\033[32msome info\033[39m\033[37;41m error\033[39;49m",
		format.Format("<error>some <info>some info</info> error</error>"),
	)
}

func TestAdjacentStyles(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	assert.Equal(
		t,
		"\033[37;41msome error\033[39;49m\033[32msome info\033[39m",
		format.Format("<error>some error</error><info>some info</info>"),
	)
}

func TestStyleMatchingNotGreedy(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	assert.Equal(
		t,
		"(\033[32m>=2.0,<2.3\033[39m)",
		format.Format("(<info>>=2.0,<2.3</info>)"),
	)
}

func TestStyleEscaping(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	assert.Equal(
		t,
		"(\033[32mz>=2.0,<<<a2.3\\\033[39m)",
		format.Format(
			fmt.Sprintf(
				"(<info>%s</info>)",
				formatter.Escape("z>=2.0,<\\<<a2.3\\"),
			),
		),
	)
	assert.Equal(
		t,
		"\033[32m<error>some error</error>\033[39m",
		format.Format(
			fmt.Sprintf(
				"<info>%s</info>",
				formatter.Escape("<error>some error</error>"),
			),
		),
	)
}

func TestDeepNestedStyles(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	assert.Equal(
		t,
		"\033[37;41merror\033[39;49m\033[32minfo\033[39m\033[33mcomment\033[39m\033[37;41merror\033[39;49m",
		format.Format("<error>error<info>info<comment>comment</info>error</error>"),
	)
}

func TestNewStyle(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	s1 := formatter.NewOutputFormatterStyle(color.Blue, color.White, nil)
	format.SetStyle("test", *s1)

	assert.Equal(t, s1, format.GetStyle("test"))
	assert.NotEqual(t, s1, format.GetStyle("info"))

	s2 := formatter.NewOutputFormatterStyle(color.Blue, color.White, nil)
	format.SetStyle("b", *s2)

	assert.Equal(
		t,
		"\033[34;47msome \033[39;49m\033[34;47mcustom\033[39;49m\033[34;47m msg\033[39;49m",
		format.Format("<test>some <b>custom</b> msg</test>"),
	)
}

func TestRedefineStyle(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	s := formatter.NewOutputFormatterStyle(color.Blue, color.White, nil)
	format.SetStyle("info", *s)

	assert.Equal(
		t,
		"\033[34;47msome custom msg\033[39;49m",
		format.Format("<info>some custom msg</info>"),
	)
}

func TestInlineStyle(t *testing.T) {
	format := formatter.NewOutputFormatter()
	format.SetDecorated(true)

	assert.Equal(
		t,
		"\033[34;41msome text\033[39;49m",
		format.Format("<fg=blue;bg=red>some text</>"),
	)

	assert.Equal(
		t,
		"\033[34;41msome text\033[39;49m",
		format.Format("<fg=blue;bg=red>some text</fg=blue;bg=red>"),
	)

	assert.Equal(
		t,
		"\033[39;49;4msome text\033[39;49;24m",
		format.Format("<options=underscore>some text</>"),
	)

	assert.Equal(
		t,
		"\033[39;49;1;4msome text\033[39;49;22;24m",
		format.Format("<options=underscore,bold>some text</>"),
	)
}
