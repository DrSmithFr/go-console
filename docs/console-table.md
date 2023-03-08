## Table

When building a console application it may be useful to display tabular data:

<p align="center">
    <img src="assets/table/table-example.png">
</p>

### Basic Usage

To display a table, use Table, set the headers, set the rows and then render the table:

```go
package main

import (
	"github.com/DrSmithFr/go-console/pkg/style"
	"github.com/DrSmithFr/go-console/pkg/table"
)

func main() {
	io := style.NewConsoleCommand().Build()

	tab := table.
		NewTable().
		AddHeadersFromString(
			[][]string{
				{"ISBN-LONG-TITLE", "Title", "Author"},
			},
		)

	tab.
		AddRowsFromString(
			[][]string{
				{"99921-58-10-7", "The Divine Comedy", "Dante Alighieri"},
				{"9971-5-0210-0", "A Tale of Two Cities", "Charles Dickens"},
				{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
			},
		)

	render := table.
		NewRender(io.GetOutput()).
		SetContent(tab)

	render.Render()
}
```

You can add a table separator anywhere in the output by passing an instance of TableSeparator as a row:

```go
	tab.
		AddRowsFromString(
			[][]string{
				{"99921-58-10-7", "The Divine Comedy", "Dante Alighieri"},
				{"9971-5-0210-0", "A Tale of Two Cities", "Charles Dickens"},
				{"---"}, // or "===" 
				{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
			},
		)
```

<p align="center">
    <img src="assets/table/table-separator-example.png">
</p>

You can optionally display titles at the top and the bottom of the table:

```go
	tab.
        SetHeaderTitle("Books").
        SetFooterTitle("Page 1/2")
```

<p align="center">
    <img src="assets/table/table-title.png">
</p>

By default, the width of the columns is calculated automatically based on their contents.
Use the SetColumnWidths() method to set the column widths explicitly:

```go
    // this is equivalent to the calling SetColumnsMinWidths() and SetColumnsMaxWidths() with the same values
    render.
        SetColumnsWidths(map[int]int{
            0: 10,
            1: 0,
            2: 30,
        })
    render.Render()
```

In this example, the first column width will be 10, 
the last column width will be 30 and the second column width will be calculated automatically because of the 0 value.

You can also set the width individually for each column with the SetColumnWidth() method. 
Its first argument is the column index (starting from 0) and the second argument is the column width:

```go
    render.SetColumnWidth(0, 10)
    render.SetColumnWidth(2, 10)

    render.Render()
```

The output of this command will be:

<p align="center">
    <img src="assets/table/table-column-width.png">
</p>

Note that you can also set the max and min width of a column individually:

```go
    render.SetColumnMaxWidth(0, 10)
    render.SetColumnMinWidth(1, 15)

    render.
        SetColumnsMinWidths(map[int]int{
            0: 10,
            1: 0,
            2: 30,
        })

    render.
        SetColumnsMaxWidths(map[int]int{
            0: 10,
            1: 0,
            2: 30,
        })

    render.Render()
```

### Table Styling

The table style can be changed to any built-in styles via SetStyleFromName()
    
```go
    // same as calling nothing
    render.SetStyleFromName("default")

    // changes the default style to compact
    render.SetStyleFromName("compact")
    render.Render()
```

This code results in:

<p align="center">
    <img src="assets/table/table-compact.png">
</p>

You can also set the style to `borderless`:

```go
    // changes the default style to compact
    render.SetStyleFromName("borderless")
    render.Render()
```

<p align="center">
    <img src="assets/table/table-borderless.png">
</p>

You can also set the style to `box`:

```go
    // changes the default style to compact
    render.SetStyleFromName("box")
    render.Render()
```

<p align="center">
    <img src="assets/table/table-box.png">
</p>

You can also set the style to `box-double`:

```go
    // changes the default style to compact
    render.SetStyleFromName("box-double")
    render.Render()
```

<p align="center">
    <img src="assets/table/table-box-double.png">
</p>

> **Note:** 
> 
> > Using shortcut "---" and "===" to insert a tableSeparator with style `box-double`
> > will result in simple or double line separator. 
> >On every other style, it will result in a simple line separator.

If the built-in styles do not fit your need, define your own:

```go
    customStyle := table.NewTableStyle().
        SetHorizontalOutsideBorderChar("═").
        SetHorizontalInsideBorderChar("─").
        SetVerticalOutsideBorderChar("║").
        SetVerticalInsideBorderChar("│").
        SetCrossingChars("┼", "╔", "╤", "╗", "╢", "╝", "╧", "╚", "╟", "╠", "╪", "╣")

    render.SetStyle(customStyle)
```

### Spanning Multiple Columns

To make a table cell that spans multiple columns you can use a TableCell:

```go
    package main

import (
	"github.com/DrSmithFr/go-console/pkg/style"
	"github.com/DrSmithFr/go-console/pkg/table"
)

func main() {
	io := style.NewConsoleCommand().Build()

	tab := table.
		NewTable().
		AddHeadersFromString(
			[][]string{
				{"ISBN-LONG-TITLE", "Title", "Author"},
			},
		)

	tab.
		AddRowsFromString(
			[][]string{
				{"99921-58-10-7", "The Divine Comedy", "Dante Alighieri"},
				{"9971-5-0210-0", "A Tale of Two Cities", "Charles Dickens"},
				{"---"},
				{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
				{"==="},
			},
		).
		AddRow(
			&table.TableRow{
				Columns: map[int]table.TableColumnInterface{
					0: &table.TableColumn{
						Cell: &table.TableCell{
							Value:   "<info>This value spans use <b>3 columns</b> to get fully displayed and now to long to feet inside the table.</info>",
							Colspan: 3,
							PadType: table.PadToCenter,
						},
					},
				},
			},
		)

	render := table.
		NewRender(io.GetOutput()).
		SetContent(tab)

	render.SetStyleFromName("box-double")

	render.Render()
}
```

This results in:

<p align="center">
    <img src="assets/table/table-colspan.png">
</p>

> **Note:** 
> 
> > You can create a title using a header cell that spans the entire table width.

### Padding management

You can set the padding type for each cell or column individually:

- `PadToLeft` (default)
-  `PadToCenter` 
- `PadToRight`

> **Note:**
> 
> > If you set a cell padding, the column padding will be ignored.
> 
> > If you set a column padding, the default padding (defined by style) will be ignored.

```go
    package main

import (
	"github.com/DrSmithFr/go-console/pkg/style"
	"github.com/DrSmithFr/go-console/pkg/table"
)

func main() {
	io := style.NewConsoleCommand().Build()

	tab := table.
		NewTable().
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
		AddRow(
			table.
				NewTableRow().
				AddColumn(
					table.
						NewTableColumn().
						SetCell(
							table.
								NewTableCell("This value spans 2 columns.").
								SetPadType(table.PadToCenter).
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
		).
		AddTableSeparator().
		AddRowsFromString(
			[][]string{
				{"960-425-059-0", "The Lord of the Rings", "J. R. R. Tolkien"},
				{"80-902734-1-6", "And Then There Were None", "Agatha Christie"},
			},
		)

	render := table.
		NewRender(io.GetOutput()).
		SetContent(tab)

	render.SetColumnMinWidth(2, 13)

	render.SetStyleFromName("box-double")

	render.Render()
}
```

<p align="center">
    <img src="assets/table/table-padding.png">
</p>