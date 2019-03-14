package main

import (
	"fmt"
	"github.com/MrSmith777/go-console/pkg/output"
)

func main() {
	channel := make(chan string, 1)

	out := output.NewChanOutput(channel)
	out.Writeln("Ceci est un test")

	msg := <-channel
	fmt.Print(msg)
}