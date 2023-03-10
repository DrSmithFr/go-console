package main

import "github.com/DrSmithFr/go-console"

func main() {
	// creating default console styler
	cmd := go_console.NewCli()

	// according to my terminal size (default: 120)
	cmd.SetMaxLineLength(80)

	// use simple strings for short messages
	cmd.PrintNote("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam nec nisl nec mi blandit malesuada. Nunc augue risus, posuere vitae feugiat quis, pulvinar non ligula.")

	// consider using arrays when displaying long messages
	cmd.PrintNotes([]string{
		"Lorem Ipsum Dolor Sit Amet",
		"Lorem Ipsum Dolor Sit Amet",
		"Lorem Ipsum Dolor Sit Amet",
	})
}
