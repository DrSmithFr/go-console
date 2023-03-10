package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console/color"
	"github.com/DrSmithFr/go-console/formatter"
)

func main() {
	// create a default style
	s1 := formatter.NewOutputFormatterStyle(color.Null, color.Null, nil)
	fmt.Printf(s1.Apply("some text without coloration\n"))

	s1.SetBackground(color.Red)
	fmt.Printf(s1.Apply("some text with red background\n"))

	s1.SetForeground(color.Green)
	fmt.Printf(s1.Apply("some text with red background and green text\n"))

	s1.SetOption(color.Bold)
	fmt.Printf(s1.Apply("some bold text with red background and green text \n"))

	// override all options in one time
	s1.SetOptions([]string{color.Underscore})
	fmt.Printf(s1.Apply("some underscore text with red background and green text \n"))

	// quick declaration
	s2 := formatter.NewOutputFormatterStyle(color.Blue, color.Yellow, nil)
	fmt.Printf(s2.Apply("some text with yellow background and blue text\n"))

	// quick declaration with options
	s3 := formatter.NewOutputFormatterStyle(color.Default, color.Default, []string{color.Underscore, color.Bold})
	fmt.Printf(s3.Apply("some bold and underscore text\n"))
}
