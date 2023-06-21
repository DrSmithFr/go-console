package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/table"
)

func main() {

	type Author struct {
		Name string
		age  int
	}

	type Book struct {
		ISBN   *string
		Title  string
		Author *Author
	}

	ptrStr := func(str string) *string {
		return &str
	}

	book := Book{
		ISBN:  ptrStr("99921-58-10-7"),
		Title: "The Divine Comedy",
		Author: &Author{
			Name: "Dante Alighieri",
			age:  56,
		},
	}

	cmd := go_console.NewScript().Build()

	tab := table.
		NewTable().
		SetParseConfig(table.ParserConfig{
			TagsFieldsOnly:   false,
			UnexportedFields: true,
		}).
		Parse(book)

	render := table.
		NewRender(cmd.Output).
		SetContent(tab)

	render.Render()
}
