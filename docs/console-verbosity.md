## How to use verbosity levels
Console commands have different verbosity levels, which determine the messages displayed in their output. 
By default, commands display only the most useful messages, 
but you can control their verbosity with the `--quiet|-q`, `--verbose|-v`, `--very-verbose|-vv`, `--debug|-vvv` options.

### Basic Usage

```go
package main

import (
	"github.com/DrSmithFr/go-console/pkg/input"
	"github.com/DrSmithFr/go-console/pkg/output"
	"github.com/DrSmithFr/go-console/pkg/style"
	"github.com/DrSmithFr/go-console/pkg/verbosity"
)

func main() {
	io := style.NewConsoleCommand()

	if io.GetVerbosity() == verbosity.Verbose {
		io.Text("Lorem Ipsum Dolor Sit Amet")
	}

	// available methods: .IsQuiet(), .IsVerbose(), .IsVeryVerbose(), .IsDebug()
	if io.IsVeryVerbose() {
		io.Text("Lorem Ipsum Dolor Sit Amet")
	}

	// or using directly the output instance
	out := io.GetOutput()

	if out.GetVerbosity() == verbosity.Verbose {
		out.Writeln("Lorem Ipsum Dolor Sit Amet")
	}

	// available methods: .IsQuiet(), .IsVerbose(), .IsVeryVerbose(), .IsDebug()
	if out.IsVeryVerbose() {
		out.Writeln("Lorem Ipsum Dolor Sit Amet")
	}
}
```

When the quiet level is used, all output is suppressed as the default write() method returns without actually printing.
