## Console Input (Arguments & Options)

The most interesting part of the commands are the arguments and options that you can make available. These arguments and
options allow you to pass dynamic information from the terminal to the command.

### Using Command Arguments

Arguments are the strings - separated by spaces - that come after the command name itself. They are ordered, and can be
optional or required. For example, to add an optional `last_name` argument to the command and make the `name` argument
required:

```go
package main

import (
	"fmt"
	"DrSmithFr/go-console/pkg/input/argument"
	"DrSmithFr/go-console/pkg/style"
)

func main() {
	io := style.
		NewConsoleStyler().
		AddInputArgument(
			argument.
				New("name", argument.REQUIRED).
				SetDescription("Who do you want to greet?"),
		).
		AddInputArgument(
			argument.
				New("last_name", argument.OPTIONAL).
				SetDescription("Your last name?"),
		).
		ParseInput().
		ValidateInput()

	//
	// You now have access to a last_name argument in your command:
	//

	text := fmt.Sprintf("Hi %s", io.GetInput().GetArgument("name"))

	lastName := io.GetInput().GetArgument("last_name")

	if lastName != "" {
		text = fmt.Sprintf("%s %s", text, lastName)
	}

	io.GetOutput().Write(text)
	io.GetOutput().Writeln("!")
}
```

The command can now be used in either of the following ways:

```
go run command-script John
> Hi John!

go run command John Smith
> Hi John Daligault!
```

---

It is also possible to let an argument take a list of values (imagine you want to greet all your friends). Only the last
argument can be a list:

```go
package main

import (
	"fmt"
	"DrSmithFr/go-console/pkg/input/argument"
	"DrSmithFr/go-console/pkg/style"
)

func main() {
	io := style.
		NewConsoleStyler().
		AddInputArgument(
			argument.
				New("names", argument.IS_ARRAY).
				SetDescription("Who do you want to greet?"),
		).
		ParseInput().
		ValidateInput()

	//
	// You can access the names argument as an array:
	//

	names := io.GetInput().GetArgumentArray("names")

	for _, name := range names {
		io.Text(fmt.Sprintf("Hi %s!", name))
	}
}
```

To use this, specify as many names as you want:

```go
go run command-script John Alex Fred
```

---

There are three argument variants you can use:

`argument.REQUIRED`
> The argument is mandatory. The command doesn't run if the argument isn't provided;

`argument.OPTIONAL`
> The argument is optional and therefore can be omitted. This is the default behavior of arguments;

`argument.IS_ARRAY`
> The argument can contain any number of values. For that reason, it must be used at the end of the argument list.

You can combine `IS_ARRAY` with `REQUIRED` and `OPTIONAL` like this:

```go
io := style.
NewConsoleStyler().
AddInputArgument(
argument.
New("names", argument.IS_ARRAY | argument.REQUIRED),
).
ParseInput().
ValidateInput()
```

### Using Command Options

Unlike arguments, options are not ordered (meaning you can specify them in any order) and are specified with two
dashes (e.g. `--yell`). Options are always optional, and can be setup to accept a value (e.g. `--dir=src`) or as a
boolean flag without a value (e.g.  `--yell`).

For example, add a new option to the command that can be used to specify how many times in a row the message should be
printed:

```go
package main

import (
	"fmt"
	"DrSmithFr/go-console/pkg/input/argument"
	"DrSmithFr/go-console/pkg/input/option"
	"DrSmithFr/go-console/pkg/style"
	"strconv"
)

func main() {
	io := style.
		NewConsoleStyler().
		AddInputArgument(
			argument.
				New("name", argument.REQUIRED).
				SetDescription("Who do you want to greet?"),
		).
		AddInputOption(
			option.
				New("iterations", option.REQUIRED).
				SetDescription("How many times should the message be printed?").
				SetDefault("1"),
		).
		ParseInput().
		ValidateInput()

	//
	// Next, use this in the command to print the message multiple times:
	//

	iterations, _ := strconv.Atoi(io.GetInput().GetOption("iterations"))

	for i := 0; i < iterations; i++ {
		io.Text(
			fmt.Sprintf("Hi %s!", io.GetInput().GetArgument("name")),
		)
	}
}
```

Now, when you run the command, you can optionally specify a `--iterations` flag:

```
# no --iterations provided, the default (1) is used
$ php bin/console app:greet John
 Hi John!

$ php bin/console app:greet John --iterations=5
 Hi John!
 Hi John!
 Hi John!
 Hi John!
 Hi John!


$ php bin/console app:greet John --iterations=5 --yell
$ php bin/console app:greet John --yell --iterations=5
$ php bin/console app:greet --yell --iterations=5 John
```

---

You can also declare a one-letter shortcut that you can call with a single dash, like `-i`:

```go
io := style.
NewConsoleStyler().
AddInputOption(
option.
New("iterations", option.REQUIRED).
SetShortcut("i"),
).
ParseInput().
ValidateInput()
```

Note that to comply with the [docopt standard](http://docopt.org/), long options can specify their values after a white
space or an = sign (e.g. `--iterations 5` or `--iterations=5`), but short options can only use white spaces or no
separation at all (e.g. `-i 5` or `-i5`).

> While it is possible to separate an option from its value with a white space, using this form leads to an ambiguity should the option appear before the command name.
> For example, `php bin/console --iterations 5 app:greet Fabien` is ambiguous; Go-console would interpret 5 as the command name. To avoid this situation, always place options after the command name, or avoid using a space to separate the option name from its value.

---

There are four option variants you can use:

`option.IS_ARRAY`
> This option accepts multiple values (e.g. `--dir=/foo --dir=/bar`);

`argument.NONE`
> Do not accept input for this option (e.g. `--yell`). This is the default behavior of options;

`argument.REQUIRED`
> This value is required (e.g. `--iterations=5` or `-i5`), the option itself is still optional;

`argument.OPTIONAL`
> This option may or may not have a value (e.g. `--yell` or `--yell=loud`).

You can combine `IS_ARRAY` with `REQUIRED` and `OPTIONAL` like this:

```go
io := style.
NewConsoleStyler().
AddInputOption(
option.New("iterations", option.IS_ARRAY | option.REQUIRED),
).
ParseInput().
ValidateInput()
```