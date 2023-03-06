package table

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/helper"
	"github.com/DrSmithFr/go-console/pkg/output"
	"math"
	"strings"
	"unicode/utf8"
)

type TableRenderInterface interface {
	GetColumnStyle(column int) TableStyleInterface
}

type TableRender struct {
	content *Table

	output output.OutputInterface
	style  TableStyleInterface

	columnsStyles map[int]TableStyleInterface
	columnsWidths map[int]int

	numberOfColumns       int
	effectiveColumnWidths map[int]int
}

// Table constructor
func NewRender(output output.OutputInterface) *TableRender {
	t := new(TableRender)

	t.output = output

	if Styles == nil {
		initStyles()
	}

	t.content = NewTable()

	t.columnsStyles = map[int]TableStyleInterface{}
	t.columnsWidths = map[int]int{}

	t.effectiveColumnWidths = map[int]int{}

	t.SetStyle("default")

	return t
}

// Implement TableInterface

var _ TableRenderInterface = (*TableRender)(nil)

func (t *TableRender) GetColumnStyle(column int) TableStyleInterface {
	if t.columnsStyles[column] != nil {
		return t.columnsStyles[column]
	}

	return t.style
}

// Implement Table fluent setters

func (t *TableRender) SetStyle(name string) *TableRender {
	t.style = GetStyleDefinition(name)
	return t
}

func (t *TableRender) SetColumnStyle(column int, name string) *TableRender {
	t.columnsStyles[column] = GetStyleDefinition(name)
	return t
}

func (t *TableRender) SetColumnWidth(column int, width int) *TableRender {
	t.columnsWidths[column] = width
	return t
}

func (t *TableRender) GetColumnWidth(column int) int {
	return t.columnsWidths[column]
}

func (t *TableRender) SetColumnsWidths(widths map[int]int) *TableRender {
	t.columnsWidths = map[int]int{}

	for column, width := range widths {
		t.SetColumnWidth(column, width)
	}

	return t
}

func (t *TableRender) SetEffectiveColumnWidth(column int, width int) *TableRender {
	t.effectiveColumnWidths[column] = width
	return t
}

func (t *TableRender) GetEffectiveColumnWidth(column int) int {
	return t.effectiveColumnWidths[column]
}

func (t *TableRender) SetEffectiveColumnsWidths(widths map[int]int) *TableRender {
	t.effectiveColumnWidths = map[int]int{}

	for column, width := range widths {
		t.SetColumnWidth(column, width)
	}

	return t
}

// Add Content

func (t *TableRender) SetContent(content *Table) *TableRender {
	t.content = content
	return t
}

func (t *TableRender) GetContent() *Table {
	return t.content
}

// Table Rendering

func (t *TableRender) Render() {
	mergedData := MergeData(t.content.GetHeaders(), t.content.GetRows())

	t.calculateNumberOfColumns(mergedData)
	t.completeTableSeparator(mergedData)

	rows := t.content.GetRows()
	headers := t.content.GetHeaders()

	rowsData := t.buildTableRows(rows)
	headersData := t.buildTableRows(headers)

	t.calculateColumnsWidth(mergedData)
	t.renderRowTitleSeparator(t.content.GetHeaderTitle())

	if len(headersData.GetRows()) > 0 {
		for _, index := range headersData.GetRowsSortedKeys() {
			header := headersData.GetRow(index)
			t.renderRow(header, t.style.GetCellHeaderFormat())

			if len(rowsData.GetRows()) == 0 {
				t.renderRowTitleSeparator(t.content.GetFooterTitle())
			} else {
				t.renderRowSeparator()
			}
		}
	}

	for _, index := range rowsData.GetRowsSortedKeys() {
		row := rowsData.GetRow(index)
		t.renderRow(row, t.style.GetCellRowFormat())
	}

	if len(rowsData.GetRows()) > 0 {
		t.renderRowTitleSeparator(t.content.GetFooterTitle())
	}

	t.cleanup()
}

func (t *TableRender) renderRowTitleSeparator(title string) {
	separator := t.getRowSeparator()

	if len(separator) == 0 {
		return
	}

	paddedTitle := fmt.Sprintf(" %s ", title)

	if len(paddedTitle) >= len(separator) {
		t.output.Writeln(paddedTitle)
		return
	}

	separatorCrop := len(separator) - len(paddedTitle)

	separatorCropLeft := separatorCrop / 2
	separatorCropRight := separatorCrop - separatorCropLeft

	separatorLeft := separator[0:separatorCropLeft]
	separatorRight := separator[len(separator)-separatorCropRight:]

	t.output.Writeln(fmt.Sprintf("%s<b>%s</b>%s", separatorLeft, paddedTitle, separatorRight))
}

func (t *TableRender) renderRowSeparator() {
	separator := t.getRowSeparator()

	if len(separator) == 0 {
		return
	}

	t.output.Writeln(separator)
}

/**
 * Return horizontal separator.
 *
 * Example:
 *
 *     +-----+-----------+-------+
 */
func (t *TableRender) getRowSeparator() string {
	count := t.numberOfColumns

	if count == 0 {
		return ""
	}

	if len(t.style.GetHorizontalBorderChar()) == 0 && len(t.style.GetCrossingChar()) == 0 {
		return ""
	}

	markup := t.style.GetCrossingChar()
	for column := 0; column < count; column++ {
		markup += strings.Repeat(t.style.GetHorizontalBorderChar(), t.GetEffectiveColumnWidth(column)) + t.style.GetCrossingChar()
	}

	return fmt.Sprintf(t.style.GetBorderFormat(), markup)
}

/**
 * Renders vertical column separator.
 */
func (t *TableRender) renderColumnSeparator() string {
	return fmt.Sprintf(t.style.GetBorderFormat(), t.style.GetVerticalBorderChar())
}

/**
 * Renders table row.
 *
 * Example:
 *
 *     | 9971-5-0210-0 | A Tale of Two Cities  | Charles Dickens  |
 */
func (t *TableRender) renderRow(row TableRowInterface, cellFormat string) {
	if len(row.GetColumns()) == 0 {
		return
	}

	rowContent := t.renderColumnSeparator()

	for _, index := range row.GetColumnsSortedKeys() {
		column := row.GetColumn(index)
		cell := column.GetCell()

		if _, ok := cell.(TableSeparatorInterface); ok {
			rowContent = t.getRowSeparator()
			break
		}

		rowContent += t.renderCell(row, index, cellFormat)
		rowContent += t.renderColumnSeparator()
	}

	t.output.Writeln(rowContent)
}

/**
 * Renders table Cell with padding.
 */
func (t *TableRender) renderCell(row TableRowInterface, columnIndex int, cellFormat string) string {
	var cell TableCellInterface

	column := row.GetColumn(columnIndex)
	if column == nil {
		cell = NewTableCell("")
	} else {
		cell = column.GetCell()
	}

	width := t.GetEffectiveColumnWidth(columnIndex)

	if cell.GetColspan() > 1 {
		nextColumns := helper.RangeInt(columnIndex+1, columnIndex+cell.GetColspan()-1)

		for _, nextColumn := range nextColumns {
			width += t.getColumnSeparatorWidth() + t.GetEffectiveColumnWidth(nextColumn)
		}
	}

	// str_pad won't work properly with multi-byte strings, we need to fix the padding
	if utf8.ValidString(cell.GetValue()) {
		width += len(cell.GetValue()) - helper.Strlen(cell.GetValue())
	}

	style := t.GetColumnStyle(columnIndex)

	if _, ok := cell.(TableSeparatorInterface); ok {
		return fmt.Sprintf(style.GetBorderFormat(), strings.Repeat(style.GetHorizontalBorderChar(), width))
	}

	width += helper.Strlen(cell.GetValue()) - helper.StrlenWithoutDecoration(t.output.GetFormatter(), cell.GetValue())
	content := fmt.Sprintf(style.GetCellRowContentFormat(), cell.GetValue())

	result := fmt.Sprintf(cellFormat, style.Pad(content, width, style.GetPaddingChar(), style.GetPadType()))

	return result
}

func (t *TableRender) calculateNumberOfColumns(data *TableData) {
	if t.numberOfColumns != 0 {
		return
	}

	columns := []int{0}
	for _, row := range data.GetRowsAsList() {
		if _, ok := row.(TableSeparatorInterface); ok {
			continue
		}

		columns = append(columns, t.getNumberOfColumns(row))
	}

	t.numberOfColumns = helper.MaxInt(columns)
}

func (t *TableRender) completeTableSeparator(data *TableData) {
	for columnIndex := 1; columnIndex < t.numberOfColumns; columnIndex++ {
		for _, rowKey := range data.GetRowsSortedKeys() {
			row := data.GetRow(rowKey)
			firstColumn := row.GetColumn(0)

			if firstColumn == nil {
				continue
			}

			if _, ok := firstColumn.GetCell().(TableSeparatorInterface); ok {
				separatorColumn := NewTableColumn().SetCell(NewTableSeparator())
				row.SetColumn(columnIndex, separatorColumn)
			}
		}
	}
}

// TODO: check
func (t *TableRender) buildTableRows(data *TableData) *TableData {
	unmergedRows := map[int]map[int]map[int]TableCellInterface{}

	for _, rowKey := range data.GetRowsSortedKeys() {
		rows := t.fillNextRows(*data, rowKey)

		// Remove any new line breaks and replace it with a new line
		for _, columnIndex := range data.rows[rowKey].GetColumnsSortedKeys() {
			column := data.rows[rowKey].GetColumn(columnIndex)
			cell := column.GetCell()

			if -1 == strings.Index(cell.GetValue(), "\n") {
				continue
			}

			lines := strings.Split(
				strings.ReplaceAll(cell.GetValue(), "\n", "<fg=default;bg=default>\n</>"),
				"\n",
			)

			for lineKey, line := range lines {
				newCell := NewTableCell(line)

				if _, ok := cell.(TableSeparatorInterface); !ok {
					newCell = NewTableCell(line).SetColspan(t.numberOfColumns)
				}
				if 0 == lineKey {
					rows.GetRow(rowKey).GetColumn(columnIndex).SetCell(newCell)
				} else {
					if _, ok := unmergedRows[rowKey]; !ok {
						unmergedRows[rowKey] = map[int]map[int]TableCellInterface{}
					}

					if _, ok := unmergedRows[rowKey][lineKey]; !ok {
						unmergedRows[rowKey][lineKey] = map[int]TableCellInterface{}
					}

					unmergedRows[rowKey][lineKey][columnIndex] = newCell
				}
			}
		}
	}

	tableRows := NewTableData()
	for _, rowKey := range data.GetRowsSortedKeys() {
		row := data.GetRow(rowKey)
		tableRows.AddRow(t.fillCells(row))
		if _, ok := unmergedRows[rowKey]; ok {
			newRow := NewTableRow()

			for _, column := range unmergedRows[rowKey] {
				for columnIndex, cell := range column {
					newRow.SetColumn(columnIndex, NewTableColumn().SetCell(cell))
				}
			}

			tableRows.AddRow(newRow)
		}
	}

	return tableRows
}

func (t *TableRender) fillNextRows(data TableData, line int) TableData {
	unmergedRows := map[int]map[int]TableCellInterface{}

	row := data.GetRow(line)
	for _, columnIndex := range row.GetColumnsSortedKeys() {
		column := row.GetColumn(columnIndex)
		cell := column.GetCell()

		if _, ok := cell.(TableSeparatorInterface); ok {
			continue
		}

		if cell.GetRowspan() > 1 {
			nbLines := cell.GetRowspan() - 1
			lines := []string{cell.GetValue()}
			if -1 != strings.Index(cell.GetValue(), "\n") {
				lines = strings.Split(strings.ReplaceAll("\n", "<fg=default;bg=default>\n</>", cell.GetValue()), "\n")

				if len(lines) > nbLines {
					nbLines = len(strings.Split(cell.GetValue(), "\n"))

					data.GetRow(line).GetColumn(columnIndex).SetCell(NewTableCell(lines[0]).SetColspan(cell.GetColspan()))
					lines = lines[1:]
				}

				// create a two dimensional array (Rowspan x Colspan)
				filler := RowMapFill(line+1, nbLines, NewTableRow())
				unmergedRows = RowMapReplaceRecursive(filler, unmergedRows)

				for unmergedRowKey := range unmergedRows {
					value := ""

					if lines[unmergedRowKey-line] != "" {
						value = lines[unmergedRowKey-line]
					}

					unmergedRows[unmergedRowKey][columnIndex] = NewTableCell(value).SetColspan(cell.GetColspan())

					if nbLines == unmergedRowKey-line {
						break
					}
				}
			}
		}
	}

	for unmergedRowKey, unmergedRow := range unmergedRows {
		// we need to know if $unmergedRow will be merged or inserted into $rows
		row := data.GetRow(unmergedRowKey)

		if row != nil && row.GetColumns() != nil && (t.getNumberOfColumns(row)+t.getNumberOfColumns(row) <= t.numberOfColumns) {
			for cellKey, cell := range unmergedRow {
				// insert Cell into row at cellKey position
				for columnIndex, cell := range MapCellSplice(unmergedRow, cellKey, cell) {
					data.GetRow(unmergedRowKey).SetColumn(columnIndex, NewTableColumn().SetCell(cell))
				}
			}
		} else {
			row = t.copyRow(data, unmergedRowKey-1)
			for columnIndex, cell := range unmergedRow {
				if cell != nil {
					row.SetColumn(columnIndex, NewTableColumn().SetCell(cell))
				}
			}
			// array_splice($rows, $unmergedRowKey, 0, [$row]);
			data.SetRows(MapRowSplice(data.rows, unmergedRowKey, row))
		}

	}

	return data
}

/**
 * fill cells for a row that contains colspan > 1.
 */
func (t *TableRender) fillCells(row TableRowInterface) TableRowInterface {
	newRow := NewTableRow()

	for _, columnIndex := range row.GetColumnsSortedKeys() {
		column := row.GetColumn(columnIndex)
		cell := column.GetCell()

		newRow.AddColumn(NewTableColumn().SetCell(cell))

		// TODO: Find why empty cells keep being inserted
		//if _, ok := cell.(TableSeparatorInterface); !ok && cell.GetColspan() > 1 {
		//	positions := helper.RangeInt(columnIndex+1, columnIndex+cell.GetColspan()-1)
		//	for _, position := range positions {
		//		newRow.SetColumn(position, NewTableColumn().SetCell(NewTableCell("")))
		//	}
		//}
	}

	if len(newRow.GetColumns()) > 0 {
		return newRow
	}

	return row
}

func (t *TableRender) copyRow(rows TableData, line int) TableRowInterface {
	row := rows.GetRow(line)

	for _, columnIndex := range row.GetColumnsSortedKeys() {
		column := row.GetColumn(columnIndex)
		row.GetColumn(columnIndex).SetCell(NewTableCell(""))

		if _, ok := column.(TableSeparatorInterface); !ok {
			row.GetColumn(columnIndex).SetCell(NewTableCell("").SetColspan(column.GetCell().GetColspan()))
		}
	}

	return row
}

func (t *TableRender) getNumberOfColumns(row TableRowInterface) int {
	columns := len(row.GetColumns())

	for _, column := range row.GetColumns() {
		if _, ok := column.(TableSeparatorInterface); !ok {
			columns += column.GetCell().GetColspan() - 1
		}
	}

	return columns
}

func (t *TableRender) getRowColumns(row []TableCellInterface) []int {
	columns := helper.RangeInt(0, t.numberOfColumns-1)

	for cellKey, cell := range row {
		if _, ok := cell.(TableCellInterface); ok && cell.GetColspan() > 1 {
			columns = helper.ArrayDiffInt(columns, helper.RangeInt(cellKey+1, cellKey+cell.GetColspan()-1))
		}
	}

	return columns
}

func (t *TableRender) calculateColumnsWidth(data *TableData) {
	for columnIndex := 0; columnIndex < t.numberOfColumns; columnIndex++ {
		lengths := []int{}

		for _, rowKey := range data.GetRowsSortedKeys() {
			row := data.GetRow(rowKey)

			for _, i := range row.GetColumnsSortedKeys() {
				column := row.GetColumn(i)
				cell := column.GetCell()

				if _, ok := cell.(TableSeparatorInterface); ok {
					continue
				}

				textContent := helper.RemoveDecoration(t.output.GetFormatter(), cell.GetValue())
				textLenght := helper.Strlen(textContent)

				if textLenght > 0 {
					contentColumns := helper.StrSplit(textContent, int(math.Ceil(float64(textLenght)/float64(cell.GetColspan()))))

					for position, content := range contentColumns {
						row.SetColumn(i+position, MakeColumnFromString(content))
					}
				}
			}

			lengths = append(lengths, t.getCellWidth(row, columnIndex))
		}

		t.SetEffectiveColumnWidth(columnIndex, helper.MaxInt(lengths)+helper.Strlen(t.style.GetCellRowContentFormat())-2)
	}
}

func (t *TableRender) getColumnSeparatorWidth() int {
	return helper.Strlen(fmt.Sprintf(t.style.GetBorderFormat(), t.style.GetVerticalBorderChar()))
}

func (t *TableRender) getCellWidth(rows TableRowInterface, columnIndex int) int {
	cellWidth := 0

	column := rows.GetColumn(columnIndex)
	if column != nil {
		cell := column.GetCell()
		cellWidth = helper.StrlenWithoutDecoration(t.output.GetFormatter(), cell.GetValue())
	}

	if cellWidth > t.GetColumnWidth(columnIndex) {
		return cellWidth
	}

	return t.GetColumnWidth(columnIndex)
}

func (t *TableRender) cleanup() {
	t.effectiveColumnWidths = map[int]int{}
	t.numberOfColumns = 0
}

// SubInternal methods

func (t *TableRender) getAllCells() []TableColumnInterface {
	return t.content.GetColumnsAsList()
}

func (t *TableRender) getAllCellsAsList() []TableCellInterface {
	return t.content.GetCellsAsList()
}
