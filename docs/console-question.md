## How to ask for user input
The QuestionHelper provides functions to ask the user for more information.
It can be used to ask for confirmation, to ask for a value, or to ask for a choice.

### Helper Usage

The Question Helper needs an io.Reader instance as the first argument and OutputInterface instance as the second argument.

```go
package main

import (
	"github.com/DrSmithFr/go-console/pkg/question"
	"github.com/DrSmithFr/go-console/pkg/style"
	"os"
)

func main() {
	io := style.NewConsoleCommand().Build()
	qh := question.NewHelper(os.Stdin, io.GetOutput())
}
```

### Asking the User for Information

The Question Helper has a single method ask() that takes a Question instance as its first argument and returns the user's answer as a string.

```go
package main

import (
	"github.com/DrSmithFr/go-console/pkg/question"
	"github.com/DrSmithFr/go-console/pkg/question/answers"
	"github.com/DrSmithFr/go-console/pkg/style"
	"os"
)

func main() {
	io := style.NewConsoleCommand().Build()
	qh := question.NewHelper(os.Stdin, io.GetOutput())

	// Simple question with default answer
	name := qh.Ask(
		question.
			NewQuestion("What is your name?").
			SetDefaultAnswer("John Doe"),
	)
	io.Text("Hello " + name)
}
```

The user will be asked "What is your name?". 
They can type some name which will be returned by the ask() method. 
If they leave it empty, the default value ("John Doe" here) is returned.

### Asking the User for Confirmation

Suppose you want to confirm an action before actually executing it. Add the following to your command:

```go
package main

import (
	"github.com/DrSmithFr/go-console/pkg/question"
	"github.com/DrSmithFr/go-console/pkg/question/answers"
	"github.com/DrSmithFr/go-console/pkg/style"
	"os"
)

func main() {
	io := style.NewConsoleCommand().Build()
	qh := question.NewHelper(os.Stdin, io.GetOutput())

	// Simple confirmation question
	answer := qh.Ask(
		question.
			NewComfirmation("Continue with this action?").
			SetDefaultAnswer(answers.YES).
			SetMaxAttempts(2),
	)
	if answer == answers.YES {
		io.Text("Great!")
	} else {
		io.Text("... ok :(")
	}
}
```

In this case, the user will be asked "Continue with this action?". If the user answers with y it returns "yes" (`answers.YES`) or "no" (`answers.NO`) if they answer with n. 
The default value to return if the user doesn't enter any valid input can be modify using the `SetDefaultAnswer()` (By default it is set to `answers.NONE` forcing the user to answer).

> **Note**
> 
> >You can customize the regex used to check if the answer means "yes" using the `SetYesRegex()` method. (By default it is set to `^(y|yes|true|1)$`)
> 
> >You can customize the regex used to check if the answer means "no" using the `SetNoRegex()` method. (By default it is set to `^(n|no|false|0)$`)
> 
> > You can define your own error message using `SetErrorMessage()`
> > by default it is set to `"Value '%s' is invalid"`

### Asking the User for a Choice

If you have a predefined set of answers the user can choose from, you could use a ChoiceQuestion which makes sure that the user can only enter a valid string from a predefined list:

```go
package main

import (
	"github.com/DrSmithFr/go-console/pkg/question"
	"github.com/DrSmithFr/go-console/pkg/question/answers"
	"github.com/DrSmithFr/go-console/pkg/style"
	"os"
)

func main() {
	io := style.NewConsoleCommand().Build()
	qh := question.NewHelper(os.Stdin, io.GetOutput())

	colors := []string{"red", "green", "blue", "yellow", "black", "white"}
	
	// Choice question with only one answer allowed
	answer := qh.Ask(
		question.
			NewChoices("What is your overall favorite color?", colors).
			SetMaxAttempts(3),
	)

	io.Text("Your overall favorite color is " + answer)
}
```

If the user enters an invalid or empty string, an error message is shown and the user is asked to provide the answer another time, 
until they enter a valid string or reach the maximum number of attempts.

> **Note**
> > The default value for the maximum number of attempts is 0. which means an infinite number of attempts.
> 
> >You can define your own error message using `SetErrorMessage()` 
  > by default it is set to `"Value '%s' is invalid"`

### Multiple Choices

Sometimes, multiple answers can be given. The ChoiceQuestion provides this feature using comma separated values.
This is disabled by default, to enable this use `SetMultiselect(true)`:

```go
package main

import (
	"github.com/DrSmithFr/go-console/pkg/question"
	"github.com/DrSmithFr/go-console/pkg/question/answers"
	"github.com/DrSmithFr/go-console/pkg/style"
	"os"
	"strings"
)

func main() {
	io := style.NewConsoleCommand().Build()
	qh := question.NewHelper(os.Stdin, io.GetOutput())

	colorList := []string{"red", "green", "blue", "yellow", "black", "white"}

	// Choice question with multiple answers allowed
	answer := qh.Ask(
		question.
			NewChoices("What is your favorite color?", colorList).
			SetMultiselect(true).
			SetMaxAttempts(3),
	)

	// Retrieve all selected colors by splitting the answer on commas
	colors := strings.Split(answer, ",")

	for _, color := range colors {
		io.Text("One of your favorite color is " + color)
	}
}
```

### Normalizing the Answer

Before validating the answer, you can "normalize" it to fix minor errors or tweak it as needed.
For instance, in the next example you ask for the user firstname. In case the user did capitalize the first error by mistake, 
you can modify the answer provided before validating it. To do so, configure a normalizer using the `SetNormalizer()` method:

```go
package main

import (
	"github.com/DrSmithFr/go-console/pkg/question"
	"github.com/DrSmithFr/go-console/pkg/style"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
)

func main() {
	io := style.NewConsoleCommand().Build()
	qh := question.NewHelper(os.Stdin, io.GetOutput())

	// Simple question with normalizer
	firstname := qh.Ask(
		question.
			NewQuestion("What is your name?").
			SetNormalizer(func(answer string) string {
				return cases.Title(language.English, cases.Compact).String(answer)
			}),
	)
	io.Text("Hello " + firstname)
}
```

> **Note**
> 
> > Adding a custom normalizer on QuestionConfirmation and QuestionChoices will override the default one.
> 
> > The normalizer is called first and the returned value is used as the input of the validator. 
> > If the answer is invalid, don't throw exceptions in the normalizer and let the validator handle those errors.


### Validating the Answer

You can even validate the answer. you can configure a validator using the `SetValidator()` method:

```go
package main

import (
    "errors"
	"github.com/DrSmithFr/go-console/pkg/question"
	"github.com/DrSmithFr/go-console/pkg/style"
	"os"
    "regexp"
)

func main() {
	io := style.NewConsoleCommand().Build()
	qh := question.NewHelper(os.Stdin, io.GetOutput())

  // Simple question with custom validator
  nickname := qh.Ask(
    question.
      NewQuestion("What is your nickname?").
            SetValidator(func(answer string) error {
              regex := regexp.MustCompile("^(\\w|_|-)*$")
              if match := regex.MatchString(answer); !match {
                return errors.New("nickname must be alphanumeric")
              }

              return nil
            }),
  )
  io.Text("Hi " + nickname)
}
```

The validator is a callback which handles the validation. 
It should return an exception if there is something wrong.
The exception message is displayed in the console, so it is a good practice to put some useful information in it.

> **Note**
>
> > Adding a custom validator on QuestionConfirmation and QuestionChoices will override the default one.