package question

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/helper"
	"github.com/DrSmithFr/go-console/question/normalizer"
	"github.com/DrSmithFr/go-console/question/validator"
	"regexp"
	"strings"
)

type QuestionChoicesInterface interface {
	QuestionBasicInterface
	GetChoices() []string
	GetPrompt() string
	IsMultiselect() bool
	GetErrorMessage() string
	GetDefaultNormalizer() normalizer.Normalizer
	GetDefaultValidator() validator.Validator
}

type QuestionChoices struct {
	QuestionBasic
	choices      []string
	multiselect  bool
	prompt       string
	errorMessage string
}

func NewChoices(question string, choices []string) *QuestionChoices {
	q := new(QuestionChoices)

	q.question = question

	q.choices = choices
	q.autocompletedValues = &choices
	q.multiselect = false
	q.prompt = " > "
	q.errorMessage = "Value '%s' is invalid"

	norm := q.GetDefaultNormalizer()
	validation := q.GetDefaultValidator()

	q.normalizer = &norm
	q.validator = &validation

	return q
}

// Implement QuestionChoicesInterface

func (q *QuestionChoices) GetChoices() []string {
	return q.choices
}

func (q *QuestionChoices) IsMultiselect() bool {
	return q.multiselect
}

func (q *QuestionChoices) GetPrompt() string {
	return q.prompt
}

func (q *QuestionChoices) GetErrorMessage() string {
	return q.errorMessage
}

// Implement Custom Methods

func (q *QuestionChoices) GetDefaultNormalizer() normalizer.Normalizer {
	return func(answer string) string {
		var selectedChoices []string

		if q.IsMultiselect() {
			// remove right last comma
			answer = strings.TrimRight(answer, ",")

			// Check for a separated comma values
			matches := regexp.MustCompile(`^[^,]+(?:,[^,]+)*$`).FindStringSubmatch(answer)
			if len(matches) == 0 {
				panic(errors.New(fmt.Sprintf(q.GetErrorMessage(), answer)))
			}

			selectedChoices = helper.Map(strings.Split(answer, ","), strings.TrimSpace)
		} else {
			selectedChoices = []string{strings.TrimSpace(answer)}
		}

		return strings.Join(selectedChoices, ",")
	}
}

func (q *QuestionChoices) GetDefaultValidator() validator.Validator {
	return func(answer string) error {
		var selectedChoices []string

		if q.IsMultiselect() {
			selectedChoices = strings.Split(answer, ",")
		} else {
			selectedChoices = []string{answer}
		}

		for _, value := range selectedChoices {
			matched := false

			for _, choice := range q.GetChoices() {
				matched = matched || choice == value
			}

			if !matched {
				return errors.New(fmt.Sprintf(q.GetErrorMessage(), value))
			}
		}

		return nil
	}
}

// Implement Fluent setters for QuestionChoices

func (q *QuestionChoices) SetMultiselect(multiselect bool) *QuestionChoices {
	q.multiselect = multiselect
	return q
}

func (q *QuestionChoices) SetPrompt(prompt string) *QuestionChoices {
	q.prompt = prompt
	return q
}

func (q *QuestionChoices) SetErrorMessage(errorMessage string) *QuestionChoices {
	q.errorMessage = errorMessage
	return q
}

// Implement Fluent setters for QuestionBasic

func (q *QuestionChoices) SetDefaultAnswer(defaultAnswer string) *QuestionChoices {
	q.defaultAnswer = defaultAnswer
	return q
}

func (q *QuestionChoices) SetHidden(hidden bool) *QuestionChoices {
	q.hidden = hidden
	return q
}

func (q *QuestionChoices) SetHiddenFallback(fallback bool) *QuestionChoices {
	q.hiddenFallback = fallback
	return q
}

func (q *QuestionChoices) SetAutocompletedValues(values *[]string) *QuestionChoices {
	q.autocompletedValues = values
	return q
}

func (q *QuestionChoices) SetValidator(validator validator.Validator) *QuestionChoices {
	q.validator = &validator
	return q
}

func (q *QuestionChoices) SetMaxAttempts(attempts int) *QuestionChoices {
	if attempts < 0 {
		panic(errors.New("maximum number of maxAttempts must be zero or a positive value"))
	}

	q.maxAttempts = attempts

	return q
}

func (q *QuestionChoices) SetNormalizer(normalizer normalizer.Normalizer) *QuestionChoices {
	q.normalizer = &normalizer
	return q
}
