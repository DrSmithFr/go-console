package main

import "github.com/MrSmith777/go-console/pkg/output"

func main() {
	out := output.NewConsoleOutput()
	out.Writeln("Ceci est un test")
	out.Write("Ceci est un test")
}