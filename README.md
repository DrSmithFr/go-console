<p align="center">
    <img src="icon.png">
</p>

# How to Style a Console Command

[![CircleCI](https://circleci.com/gh/MrSmith777/go-console.svg?style=shield)](https://circleci.com/gh/MrSmith777/go-console)
[![GolangCI](https://golangci.com/badges/github.com/MrSmith777/go-console.svg)](https://golangci.com/r/github.com/MrSmith777/go-console)
[![Go Report Card](https://goreportcard.com/badge/github.com/MrSmith777/go-console)](https://goreportcard.com/report/github.com/MrSmith777/go-console)

One of the most boring tasks when creating console commands is to deal with the styling of the command's output, 
this library provide several helper for that. 

GoConsole is the Go equivalent to the [Console Component](https://github.com/symfony/console) of Symfony PHP framework.

## How to Color the Console Output

By using colors in the command output, you can distinguish different types of output (e.g. important messages, titles, comments, etc.).

Whenever you output text, you can surround the text with tags to color its output. For example:

```go
import "github.com/MrSmith777/go-console/pkg/output"

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

> The closing tag can be replaced by </>, which revokes all formatting options established by the last opened tag.

---

It is possible to define your own styles using the OutputFormatterStyle

```go
import (
    "github.com/MrSmith777/go-console/pkg/output"
    "github.com/MrSmith777/go-console/pkg/color"
)

func main() {
    // create new style
    s := formatter.NewOutputFormatterStyle(color.RED, color.YELLOW, []string{color.BOLD, color.BLINK})
 
    // add style to formatter
    out.GetFormatter().SetStyle("fire", *s)

    // use the new style
    out.Writeln("<fire>foo</>")
}
```

> Available foreground and background colors are: black, red, green, yellow, blue, magenta, cyan and white.
> And available options are: bold, underscore, blink, reverse (enables the "reverse video" mode where the background and foreground colors are swapped) and conceal (sets the foreground color to transparent, making the typed text invisible - although it can be selected and copied; this option is commonly used when asking the user to type sensitive information).

---

You can also set these colors and options directly inside the tag name:

```go
import "github.com/MrSmith777/go-console/pkg/output"

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

> If you need to render a tag literally, escape it with a backslash: \<info> or use the escape() method to escape all the tags included in the given string.

## How to Style the Console Output

One of the most boring tasks when creating console commands is to deal with the styling of the command's input and output. Displaying titles and tables or asking questions to the user involves a lot of repetitive code.

Consider for example the code used to display the title of the following command:

```go
import "github.com/MrSmith777/go-console/pkg/output"

func main() {
    // creating new output
    out := output.NewConsoleOutput(true, nil)
    
    out.Writeln("<info>Lorem Ipsum Dolor Sit Amet</>")
    out.Writeln("<info>==========================</>\n")
}
```

Displaying a simple title requires two lines of code, to change the font color, underline the contents and leave an additional blank line after the title. Dealing with styles is required for well-designed commands, but it complicates their code unnecessarily.

In order to reduce that boilerplate code, go-console can optionally use the Go Style Guide. These styles are implemented as a set of helper methods which allow to create semantic commands and forget about their styling.

### Basic Usage

```go
package main

import (
	"github.com/MrSmith777/go-console/pkg/output"
	"github.com/MrSmith777/go-console/pkg/style"
)

func main() {
	// creating new output
	out := output.NewConsoleOutput(true, nil)

	// creating new styler
	io := style.NewGoStyler(out)

	// use Go styler
	io.Title("Lorem Ipsum Dolor Sit Amet")
}
```

### Helper Methods

#### Titling Methods

##### title()

It displays the given string as the command title. This method is meant to be used only once in a given command, but nothing prevents you to use it repeatedly:

```go
io.Title("Lorem Ipsum Dolor Sit Amet")
```

