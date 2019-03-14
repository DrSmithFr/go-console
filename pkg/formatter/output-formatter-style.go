package formatter

type OutputFormatterStyle struct {

}

// Sets style foreground color.
func (style *OutputFormatterStyle) SetForeground(color *string) {

}

// Sets style background color.
func (style *OutputFormatterStyle) SetBackground(color *string) {

}

//Sets some specific style option.
func (style *OutputFormatterStyle) SetOption(option string) {

}

//Unsets some specific style option.
func (style *OutputFormatterStyle) UnsetOption(option string) {

}

// Sets multiple style options at once.
func (style *OutputFormatterStyle) SetOptions(options []string) {

}

// Applies the style to a given text.
func (style *OutputFormatterStyle) Apply(text string) string {
	return text
}