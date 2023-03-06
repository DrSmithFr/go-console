package table

import "sort"

type TableRowInterface interface {
	SetColumns(columns map[int]TableColumnInterface) *TableRow
	GetColumns() map[int]TableColumnInterface

	GetColumnsSortedKeys() []int

	SetColumn(index int, cell TableColumnInterface) *TableRow
	GetColumn(column int) TableColumnInterface

	GetCellsAsList() []TableCellInterface
}

type TableRow struct {
	Columns map[int]TableColumnInterface
}

// TableRow constructors

func NewTableRow() *TableRow {
	return &TableRow{
		Columns: make(map[int]TableColumnInterface),
	}
}

func MakeRowFromStrings(column []string) *TableRow {
	return NewTableRow().setRowFromString(column)
}

// Implements TableRowInterface.

var _ TableRowInterface = (*TableRow)(nil)

func (t *TableRow) SetColumns(columns map[int]TableColumnInterface) *TableRow {
	t.Columns = columns
	return t
}

func (t *TableRow) GetColumns() map[int]TableColumnInterface {
	return t.Columns
}

func (t *TableRow) GetColumnsSortedKeys() []int {
	keys := make([]int, 0, len(t.Columns))

	for k := range t.Columns {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	return keys
}

func (t *TableRow) SetColumn(index int, column TableColumnInterface) *TableRow {
	t.Columns[index] = column
	return t
}

func (t *TableRow) GetColumn(column int) TableColumnInterface {
	return t.Columns[column]
}

func (t *TableRow) GetCellsAsList() []TableCellInterface {
	cells := []TableCellInterface{}

	for _, column := range t.Columns {
		cells = append(cells, column.GetCell())
	}

	return cells
}

// Data injections Helpers

func (t *TableRow) setRowFromString(row []string) *TableRow {
	for index, column := range row {
		t.SetColumn(index, MakeColumnFromString(column))
	}

	return t
}

func (t *TableRow) AddColumn(column TableColumnInterface) *TableRow {
	addIndex := -1
	colspanOffset := 0

	for index, column := range t.Columns {
		if index > addIndex {
			addIndex = index
		}

		colspanOffset += column.GetCell().GetColspan() - 1
	}

	t.Columns[addIndex+colspanOffset+1] = column

	return t
}
