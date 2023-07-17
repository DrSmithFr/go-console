package main

import (
	"fmt"
	go_console "github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/cursor"
)

func main() {
	// creating default console styler
	cmd := go_console.NewScript()
	out := cmd.Output

	c := cursor.NewCursor(out, nil)

	c.MoveToPosition(7, 11)

	// and write text on this position using the output
	out.Print("my text")

	line, col := c.GetCurrentPosition()

	fmt.Printf("line: %d, col: %d\n", line, col)
}
