package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/color"
	"github.com/DrSmithFr/go-console/pkg/formatter"
)

func main() {
	// create a default style
	s1 := formatter.NewOutputFormatterStyle(color.NULL, color.NULL, nil)
	fmt.Printf(s1.Apply("some text without coloration\n"))

	s1.SetBackground(color.RED)
	fmt.Printf(s1.Apply("some text with red background\n"))

	s1.SetForeground(color.GREEN)
	fmt.Printf(s1.Apply("some text with red background and green text\n"))

	s1.SetOption(color.BOLD)
	fmt.Printf(s1.Apply("some bold text with red background and green text \n"))

	// override all options in one time
	s1.SetOptions([]string{color.UNDERSCORE})
	fmt.Printf(s1.Apply("some underscore text with red background and green text \n"))

	// quick declaration
	s2 := formatter.NewOutputFormatterStyle(color.BLUE, color.YELLOW, nil)
	fmt.Printf(s2.Apply("some text with yellow background and blue text\n"))

	// quick declaration with options
	s3 := formatter.NewOutputFormatterStyle(color.DEFAULT, color.DEFAULT, []string{color.UNDERSCORE, color.BOLD})
	fmt.Printf(s3.Apply("some bold and underscore text\n"))
}
