package question

import (
	"errors"
	"github.com/DrSmithFr/go-console/question/normalizer"
	"github.com/DrSmithFr/go-console/question/validator"
)

type QuestionBasic struct {
	question            string
	defaultAnswer       string
	maxAttempts         int
	hidden              bool
	hiddenFallback      bool
	autocompletedValues *[]string
	validator           *validator.Validator
	normalizer          *normalizer.Normalizer
}

func NewQuestion(question string) *QuestionBasic {
	q := new(QuestionBasic)

	q.question = question

	q.hidden = false
	q.hiddenFallback = true
	q.autocompletedValues = nil
	q.validator = nil
	q.normalizer = nil

	return q
}

// Implement QuestionBasicInterface

func (q *QuestionBasic) GetQuestion() string {
	return q.question
}

func (q *QuestionBasic) GetDefaultAnswer() string {
	return q.defaultAnswer
}

func (q *QuestionBasic) IsHidden() bool {
	return q.hidden
}

func (q *QuestionBasic) IsHiddenFallback() bool {
	return q.hiddenFallback
}

func (q *QuestionBasic) GetAutocompletedValues() *[]string {
	return q.autocompletedValues
}

func (q *QuestionBasic) GetValidator() func(string) error {
	if q.validator == nil {
		return nil
	}

	return *q.validator
}

func (q *QuestionBasic) GetMaxAttempts() int {
	return q.maxAttempts
}

func (q *QuestionBasic) GetNormalizer() func(string) string {
	if q.normalizer == nil {
		return nil
	}

	return *q.normalizer
}

// Fluent setters

func (q *QuestionBasic) SetDefaultAnswer(defaultAnswer string) *QuestionBasic {
	q.defaultAnswer = defaultAnswer
	return q
}

func (q *QuestionBasic) SetHidden(hidden bool) *QuestionBasic {
	q.hidden = hidden
	return q
}

func (q *QuestionBasic) SetHiddenFallback(fallback bool) *QuestionBasic {
	q.hiddenFallback = fallback
	return q
}

func (q *QuestionBasic) SetAutocompletedValues(values *[]string) *QuestionBasic {
	q.autocompletedValues = values
	return q
}

func (q *QuestionBasic) SetValidator(validator validator.Validator) *QuestionBasic {
	q.validator = &validator
	return q
}

func (q *QuestionBasic) SetMaxAttempts(attempts int) *QuestionBasic {
	if attempts < 0 {
		panic(errors.New("maximum number of maxAttempts must be zero or a positive value"))
	}

	q.maxAttempts = attempts

	return q
}

func (q *QuestionBasic) SetNormalizer(normalizer normalizer.Normalizer) *QuestionBasic {
	q.normalizer = &normalizer
	return q
}
