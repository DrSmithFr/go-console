package main

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/question"
	"github.com/DrSmithFr/go-console/pkg/question/answers"
	"github.com/DrSmithFr/go-console/pkg/style"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"regexp"
	"strings"
)

func main() {
	io := style.NewConsoleCommand().Build()
	qh := question.NewHelper(os.Stdin, io.GetOutput())

	// Simple question with default answer
	firstname := qh.Ask(
		question.
			NewQuestion("What is your firstname?").
			SetDefaultAnswer("Doe"),
	)
	io.Text("Hello " + firstname)

	// Simple question with normalizer
	lastname := qh.Ask(
		question.
			NewQuestion("What is your lastname?").
			SetNormalizer(func(answer string) string {
				return cases.Title(language.English, cases.Compact).String(answer)
			}),
	)

	io.Text(fmt.Sprintf("Hello %s %s", firstname, lastname))

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

	// Simple confirmation question
	answer := qh.Ask(
		question.
			NewComfirmation("Do you like this library?").
			SetDefaultAnswer(answers.YES).
			SetMaxAttempts(2),
	)
	if answer == answers.YES {
		io.Text("Great!")
	} else {
		io.Text("... ok :(")
	}

	// Choice question with multiple answers allowed
	answer = qh.Ask(
		question.
			NewChoices("What is your favorite color?", []string{"red", "blue", "yellow"}).
			SetMultiselect(true).
			SetMaxAttempts(3),
	)

	// Retrieve all selected colors by splitting the answer on commas
	colors := strings.Split(answer, ",")

	for _, color := range colors {
		io.Text("One of your favorite color is " + color)
	}

	if len(colors) > 1 {
		// Choice question with only one answer allowed
		answer = qh.Ask(
			question.
				NewChoices("What is your overall favorite color?", colors).
				SetMaxAttempts(3),
		)

		io.Text("Your overall favorite color is " + answer)
	}
}
