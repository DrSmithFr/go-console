package main

import (
	"github.com/DrSmithFr/go-console/pkg/style"
	"github.com/DrSmithFr/go-console/pkg/table"
)

func main() {
	io := style.NewConsoleCommand().Build()
	out := io.GetOutput()

	tab := table.
		NewTable().
		SetHeadersFromString(
			[][]string{
				{"ISBN", "Title", "Author"},
			},
		).
		SetRowsFromString(
			[][]string{
				{"99921-58-10-7", "Divine Comedy", "Dante Alighieri"},
				{"9971-5-0210-0", "A Tale of Two Cities", "Charles Dickens"},
				{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
			},
		)

	render := table.
		NewRender(out).
		SetContent(tab)

	render.Render()
}
