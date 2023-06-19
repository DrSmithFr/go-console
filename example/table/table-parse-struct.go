package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/table"
)

type Author struct {
	Name string
	Age  int
}

type Book struct {
	ISBN   *string
	Title  *string
	Author *Author
}

func (b Book) String() string {
	return *b.ISBN
}

func ptrStr(str string) *string {
	return &str
}

func main() {
	cmd := go_console.NewScript().Build()

	book := Book{
		ISBN:  ptrStr("99921-58-10-7"),
		Title: ptrStr("The Divine Comedy"),
		Author: &Author{
			Name: "Dante Alighieri",
			Age:  56,
		},
	}

	tab := table.
		NewTable().
		ParseData(book)

	render := table.
		NewRender(cmd.Output).
		SetContent(tab)

	render.Render()
}
