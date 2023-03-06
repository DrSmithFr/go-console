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
		SetHeadersFromString(
			[][]string{
				{"ISBN", "Title", "Author"},
			},
		).
		AddRowsFromString(
			[][]string{
				{"99921-58-10-7", "Divine Comedy", "Dante Alighieri"},
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

	//tab.
	//	AddTableSeparator().
	//	AddRow(
	//		table.
	//			NewTableRow().
	//			AddColumn(
	//				table.
	//					NewTableColumn().
	//					SetCell(
	//						table.
	//							NewTableCell("This value spans 3 columns.").
	//							SetColspan(3),
	//					),
	//			),
	//	)

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
