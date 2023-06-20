package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/table"
)

func main() {

	type Author struct {
		Name string
		Age  int
	}

	type Book struct {
		ISBN   string
		Title  string
		Author *Author `header:"inline"`
	}

	book := Book{
		ISBN:  "99921-58-10-7",
		Title: "The Divine Comedy",
		Author: &Author{
			Name: "Dante Alighieri",
			Age:  56,
		},
	}

	cmd := go_console.NewScript().Build()

	tab := table.
		NewTable().
		ParseData(book)

	render := table.
		NewRender(cmd.Output).
		SetContent(tab)

	render.Render()
}
