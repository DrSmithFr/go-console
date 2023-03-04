package question

import (
	"errors"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/question/answers"
	"github.com/DrSmithFr/go-console/pkg/question/normalizer"
	"github.com/DrSmithFr/go-console/pkg/question/validator"
	"regexp"
)

type QuestionConfirmationInterface interface {
	QuestionBasicInterface
	GetYesRegex() *regexp.Regexp
	GetNoRegex() *regexp.Regexp
	GetErrorMessage() string
	GetDefaultNormalizer() normalizer.Normalizer
	GetDefaultValidator() validator.Validator
}

type QuestionConfirmation struct {
	QuestionBasic
	yesRegex     *regexp.Regexp
	noRegex      *regexp.Regexp
	errorMessage string
}

func NewComfirmation(question string) *QuestionConfirmation {
	q := &QuestionConfirmation{
		yesRegex: regexp.MustCompile("^(y|yes|true|1)$"),
		noRegex:  regexp.MustCompile("^(n|no|false|0)$"),
	}

	q.question = question

	q.defaultAnswer = answers.NONE
	q.errorMessage = "Value '%s' is invalid. yes or no is expected."

	norm := q.GetDefaultNormalizer()
	validation := q.GetDefaultValidator()

	q.normalizer = &norm
	q.validator = &validation

	return q
}

// Implement QuestionConfirmationInterface

func (q *QuestionConfirmation) GetYesRegex() *regexp.Regexp {
	return q.yesRegex
}

func (q *QuestionConfirmation) GetNoRegex() *regexp.Regexp {
	return q.noRegex
}

func (q *QuestionConfirmation) GetErrorMessage() string {
	return q.errorMessage
}

// Implement Custom Methods

func (q *QuestionConfirmation) GetDefaultNormalizer() normalizer.Normalizer {
	return func(answer string) string {
		if answer == "" {
			return q.GetDefaultAnswer()
		}

		if q.GetYesRegex().MatchString(answer) {
			return answers.YES
		}

		if q.GetNoRegex().MatchString(answer) {
			return answers.NO
		}

		return answer
	}
}

func (q *QuestionConfirmation) GetDefaultValidator() validator.Validator {
	return func(answer string) error {
		valid := q.GetYesRegex().MatchString(answer) || q.GetNoRegex().MatchString(answer)

		if !valid {
			return errors.New(fmt.Sprintf(q.GetErrorMessage(), answer))
		}

		return nil
	}
}

// Implement Fluent setters for QuestionConfirmation

func (q *QuestionConfirmation) SetYesRegex(regex *regexp.Regexp) *QuestionConfirmation {
	q.yesRegex = regex
	return q
}

func (q *QuestionConfirmation) SetNoRegex(regex *regexp.Regexp) *QuestionConfirmation {
	q.noRegex = regex
	return q
}

// Implement Fluent setters for QuestionBasic

func (q *QuestionConfirmation) SetDefaultAnswer(defaultAnswer string) *QuestionConfirmation {
	q.defaultAnswer = defaultAnswer
	return q
}

func (q *QuestionConfirmation) SetHidden(hidden bool) *QuestionConfirmation {
	q.hidden = hidden
	return q
}

func (q *QuestionConfirmation) SetHiddenFallback(fallback bool) *QuestionConfirmation {
	q.hiddenFallback = fallback
	return q
}

func (q *QuestionConfirmation) SetAutocompletedValues(values *[]string) *QuestionConfirmation {
	q.autocompletedValues = values
	return q
}

func (q *QuestionConfirmation) SetValidator(validator validator.Validator) *QuestionConfirmation {
	q.validator = &validator
	return q
}

func (q *QuestionConfirmation) SetMaxAttempts(attempts int) *QuestionConfirmation {
	if attempts < 0 {
		panic(errors.New("maximum number of maxAttempts must be zero or a positive value"))
	}

	q.maxAttempts = attempts

	return q
}

func (q *QuestionConfirmation) SetNormalizer(normalizer normalizer.Normalizer) *QuestionConfirmation {
	q.normalizer = &normalizer
	return q
}
