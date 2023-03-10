package main

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/question"
	"github.com/DrSmithFr/go-console/question/answers"
	"github.com/DrSmithFr/go-console/question/normalizer"
	"github.com/DrSmithFr/go-console/question/validator"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"regexp"
	"strings"
)

func main() {
	io := go_console.NewCli().Build()
	qh := question.NewHelper(os.Stdin, io.Output())

	// Simple question with default answer
	firstname := qh.Ask(
		question.NewQuestion("What is your firstname?").
			SetDefaultAnswer("Doe"),
	)
	io.PrintText("Hello " + firstname)

	// Simple question with normalizer
	lastname := qh.Ask(
		question.NewQuestion("What is your lastname?").
			SetNormalizer(func(answer string) string {
				return cases.Title(language.English, cases.Compact).String(answer)
			}),
	)

	io.PrintText(fmt.Sprintf("Hello %s %s", firstname, lastname))

	// Simple question with custom validator
	nickname := qh.Ask(
		question.NewQuestion("What is your nickname?").
			SetValidator(func(answer string) error {
				regex := regexp.MustCompile("^(\\w|_|-)*$")
				if match := regex.MatchString(answer); !match {
					return errors.New("nickname must be alphanumeric")
				}

				return nil
			}),
	)
	io.PrintText("Hi " + nickname)

	// Simple question with hidden input
	password := qh.Ask(
		question.NewQuestion("What is your password?").
			SetHidden(true),
	)
	io.PrintText("Password: " + password)

	// Simple confirmation question
	answer := qh.Ask(
		question.NewComfirmation("Do you like this library?").
			SetDefaultAnswer(answers.Yes).
			SetMaxAttempts(2),
	)
	if answer == answers.Yes {
		io.PrintText("Great!")
	} else {
		io.PrintText("... ok :(")
	}

	// Choice question with multiple answers allowed
	answer = qh.Ask(
		question.NewChoices("What is your favorite color?", []string{"red", "blue", "yellow"}).
			SetMultiselect(true).
			SetMaxAttempts(3),
	)

	// Retrieve all selected colors by splitting the answer on commas
	colors := strings.Split(answer, ",")

	for _, color := range colors {
		io.PrintText("One of your favorite color is " + color)
	}

	if len(colors) > 1 {
		// Choice question with only one answer allowed
		answer = qh.Ask(
			question.NewChoices("What is your overall favorite color?", colors).
				SetMaxAttempts(3),
		)

		io.PrintText("Your overall favorite color is " + answer)
	}

	// simple chain normalizer example
	answer = qh.Ask(
		question.NewChoices("What is your favorite color?", []string{"red", "blue", "yellow"}).
			SetMultiselect(true).
			SetMaxAttempts(3).
			SetNormalizer(
				normalizer.MakeChainedNormalizer(
					normalizer.Ucfirst,
					strings.ToLower,
				),
			),
	)
	io.PrintText(answer)

	// chain normalizer example using including the default normalizer
	q1 := question.NewChoices("What is your favorite color?", []string{"red", "blue", "yellow"}).
		SetMultiselect(true).
		SetMaxAttempts(3)

	customNormalizer := normalizer.MakeChainedNormalizer(
		strings.ToLower,
		q1.GetDefaultNormalizer(),
	)

	data := qh.Ask(
		q1.SetNormalizer(customNormalizer),
	)
	io.PrintText(data)

	// chain validator example
	answer = qh.Ask(
		question.NewQuestion("What is your favorite color?").
			SetValidator(
				validator.MakeChainedValidator(
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
	io.PrintText(answer)
}
