package table

import "sort"

type TableDataInterface interface {
	SetRows(lines map[int]TableRowInterface) *TableData
	GetRows() map[int]TableRowInterface

	GetRowsSortedKeys() []int

	SetRow(index int, row TableRowInterface) *TableData
	GetRow(index int) TableRowInterface

	GetColumnsAsList() []TableColumnInterface
	GetCellsAsList() []TableCellInterface
}

type TableData struct {
	rows map[int]TableRowInterface
}

// TableData constructors

func NewTableData() *TableData {
	return &TableData{
		rows: make(map[int]TableRowInterface),
	}
}

func MakeDataFromStrings(rows [][]string) *TableData {
	return NewTableData().setDataFromString(rows)
}

func MergeData(datas ...*TableData) *TableData {
	mergedData := NewTableData()

	index := 0

	for _, data := range datas {
		for _, row := range data.GetRows() {
			mergedData.SetRow(index, row)
			index++
		}
	}

	return mergedData
}

var _ TableDataInterface = (*TableData)(nil)

// Implements TableDataInterface.

func (t *TableData) SetRows(lines map[int]TableRowInterface) *TableData {
	t.rows = lines
	return t
}

func (t *TableData) GetRows() map[int]TableRowInterface {
	return t.rows
}

func (t *TableData) GetRowsSortedKeys() []int {
	keys := []int{}

	for k := range t.rows {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	return keys
}

func (t *TableData) SetRow(index int, row TableRowInterface) *TableData {
	t.rows[index] = row
	return t
}

func (t *TableData) GetRow(index int) TableRowInterface {
	return t.rows[index]
}

func (t *TableData) GetRowsAsList() []TableRowInterface {
	rows := []TableRowInterface{}

	for _, row := range t.rows {
		rows = append(rows, row)
	}

	return rows
}

func (t *TableData) GetColumnsAsList() []TableColumnInterface {
	columns := []TableColumnInterface{}

	for _, line := range t.rows {
		for _, column := range line.GetColumns() {
			columns = append(columns, column)
		}
	}

	return columns
}

func (t *TableData) GetCellsAsList() []TableCellInterface {
	cells := []TableCellInterface{}

	for _, row := range t.rows {
		for _, cell := range row.GetCellsAsList() {
			cells = append(cells, cell)
		}
	}

	return cells
}

// Data injections Helpers

func (t *TableData) setDataFromString(rows [][]string) *TableData {
	for index, column := range rows {
		t.rows[index] = MakeRowFromStrings(column)
	}

	return t
}

func (t *TableData) AddRows(rows []TableRowInterface) *TableData {
	for _, row := range rows {
		t.AddRow(row)
	}

	return t
}

func (t *TableData) AddRowsFromString(rows [][]string) *TableData {
	for _, row := range rows {
		t.AddRowFromString(row)
	}

	return t
}

func (t *TableData) AddRow(row TableRowInterface) *TableData {
	addIndex := -1
	rowspanOffset := 0

	for index := range t.rows {
		if index > addIndex {
			addIndex = index
		}

		for _, column := range t.rows[index].GetColumnsSortedKeys() {
			column := t.rows[index].GetColumn(column)
			cell := column.GetCell()

			if _, ok := cell.(TableSeparatorInterface); ok {
				rowspanOffset++
			} else {
				rowspanOffset += column.GetCell().getRowspan() - 1
			}
		}
	}

	t.rows[addIndex+rowspanOffset+1] = row

	return t
}

func (t *TableData) AddRowFromString(rowData []string) *TableData {
	t.AddRow(MakeRowFromStrings(rowData))
	return t
}

func (t *TableData) setRow(column int, row TableRowInterface) *TableData {
	t.rows[column] = row
	return t
}

func (t *TableData) setRowFromString(column int, row []string) *TableData {
	t.rows[column] = MakeRowFromStrings(row)
	return t
}
