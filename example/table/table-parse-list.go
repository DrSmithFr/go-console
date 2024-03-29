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
		ISBN   *string
		Title  *string
		Author *Author
	}

	ptrStr := func(str string) *string {
		return &str
	}

	books := []Book{
		{ISBN: ptrStr("99921-58-10-7"), Title: ptrStr("The Divine Comedy"), Author: &Author{Name: "Dante Alighieri", Age: 56}},
		{ISBN: ptrStr("9971-5-0210-0"), Title: ptrStr("A Tale of Two Cities"), Author: &Author{Name: "Charles Dickens", Age: 58}},
		{ISBN: ptrStr("960-425-059-0"), Title: ptrStr("The Lord of the Rings"), Author: &Author{Name: "J. R. R. Tolkien", Age: 81}},
		{ISBN: ptrStr("80-902734-1-6"), Title: ptrStr("And Then There Were None"), Author: &Author{Name: "Agatha Christie", Age: 85}},
	}

	cmd := go_console.NewScript().Build()

	tab := table.
		NewTable().
		Parse(books)

	render := table.
		NewRender(cmd.Output).
		SetContent(tab)

	render.Render()
}
