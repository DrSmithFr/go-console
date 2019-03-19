package formatter

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/formatter"
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
