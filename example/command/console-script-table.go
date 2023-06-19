package main

import (
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/table"
)

func main() {
	cmd := go_console.NewScript().Build()
	out := cmd.Output

	tab := table.NewTable().
		AddHeader(
			&table.TableRow{
				Columns: map[int]table.TableColumnInterface{
					0: &table.TableColumn{
						Cell: &table.TableCell{
							Value: "ISBN",
						},
					},
					1: &table.TableColumn{
						Cell: &table.TableCell{
							Value: "Title",
						},
					},
					2: &table.TableColumn{
						Cell: &table.TableCell{
							Value:   "Author",
							Colspan: 2,
							PadType: table.PadToCenter,
						},
					},
				},
			},
		).
		AddHeadersFromString(
			[][]string{
				{"", "", "Name", "Age"},
			},
		)

	tab.
		AddRowsFromString(
			[][]string{
				{"99921-58-10-7", "The Divine Comedy", "Dante Alighieri", "56"},
				{"9971-5-0210-0", "A Tale of Two Cities", "Charles Dickens", "58"},
				{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien", "81"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie", "85"},
			},
		)

	render := table.NewRender(out).
		SetContent(tab)

	render.Render()
}
