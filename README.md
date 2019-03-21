<p align="center">
    <img src="assets/icon.png">
</p>

# How to Style a Console Command

[![CircleCI](https://circleci.com/gh/DrSmithFr/go-console.svg?style=shield)](https://circleci.com/gh/DrSmithFr/go-console)
[![GolangCI](https://golangci.com/badges/github.com/DrSmithFr/go-console.svg)](https://golangci.com/r/github.com/DrSmithFr/go-console)
[![Go Report Card](https://goreportcard.com/badge/github.com/DrSmithFr/go-console)](https://goreportcard.com/report/github.com/DrSmithFr/go-console)

One of the most boring tasks when creating console commands is to deal with the styling of the command's output, 
this library provide several helper for that. 

GoConsole is the Go equivalent to the [Console Component](https://github.com/symfony/console) of Symfony PHP framework.

## How to Color the Console Output

By using colors in the command output, you can distinguish different types of output (e.g. important messages, titles, comments, etc.).

```go
package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/color"
	"github.com/DrSmithFr/go-console/pkg/formatter"
)

func main() {
	// create a default style
	s1 := formatter.NewOutputFormatterStyle(color.NULL, color.NULL, nil)
	fmt.Printf(s1.Apply("some text without coloration\n"))

	s1.SetBackground(color.RED)
	fmt.Printf(s1.Apply("some text with red background\n"))

	s1.SetForeground(color.GREEN)
	fmt.Printf(s1.Apply("some text with red background and green text\n"))

	s1.SetOption(color.BOLD)
	fmt.Printf(s1.Apply("some bold text with red background and green text \n"))

	// override all options in one time
	s1.SetOptions([]string{color.UNDERSCORE})
	fmt.Printf(s1.Apply("some underscore text with red background and green text \n"))

	// quick declaration
	s2 := formatter.NewOutputFormatterStyle(color.BLUE, color.YELLOW, nil)
	fmt.Printf(s2.Apply("some text with yellow background and blue text\n"))

	// quick declaration with options
	s3 := formatter.NewOutputFormatterStyle(color.DEFAULT, color.DEFAULT, []string{color.UNDERSCORE, color.BOLD})
	fmt.Printf(s3.Apply("some bold and underscore text\n"))
}
```

<p align="center">
    <img src="assets/direct-use-formatter-style.png">
</p>

The OutputFormatterStyle is the simplest way to color output. It is not mean to be use directly, but to defined custom tags used by OutputFormatterInterface.

---

Whenever you output text, you can use OutputInterface to surround the text with tags to color its output. For example:

```go
package main

import "github.com/DrSmithFr/go-console/pkg/output"

func main() {
    // creating new output
    out := output.NewConsoleOutput(true, nil)
    
    // white text on a red background
    out.Writeln("<error>An error</error>")
    
    // green text
    out.Writeln("<info>An information</info>")
    
    // yellow text
    out.Writeln("<comment>An comment</comment>")
    
    // black text on a cyan background
    out.Writeln("<question>A question</question>")
    
    // underscore text
    out.Writeln("<u>Some underscore text</u>")
    
    // bold text
    out.Writeln("<b>Some bold text</b>")
}
```

<p align="center">
    <img src="assets/test-default-output-console.png">
</p>

> The closing tag can be replaced by </>, which revokes all formatting options established by the last opened tag.

---

It is possible to define your own styles using the OutputFormatterStyle

```go
package main

import (
	"github.com/DrSmithFr/go-console/pkg/color"
	"github.com/DrSmithFr/go-console/pkg/formatter"
	"github.com/DrSmithFr/go-console/pkg/output"
)

func main() {
	// creating new output
	out := output.NewConsoleOutput(true, nil)

	// create new style
	s := formatter.NewOutputFormatterStyle(color.RED, color.YELLOW, []string{color.BOLD, color.BLINK})

	// add style to formatter
	out.GetFormatter().SetStyle("fire", *s)

	// use the new style
	out.Writeln("<fire>foo</>")
}
```

<p align="center">
    <img src="assets/custom-console-style.png">
</p>

> Available foreground and background colors are: black, red, green, yellow, blue, magenta, cyan and white.
> And available options are: bold, underscore, blink, reverse (enables the "reverse video" mode where the background and foreground colors are swapped) and conceal (sets the foreground color to transparent, making the typed text invisible - although it can be selected and copied; this option is commonly used when asking the user to type sensitive information).

---

You can also set these colors and options directly inside the tag name:

```go
package main

import "github.com/DrSmithFr/go-console/pkg/output"

func main() {
    // green text
    out := output.NewConsoleOutput(true, nil)
    
    // black text on a cyan background
    out.Writeln("<fg=green>foo</>")
    
    // green text
    out.Writeln("<fg=black;bg=cyan>foo</>")
    
    // bold text on a yellow background
    out.Writeln("<bg=yellow;options=bold>foo</>")
    
    // bold text with underscore
    out.Writeln("<options=bold,underscore>foo</>")
}
```

<p align="center">
    <img src="assets/custom-style-tags.png">
</p>

> If you need to render a tag literally, escape it with a backslash: \<info> or use the escape() method to escape all the tags included in the given string.
