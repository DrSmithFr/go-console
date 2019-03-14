package formatter

type OutputFormatterStyleInterface interface {
	// Sets style foreground color.
	SetForeground(color *string)

	// Sets style background color.
	SetBackground(color *string)

	//Sets some specific style option.
	SetOption(option string)

	//Unsets some specific style option.
	UnsetOption(option string)

	// Sets multiple style options at once.
	SetOptions(options []string)

	// Applies the style to a given text.
	Apply(text string) string
}
