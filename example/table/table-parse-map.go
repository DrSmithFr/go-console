package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/table"
)

type Book struct {
	ISBN  string
	Title string
}

func (b Book) String() string {
	return b.ISBN
}

func main() {
	books := map[string][]Book{
		"bookshelves 1": {
			{ISBN: "99921-58-10-7", Title: "The Divine Comedy"},
			{ISBN: "9971-5-0210-0", Title: "A Tale of Two Cities"},
		},
		"bookshelves 2": {
			{ISBN: "960-425-059-0", Title: "The Lord of the Rings"},
			{ISBN: "80-902734-1-6", Title: "And Then There Were None"},
		},
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
