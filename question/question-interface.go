package question

type QuestionBasicInterface interface {
	// GetQuestion Returns the question
	GetQuestion() string

	// GetDefaultAnswer returns the default answer
	GetDefaultAnswer() string

	// IsHidden returns whether the user response must be hidden.
	IsHidden() bool

	// IsHiddenFallback Returns whether to fallback on non-hidden question if the response can not be hidden.
	IsHiddenFallback() bool

	// GetAutocompletedValues returns values for the autocompletion.
	GetAutocompletedValues() *[]string

	// GetValidator returns the validator for the question.
	GetValidator() func(string) error

	// GetMaxAttempts returns the maximum number of times to ask before giving up.
	GetMaxAttempts() int

	// GetNormalizer returns the normalizer for the question.
	GetNormalizer() func(string) string
}
