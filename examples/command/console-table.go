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
		SetColumnPadding(3, table.PadToRight).
		AddHeader(
			&table.TableRow{
				Columns: map[int]table.TableColumnInterface{
					0: &table.TableColumn{
						Cell: &table.TableCell{
							Value:   "Centred Header Cell",
							Colspan: 3,
							PadType: table.PadToCenter,
						},
					},
				},
			},
		).
		AddHeadersFromString(
			[][]string{
				{"ISBN", "Title", "Author"},
			},
		)

	//tab.
	//	AddRowsFromString(
	//		[][]string{
	//			{"99921-58-10-7", "The Divine Comedy", "Dante \nAlighieri"},
	//			{"9971-5-0210-0", "A Tale of Two Cities", "Charles Dickens"},
	//		},
	//	).
	//	AddTableSeparator().
	//	AddRowsFromString(
	//		[][]string{
	//			{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien"},
	//			{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
	//		},
	//	)
	//
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
	//							NewTableCell("This value spans 2 columns.").
	//							SetColspan(2),
	//					),
	//			).
	//			AddColumn(
	//				table.
	//					NewTableColumn().
	//					SetCell(
	//						table.
	//							NewTableCell("stand alone value"),
	//					),
	//			),
	//	)

	tab.
		//AddTableSeparator().
		AddRow(
			&table.TableRow{
				Columns: map[int]table.TableColumnInterface{
					0: &table.TableColumn{
						Cell: &table.TableCell{
							Value:   "<question>This value spans use <b>3 lines</b> to get fully displayed.</question>",
							Colspan: 3,
						},
					},
				},
			},
		)

	tab.
		AddTableSeparator().
		AddRowsFromString(
			[][]string{
				{"<b>960-425-059-0</b>", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
			},
		)

	render := table.
		NewRender(out).
		SetContent(tab)

	//render.SetColumnsMinWidths(map[int]int{
	//	0: 10,
	//	1: 0,
	//	2: 30,
	//})

	render.SetColumnMaxWidth(0, 10)
	render.SetColumnMaxWidth(1, 10)
	render.SetColumnMaxWidth(2, 10)

	//render.SetStyle("compact")
	//render.SetStyle("borderless")

	render.Render()
}
