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

<p align="center">
    <img src="assets/question/asking-user-info.png">
</p>

If they leave it empty, the default value ("John Doe" here) is returned.

<p align="center">
    <img src="assets/question/asking-user-info-empty.png">
</p>

#### Hiding the User's Response

You can also ask a question and hide the response. This is particularly convenient for passwords:

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

	// Simple question with hidden answer
	pass := qh.Ask(
		question.
			NewQuestion("What is your password?").
			SetHidden(true),
	)
	io.Text("Password: " + pass)
}
```

<p align="center">
    <img src="assets/question/asking-user-password.png">
</p>

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

<p align="center">
    <img src="assets/question/asking-user-confirmation.png">
</p>

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

<p align="center">
    <img src="assets/question/asking-user-choice.png">
</p>

If the user enters an invalid or empty string, an error message is shown and the user is asked to provide the answer another time, 
until they enter a valid string or reach the maximum number of attempts.

<p align="center">
    <img src="assets/question/asking-user-choice-max-attempts.png">
</p>

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

<p align="center">
    <img src="assets/question/asking-user-choice-multiple.png">
</p>

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

<p align="center">
    <img src="assets/question/normalizing-input.png">
</p>

> **Note**
> 
> > Adding a custom normalizer on QuestionConfirmation and QuestionChoices will override the default one.
> > If you want to keep the default behaviour and add your own logic before or after, see the next section about ChainNormalizer.
> 
> > The normalizer is called first and the returned value is used as the input of the validator. 
> > If the answer is invalid, don't throw exceptions in the normalizer and let the validator handle those errors.

#### The ChainNormalizer

The `MakeChainNormalizer` allows you to merge multiple normalizers. Each normalizer will be called in the order they are defined.

```go
package main

import (
    "github.com/DrSmithFr/go-console/pkg/question"
    "github.com/DrSmithFr/go-console/pkg/question/normalizer"
    "github.com/DrSmithFr/go-console/pkg/style"
    "os"
    "strings"
)

func main() {
    io := style.NewConsoleCommand().Build()
    qh := question.NewHelper(os.Stdin, io.GetOutput())
  
    // Simple question with normalizer
    firstname := qh.Ask(
        question.
            NewQuestion("What is your name?").
            SetNormalizer(
                normalizer.MakeChainedNormalizer(
                    strings.ToLower,
                    normalizer.Ucfirst,
                    func(answer string) string {
                      return answer + "!"
                    },
                ),
            ),
    )
    io.Text("Hello " + firstname)
}
```

<p align="center">
    <img src="assets/question/normalizing-chain.png">
</p>

> **Note**
> 
> > a normalizer is a function that takes a string as input and returns a string as output.
> > That means you can use any function that respects this signature, from the standard library to your own functions.

With the same logic, you can encapsulate the default normalizer using the `MakeChainNormalizer` method, 
however, you will need to pass the question as a parameter of `normalizer.DefaultChoicesNormalizer()`.

```go
package main

import (
    "github.com/DrSmithFr/go-console/pkg/question"
    "github.com/DrSmithFr/go-console/pkg/question/normalizer"
    "strings"
)

func main() {
    // chain normalizer example using including the default normalizer
    q := question.
        NewChoices("What is your favorite color?", []string{"red", "blue", "yellow"}).
        SetMultiselect(true).
        SetMaxAttempts(3)

    customNormalizer := normalizer.
        MakeChainedNormalizer(
            strings.ToLower,
            q.GetDefaultNormalizer(),
            normalizer.Ucfirst,
            func(answer string) string {
              return answer + "!"
            },
        )
  
    q.SetNormalizer(customNormalizer)
}
```

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

It should return an error if there is something wrong.
The error message is displayed in the console, so it is a good practice to put some useful information in it.

<p align="center">
    <img src="assets/question/validation.png">
</p>

> **Note**
>
> > Adding a custom validator on QuestionConfirmation and QuestionChoices will override the default one.
> > If you want to keep the default behaviour and add your own logic before or after, see the next section about ChainedValidator.


#### The ChainedValidator

The `MakeChainedValidator` allows you to merge multiple validators.
Each validator will be called in the order they are defined.

```go
package main

import (
    "errors"
	"github.com/DrSmithFr/go-console/pkg/question"
	"github.com/DrSmithFr/go-console/pkg/style"
    "github.com/DrSmithFr/go-console/pkg/question/validator"
	"os"
)

func main() {
	io := style.NewConsoleCommand().Build()
	qh := question.NewHelper(os.Stdin, io.GetOutput())

    // chain validator example
    answer := qh.Ask(
        question.
            NewQuestion("What is your favorite color?").
            SetValidator(
                validator.
                    MakeChainedValidator(
                        func(answer string) error {
                          if answer == "red" {
                            return errors.New("red is mine")
                          }
  
                          return nil
                        },
                        func(answer string) error {
                          if answer == "blue" {
                            return errors.New("blue is disgusting")
                          }
  
                          return nil
                        },
                    ),
            ),
    )
    io.Text(answer)
}
```

<p align="center">
    <img src="assets/question/validation-chain.png">
</p>