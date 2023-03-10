package table

import (
	"fmt"
	"github.com/DrSmithFr/go-console/helper"
	"github.com/DrSmithFr/go-console/output"
	"sort"
	"strings"
	"unicode/utf8"
)

type rowType int

const (
	rowSimple rowType = 0
	rowTop    rowType = 2
	rowDouble rowType = 1
	rowBottom rowType = 3
)

type columnType int

const (
	columnInside  columnType = 0
	columnOutside columnType = 1
)

type TableRenderInterface interface {
	GetColumnStyle(column int) TableStyleInterface
}

type TableRender struct {
	content *Table

	output output.OutputInterface
	style  TableStyleInterface

	columnsStyles map[int]TableStyleInterface

	columnsMinWidths map[int]int
	columnsMaxWidths map[int]int

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

	t.columnsMinWidths = map[int]int{}
	t.columnsMaxWidths = map[int]int{}

	t.effectiveColumnWidths = map[int]int{}

	t.SetStyleFromName("default")

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

func (t *TableRender) SetStyle(style TableStyleInterface) *TableRender {
	t.style = style
	return t
}

func (t *TableRender) SetStyleFromName(name string) *TableRender {
	t.style = GetStyleDefinition(name)
	return t
}

func (t *TableRender) SetColumnStyle(column int, name string) *TableRender {
	t.columnsStyles[column] = GetStyleDefinition(name)
	return t
}

// External width management

func (t *TableRender) SetColumnMinWidth(column int, width int) *TableRender {
	maxWidth := t.GetColumnMaxWidth(column)

	if maxWidth > 0 && width > t.GetColumnMaxWidth(column) {
		panic(fmt.Sprintf("The minimum width of a column (%d) cannot be greater than the maximum width (%d).", width, t.GetColumnMaxWidth(column)))
	}

	t.columnsMinWidths[column] = width
	return t
}

func (t *TableRender) GetColumnMinWidth(column int) int {
	return t.columnsMinWidths[column]
}

func (t *TableRender) SetColumnsMinWidths(widths map[int]int) *TableRender {
	t.columnsMinWidths = map[int]int{}

	for column, width := range widths {
		t.SetColumnMinWidth(column, width)
	}

	return t
}

func (t *TableRender) SetColumnMaxWidth(column int, width int) *TableRender {
	if width < t.GetColumnMinWidth(column) {
		panic(fmt.Sprintf("The maximum width of a column (%d) cannot be smaller than the minimum width (%d).", width, t.GetColumnMinWidth(column)))
	}

	t.columnsMaxWidths[column] = width
	return t
}

func (t *TableRender) GetColumnMaxWidth(column int) int {
	return t.columnsMaxWidths[column]
}

func (t *TableRender) SetColumnsMaxWidths(widths map[int]int) *TableRender {
	t.columnsMaxWidths = map[int]int{}

	for column, width := range widths {
		t.SetColumnMaxWidth(column, width)
	}

	return t
}

func (t *TableRender) SetColumnWidth(column int, width int) *TableRender {
	t.columnsMinWidths[column] = width
	t.columnsMaxWidths[column] = width
	return t
}

func (t *TableRender) SetColumnsWidths(widths map[int]int) *TableRender {
	t.columnsMinWidths = map[int]int{}
	t.columnsMaxWidths = map[int]int{}

	for column, width := range widths {
		t.SetColumnMinWidth(column, width)
		t.SetColumnMaxWidth(column, width)
	}

	return t
}

// Internal width management

func (t *TableRender) setEffectiveColumnWidth(column int, width int) *TableRender {
	t.effectiveColumnWidths[column] = width
	return t
}

func (t *TableRender) getEffectiveColumnWidth(column int) int {
	return t.effectiveColumnWidths[column]
}

func (t *TableRender) setEffectiveColumnsWidths(widths map[int]int) *TableRender {
	t.effectiveColumnWidths = map[int]int{}

	for column, width := range widths {
		t.SetColumnMinWidth(column, width)
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

	t.calculateColumnsWidth(mergedData)

	rows := t.content.GetRows()
	headers := t.content.GetHeaders()

	rowsData := t.buildTableRows(rows)
	headersData := t.buildTableRows(headers)

	t.renderRowTitleSeparator(t.content.GetHeaderTitle(), rowTop)

	// rendering headers
	if len(headersData.GetRows()) > 0 {
		for _, index := range headersData.GetRowsSortedKeys() {
			header := headersData.GetRow(index)
			t.renderRow(header, t.style.GetCellHeaderFormat())
		}

		if len(rowsData.GetRows()) != 0 {
			t.renderRowSeparator(rowDouble)
		} else {
			t.renderRowTitleSeparator(t.content.GetFooterTitle(), rowBottom)
		}
	}

	// rendering rows
	for _, index := range rowsData.GetRowsSortedKeys() {
		row := rowsData.GetRow(index)
		t.renderRow(row, t.style.GetCellRowFormat())
	}

	if len(rowsData.GetRows()) > 0 {
		t.renderRowTitleSeparator(t.content.GetFooterTitle(), rowBottom)
	}

	t.cleanup()
}

func (t *TableRender) renderRowTitleSeparator(title string, direction rowType) {
	if utf8.RuneCountInString(title) == 0 {
		t.renderRowSeparator(direction)
		return
	}

	count := t.numberOfColumns

	if count == 0 {
		return
	}

	paddedTitle := fmt.Sprintf(" %s ", title)

	if utf8.RuneCountInString(t.style.GetHorizontalOutsideBorderChar()) == 0 && utf8.RuneCountInString(t.style.GetCrossingChar()) == 0 {
		t.output.Writeln(paddedTitle)
		return
	}

	separator := t.getRowSeparator(direction)

	paddedTitleLength := utf8.RuneCountInString(paddedTitle)
	separatorLength := utf8.RuneCountInString(separator)
	separatorLengthCrop := separatorLength - paddedTitleLength

	separatorCropLeft := separatorLengthCrop / 2
	//separatorCropRight := separatorLengthCrop - separatorCropLeft

	var stylisedTitle string

	if direction == rowTop {
		stylisedTitle = fmt.Sprintf(t.style.GetHeaderTitleFormat(), title)
	} else if direction == rowBottom {
		stylisedTitle = fmt.Sprintf(t.style.GetFooterTitleFormat(), title)
	} else {
		stylisedTitle = paddedTitle
	}

	titleSeparator := ""
	index := 0
	for _, char := range separator {
		if index == separatorCropLeft {
			titleSeparator += stylisedTitle
			break
		}

		titleSeparator += string(char)
		index++
	}

	index = 0
	for _, char := range separator {
		if index < separatorCropLeft+paddedTitleLength {
			index++
			continue
		}

		titleSeparator += string(char)
		index++
	}

	t.output.Writeln(titleSeparator)
}

func (t *TableRender) renderRowSeparator(direction rowType) {
	separator := t.getRowSeparator(direction)

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
func (t *TableRender) getRowSeparator(direction rowType) string {

	var horizontalBorderChar string
	if direction == rowTop || direction == rowBottom || direction == rowDouble {
		horizontalBorderChar = t.style.GetHorizontalOutsideBorderChar()
	} else if direction == rowSimple {
		horizontalBorderChar = t.style.GetHorizontalInsideBorderChar()
	}

	count := t.numberOfColumns

	if count == 0 {
		return ""
	}

	if utf8.RuneCountInString(horizontalBorderChar) == 0 && utf8.RuneCountInString(t.style.GetCrossingChar()) == 0 {
		return ""
	}

	var markup string
	if direction == rowTop {
		markup = t.style.GetCrossingTopLeftChar()
	} else if direction == rowBottom {
		markup = t.style.GetCrossingBottomLeftChar()
	} else if direction == rowSimple {
		markup = t.style.GetCrossingMidLeftChar()
	} else if direction == rowDouble {
		markup = t.style.GetCrossingTopLeftBottomChar()
	}

	for column := 0; column < count; column++ {
		markup += strings.Repeat(horizontalBorderChar, t.getEffectiveColumnWidth(column))

		if column == count-1 {
			if direction == rowTop {
				markup += t.style.GetCrossingTopRightChar()
			} else if direction == rowBottom {
				markup += t.style.GetCrossingBottomRightChar()
			} else if direction == rowSimple {
				markup += t.style.GetCrossingMidRightChar()
			} else if direction == rowDouble {
				markup += t.style.GetCrossingTopRightBottomChar()
			}
		} else {
			if direction == rowTop {
				markup += t.style.GetCrossingTopMidChar()
			} else if direction == rowBottom {
				markup += t.style.GetCrossingBottomMidChar()
			} else if direction == rowSimple {
				markup += t.style.GetCrossingChar()
			} else if direction == rowDouble {
				markup += t.style.GetCrossingTopMidBottomChar()
			}
		}
	}

	return fmt.Sprintf(t.style.GetBorderFormat(), markup)
}

/**
 * Renders vertical column separator.
 */
func (t *TableRender) renderColumnSeparator(direction columnType) string {
	if direction == columnOutside {
		return fmt.Sprintf(t.style.GetBorderFormat(), t.style.GetVerticalOutsideBorderChar())
	}

	return fmt.Sprintf(t.style.GetBorderFormat(), t.style.GetVerticalInsideBorderChar())
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

	rowContent := t.renderColumnSeparator(columnOutside)

	for index := 0; index < t.numberOfColumns; {
		//for _, index := range row.GetColumnsSortedKeys() {
		column := row.GetColumn(index)

		if column == nil {
			rowContent += t.renderCell(row, index, cellFormat)

			if index == t.numberOfColumns-1 {
				rowContent += t.renderColumnSeparator(columnOutside)
			} else {
				rowContent += t.renderColumnSeparator(columnInside)
			}

			index++
			continue
		}

		cell := column.GetCell()

		if separator, ok := cell.(TableSeparatorInterface); ok {
			if separator.IsDouble() {
				rowContent = t.getRowSeparator(rowDouble)
			} else {
				rowContent = t.getRowSeparator(rowSimple)
			}
			break
		}

		rowContent += t.renderCell(row, index, cellFormat)

		if index+(cell.GetColspan()-1) == t.numberOfColumns-1 {
			rowContent += t.renderColumnSeparator(columnOutside)
		} else {
			rowContent += t.renderColumnSeparator(columnInside)
		}

		index += cell.GetColspan()
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

	width := t.getEffectiveColumnWidth(columnIndex)

	if cell.GetColspan() > 1 {
		nextColumns := helper.RangeInt(columnIndex+1, columnIndex+cell.GetColspan()-1)

		for _, nextColumn := range nextColumns {
			width += t.getColumnSeparatorWidth() + t.getEffectiveColumnWidth(nextColumn)
		}
	}

	style := t.GetColumnStyle(columnIndex)

	if _, ok := cell.(TableSeparatorInterface); ok {
		return fmt.Sprintf(style.GetBorderFormat(), strings.Repeat(style.GetHorizontalInsideBorderChar(), width))
	}

	cellPad := cell.GetPadType()

	if cellPad == PadDefault {
		cellPad = t.content.GetColumnPadding(columnIndex)
	}

	if cellPad == PadDefault {
		cellPad = style.GetPadType()
	}

	width += helper.Strlen(cell.GetValue()) - helper.StrlenWithoutDecoration(t.output.GetFormatter(), cell.GetValue())
	content := fmt.Sprintf(style.GetCellRowContentFormat(), cell.GetValue())

	hasResetOpeningTag := false
	if strings.Index(content, "<fg=default;bg=default>") != -1 {
		content = strings.Replace(content, "<fg=default;bg=default>", "", -1)
		width -= len("<fg=default;bg=default>")
		hasResetOpeningTag = true
	}

	hasResetClosingTag := false
	if strings.Index(content, "</>") != -1 {
		content = strings.Replace(content, "</>", "", -1)
		width -= len("</>")
		hasResetClosingTag = true
	}

	result := fmt.Sprintf(cellFormat, style.Pad(content, width, style.GetPaddingChar(), cellPad))

	if hasResetOpeningTag {
		result = "<fg=default;bg=default>" + result
	}

	if hasResetClosingTag {
		result = result + "</>"
	}

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

			// Managing column max width
			maxWidth := t.getEffectiveColumnWidth(columnIndex) - utf8.RuneCountInString(t.style.GetCellRowContentFormat()) + 2
			if maxWidth > 0 {

				if cell.GetColspan() > 1 {
					maxWidth = t.getEffectiveColumnWidth(columnIndex)

					for i := 1; i < cell.GetColspan(); i++ {
						maxWidth += t.getEffectiveColumnWidth(columnIndex + i)

						if columnIndex+i < t.numberOfColumns {
							maxWidth += t.getColumnSeparatorWidth()
						}
					}

					maxWidth -= utf8.RuneCountInString(t.style.GetCellRowContentFormat()) - 2
				}

				cellValue := cell.GetValue()
				cellRawValue := helper.RemoveDecoration(t.output.GetFormatter(), cellValue)

				cellRawWidth := utf8.RuneCountInString(cellRawValue)
				if cellRawWidth > maxWidth {

					var newValue string
					if cellValue == cellRawValue {
						newValue = helper.InsertNth(cellRawValue, maxWidth, '\n')
					} else {
						newRawValue := helper.InsertNth(cellRawValue, maxWidth, '\n')
						tags := t.output.GetFormatter().FindTagsInString(cellValue)
						newValue = helper.InsertTagsIgnoringNewLines(cellRawValue, newRawValue, tags)
					}

					cell.SetValue(newValue)
				}
			}

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
					newCell.
						SetColspan(cell.GetColspan()).
						SetColspan(cell.GetColspan()).
						SetPadType(cell.GetPadType())
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
	rowKeys := data.GetRowsSortedKeys()
	for _, rowKey := range rowKeys {
		row := data.GetRow(rowKey)
		tableRows.AddRow(t.fillCells(row))

		if _, ok := unmergedRows[rowKey]; ok {

			unmergedColumnsKeys := []int{}
			for unmergedColumnKey, _ := range unmergedRows[rowKey] {
				unmergedColumnsKeys = append(unmergedColumnsKeys, unmergedColumnKey)
			}
			sort.Ints(unmergedColumnsKeys)

			for _, unmergedColumnKey := range unmergedColumnsKeys {
				column := unmergedRows[rowKey][unmergedColumnKey]
				newRow := NewTableRow()
				for columnIndex, cell := range column {
					newRow.SetColumn(columnIndex, NewTableColumn().SetCell(cell))
				}
				tableRows.AddRow(newRow)
			}
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

		if cell.getRowspan() > 1 {
			nbLines := cell.getRowspan() - 1
			lines := []string{cell.GetValue()}
			if -1 != strings.Index(cell.GetValue(), "\n") {
				lines = strings.Split(strings.ReplaceAll("\n", "<fg=default;bg=default>\n</>", cell.GetValue()), "\n")

				if len(lines) > nbLines {
					nbLines = len(strings.Split(cell.GetValue(), "\n"))

					data.GetRow(line).GetColumn(columnIndex).SetCell(NewTableCell(lines[0]).SetColspan(cell.GetColspan()))
					lines = lines[1:]
				}

				// create a two dimensional array (rowspan x Colspan)
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

			if row == nil {
				continue
			}

			column := row.GetColumn(columnIndex)

			if column == nil {
				continue
			}

			cell := column.GetCell()

			if column == nil {
				continue
			}

			if _, ok := cell.(TableSeparatorInterface); ok {
				continue
			}

			if cell.GetColspan() > 1 {
				continue
			}

			cellWidth := t.getCellWidth(row, columnIndex)
			lengths = append(lengths, cellWidth)
		}

		columnWith := helper.MaxInt(lengths) + utf8.RuneCountInString(t.style.GetCellRowContentFormat()) - 2

		t.setEffectiveColumnWidth(columnIndex, columnWith)
	}
}

func (t *TableRender) getColumnSeparatorWidth() int {
	return utf8.RuneCountInString(fmt.Sprintf(t.style.GetBorderFormat(), t.style.GetVerticalInsideBorderChar()))
}

func (t *TableRender) getCellWidth(rows TableRowInterface, columnIndex int) int {
	column := rows.GetColumn(columnIndex)

	cellWidth := 0

	if column == nil {
		return t.getCellWidthOverride(columnIndex, cellWidth)
	}

	cell := column.GetCell()

	cellRawValue := helper.RemoveDecoration(t.output.GetFormatter(), cell.GetValue())
	cellWidth = utf8.RuneCountInString(cellRawValue)

	if -1 == strings.Index(cell.GetValue(), "\n") {
		return t.getCellWidthOverride(columnIndex, cellWidth)
	}

	for _, lines := range strings.Split(cellRawValue, "\n") {
		if utf8.RuneCountInString(lines) > cellWidth {
			cellWidth = utf8.RuneCountInString(lines)
		}
	}

	return t.getCellWidthOverride(columnIndex, cellWidth)
}

func (t *TableRender) getCellWidthOverride(columnIndex int, cellWidth int) int {
	maxWidth := t.GetColumnMaxWidth(columnIndex)

	if maxWidth > 0 && cellWidth > maxWidth {
		return maxWidth
	}

	minWidth := t.GetColumnMinWidth(columnIndex)

	if minWidth > 0 && cellWidth < minWidth {
		return minWidth
	}

	return cellWidth
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
