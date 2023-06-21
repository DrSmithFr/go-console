package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/table"
)

func main() {

	type Address struct {
		City    string
		Country string
	}

	type Author struct {
		Name    string
		Age     int64    `header:"Age,timestamp(ms|utc|RFC850)"`
		Address *Address `display:"inline"`
	}

	type Book struct {
		ISBN   string
		Title  string
		Author *Author
	}

	book := Book{
		ISBN:  "99921-58-10-7",
		Title: "The Divine Comedy",
		Author: &Author{
			Name: "Dante Alighieri",
			Age:  393529994,
			Address: &Address{
				City:    "Florence",
				Country: "Italy",
			},
		},
	}

	cmd := go_console.NewScript().Build()

	tab := table.
		NewTable().
		SetParseConfig(table.ParserConfig{
			MaxDepth: 1,
		}).
		Parse(book)

	render := table.
		NewRender(cmd.Output).
		SetContent(tab)

	render.Render()
}
