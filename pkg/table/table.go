package table

type TableInterface interface {
	SetHeaders(data *TableData) *Table
	GetHeaders() *TableData

	SetRows(data *TableData) *Table
	GetRows() *TableData

	GetLinesAsList() []TableRowInterface
	GetColumnsAsList() []TableColumnInterface
	GetCellsAsList() []TableCellInterface
}

type Table struct {
	headers *TableData
	rows    *TableData
}

// Table constructors

func NewTable() *Table {
	return &Table{
		headers: NewTableData(),
		rows:    NewTableData(),
	}
}

// Implements TableInterface.

var _ TableInterface = (*Table)(nil)

func (t *Table) SetHeaders(data *TableData) *Table {
	t.headers = data
	return t
}

func (t *Table) GetHeaders() *TableData {
	return t.headers
}

func (t *Table) SetRows(data *TableData) *Table {
	t.rows = data
	return t
}

func (t *Table) GetRows() *TableData {
	return t.rows
}

func (t *Table) GetLinesAsList() []TableRowInterface {
	lines := []TableRowInterface{}

	for _, line := range t.headers.GetRows() {
		lines = append(lines, line)
	}

	for _, line := range t.rows.GetRows() {
		lines = append(lines, line)
	}

	return lines
}

// Computations Helpers

func (t *Table) GetColumnsAsList() []TableColumnInterface {
	columns := []TableColumnInterface{}

	columns = append(columns, t.headers.GetColumnsAsList()...)
	columns = append(columns, t.rows.GetColumnsAsList()...)

	return columns
}

func (t *Table) GetCellsAsList() []TableCellInterface {
	cells := []TableCellInterface{}

	cells = append(cells, t.headers.GetCellsAsList()...)
	cells = append(cells, t.rows.GetCellsAsList()...)

	return cells
}

// Data injections Helpers for Headers

func (t *Table) SetHeadersFromString(rows [][]string) *Table {
	data := NewTableData()
	t.SetHeaders(data.setDataFromString(rows))
	return t
}

// Data injections Helpers for Rows

func (t *Table) SetRowsFromString(rows [][]string) *Table {
	data := NewTableData()
	t.SetRows(data.setDataFromString(rows))
	return t
}

func (t *Table) AddRows(rows []TableRowInterface) *Table {
	for _, row := range rows {
		t.AddRow(row)
	}

	return t
}

func (t *Table) AddRowsFromString(rows [][]string) *Table {
	for _, row := range rows {
		t.AddRowFromString(row)
	}

	return t
}

func (t *Table) AddRow(row TableRowInterface) *Table {
	t.rows.AddRow(row)
	return t
}

func (t *Table) AddRowFromString(row []string) *Table {
	t.rows.AddRowFromString(row)
	return t
}

func (t *Table) setRow(column int, row TableRowInterface) *Table {
	t.rows.SetRow(column, row)
	return t
}

func (t *Table) setRowFromString(column int, rowData []string) *Table {
	row := MakeRowFromStrings(rowData)
	t.rows.SetRow(column, row)
	return t
}
