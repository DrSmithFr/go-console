package main

import "github.com/DrSmithFr/go-console/pkg/output"

func main() {
	out := output.NewNullOutput(true, nil)
	out.Writeln("Ceci est un test")
	out.Write("Ceci est un test")
}
