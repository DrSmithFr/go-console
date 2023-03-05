package table

import "sort"

type TableRowInterface interface {
	SetColumns(columns map[int]TableColumnInterface) *TableRow
	GetColumns() map[int]TableColumnInterface

	SortedKeys() []int

	SetColumn(index int, cell TableColumnInterface) *TableRow
	GetColumn(column int) TableColumnInterface

	GetCellsAsList() []TableCellInterface
}

type TableRow struct {
	columns map[int]TableColumnInterface
}

// TableRow constructors

func NewTableRow() *TableRow {
	return &TableRow{
		columns: make(map[int]TableColumnInterface),
	}
}

func MakeRowFromStrings(column []string) *TableRow {
	return NewTableRow().setRowFromString(column)
}

// Implements TableRowInterface.

var _ TableRowInterface = (*TableRow)(nil)

func (t *TableRow) SetColumns(columns map[int]TableColumnInterface) *TableRow {
	t.columns = columns
	return t
}

func (t *TableRow) GetColumns() map[int]TableColumnInterface {
	return t.columns
}

func (t *TableRow) SortedKeys() []int {
	keys := make([]int, 0, len(t.columns))

	for k := range t.columns {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	return keys
}

func (t *TableRow) SetColumn(index int, column TableColumnInterface) *TableRow {
	t.columns[index] = column
	return t
}

func (t *TableRow) GetColumn(column int) TableColumnInterface {
	return t.columns[column]
}

func (t *TableRow) GetCellsAsList() []TableCellInterface {
	cells := []TableCellInterface{}

	for _, column := range t.columns {
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
	addIndex := 0

	for index := range t.columns {
		if index > addIndex {
			addIndex = index
		}
	}

	t.columns[addIndex+1] = column

	return t
}
