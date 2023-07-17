package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console/terminal"
)

func main() {
	term := terminal.New()
	stty := term.HasSttyAvailable()

	if !stty {
		fmt.Println("stty is not available")
	}

	fmt.Println("stty available")

	width, height := term.GetSize()

	fmt.Printf("width: %d, height: %d\n", width, height)

	colorMode := term.GetColorMode()

	fmt.Printf("color mode: %s\n", colorMode)
}
