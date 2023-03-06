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
		SetHeaderTitle("Books").
		SetFooterTitle("Page 1/2").
		AddHeader(
			&table.TableRow{
				Columns: map[int]table.TableColumnInterface{
					0: &table.TableColumn{
						Cell: &table.TableCell{
							Value:   "Main table title",
							Colspan: 3,
						},
					},
				},
			},
		).
		AddHeadersFromString(
			[][]string{
				{"ISBN", "Title", "Author"},
			},
		).
		AddRowsFromString(
			[][]string{
				{"99921-58-10-7", "The \nDivine \nComedy", "Dante Alighieri"},
				{"9971-5-0210-0", "A Tale of Two Cities", "Charles Dickens"},
			},
		).
		AddTableSeparator().
		AddRowsFromString(
			[][]string{
				{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
			},
		)

	tab.
		AddTableSeparator().
		AddRow(
			table.
				NewTableRow().
				AddColumn(
					table.
						NewTableColumn().
						SetCell(
							table.
								NewTableCell("This value spans 2 columns.").
								SetColspan(2),
						),
				).
				AddColumn(
					table.
						NewTableColumn().
						SetCell(
							table.
								NewTableCell("stand alone value"),
						),
				),
		)

	tab.
		AddTableSeparator().
		AddRow(
			&table.TableRow{
				Columns: map[int]table.TableColumnInterface{
					0: &table.TableColumn{
						Cell: &table.TableCell{
							Value:   "This value spans use 3 lines \nto get fully displayed.",
							Colspan: 3,
						},
					},
				},
			},
		).
		AddTableSeparator().
		AddRowsFromString(
			[][]string{
				{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
			},
		)

	render := table.
		NewRender(out).
		SetContent(tab)

	//render.SetColumnsWidths(map[int]int{
	//	0: 10,
	//	1: 0,
	//	2: 30,
	//})

	render.SetColumnWidth(0, 5)
	render.SetColumnWidth(1, 10)

	//render.SetStyle("compact")
	//render.SetStyle("borderless")

	render.Render()
}
