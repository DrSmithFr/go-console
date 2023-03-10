package table

func RowMapFill(startIndex int, count int, value TableRowInterface) TableData {
	result := NewTableData()

	for i := startIndex; i < startIndex+count; i++ {
		result.setRow(i, value)
	}

	return *result
}

func RowMapReplaceRecursive(
	base TableData,
	replacement map[int]map[int]TableCellInterface,
) map[int]map[int]TableCellInterface {

	result := map[int]map[int]TableCellInterface{}

	for baseRowIndex, baseRow := range base.GetRows() {
		for baseColumnIndex, baseColumn := range baseRow.GetColumns() {
			result[baseRowIndex][baseColumnIndex] = baseColumn.GetCell()
		}
	}

	for rowIndex, rowData := range replacement {
		for columnIndex, cell := range rowData {
			result[rowIndex][columnIndex] = cell
		}
	}

	return result
}

func MapCellSplice(data map[int]TableCellInterface, offset int, replacer TableCellInterface) map[int]TableCellInterface {
	for i := 0; i < offset; i++ {
		data[i] = replacer
	}

	return data
}

func MapRowSplice(data map[int]TableRowInterface, offset int, replacer TableRowInterface) map[int]TableRowInterface {
	for i := 0; i < offset; i++ {
		data[i] = replacer
	}

	return data
}
