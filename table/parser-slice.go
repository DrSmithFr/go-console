package table

import (
	"reflect"
)

type sliceParser struct {
	Config *ParserConfig
}

func (p *sliceParser) SetConfig(config *ParserConfig) {
	p.Config = config
}

var emptyStruct = struct{}{}

func (p *sliceParser) Parse(v reflect.Value, filters []RowFilter) ([]TableRowInterface, [][]string, []int) {
	headers := p.ParseHeaders(v)
	rows, nums := p.ParseRows(v, filters)
	return headers, rows, nums
}

func (p *sliceParser) ParseRows(v reflect.Value, filters []RowFilter) (rows [][]string, nums []int) {
	for i, n := 0, v.Len(); i < n; i++ {
		item := indirectValue(v.Index(i))
		if !CanAcceptRow(item, filters) {
			continue
		}

		if item.Kind() != reflect.Struct {
			// if not struct, don't search its fields, just put a row as it's.
			c, r := extractCells(i, emptyHeader, indirectValue(item), *p.Config)
			rows = append(rows, r)
			nums = append(nums, c...)
			continue
		}

		r, c := getRowFromStruct(item, *p.Config, 0)

		nums = append(nums, c...)
		rows = append(rows, r)
	}

	return
}

func (p *sliceParser) ParseHeaders(v reflect.Value) (headers []TableRowInterface) {
	tmp := make(map[reflect.Type]struct{})

	for i, n := 0, v.Len(); i < n; i++ {
		item := indirectValue(v.Index(i))

		// no filters.
		itemTyp := item.Type()
		if _, ok := tmp[itemTyp]; !ok {
			// make headers once per type.
			tmp[itemTyp] = emptyStruct
			hs := extractHeadersFromStruct(itemTyp, *p.Config, 0)

			if len(hs) == 0 {
				continue
			}

			headers = makeHeadersCellsFromStructHeaders(hs)
		}
	}

	return
}
