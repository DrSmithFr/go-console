package table

import (
	"reflect"
)

type sliceParser struct {
	TagsOnly bool
}

var emptyStruct = struct{}{}

func (p *sliceParser) Parse(v reflect.Value, filters []RowFilter) ([][]string, [][]string, []int) {
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
			c, r := extractCells(i, emptyHeader, indirectValue(item), p.TagsOnly)
			rows = append(rows, r)
			nums = append(nums, c...)
			continue
		}

		r, c := getRowFromStruct(item, p.TagsOnly)

		nums = append(nums, c...)
		rows = append(rows, r)
	}

	return
}

func (p *sliceParser) ParseHeaders(v reflect.Value) (headers [][]string) {
	tmp := make(map[reflect.Type]struct{})

	for i, n := 0, v.Len(); i < n; i++ {
		item := indirectValue(v.Index(i))

		// no filters.
		itemTyp := item.Type()
		if _, ok := tmp[itemTyp]; !ok {
			// make headers once per type.
			tmp[itemTyp] = emptyStruct
			hs := extractHeadersFromStruct(itemTyp, p.TagsOnly, 0)

			if len(hs) == 0 {
				continue
			}

			length, depth := headersLengthAndDepth(hs)

			for d := 0; d <= depth; d++ {
				headers = append(headers, []string{})

				for l := 0; l <= length; {
					h := getHeaderAt(hs, d, l)
					headers[d] = append(headers[d], h.Name)

					if h.ColSpan == 0 {
						l++
					} else {
						l += h.ColSpan
					}
				}
			}
		}
	}

	return
}

func headersLengthAndDepth(headers []StructHeader) (length int, depth int) {
	for _, h := range headers {
		if h.Depth > depth {
			depth = h.Depth
		}

		if h.Position > length {
			length = h.Position
		}
	}

	return
}

func getHeaderAt(headers []StructHeader, depth int, length int) (header StructHeader) {
	for _, h := range headers {
		if h.Depth == depth && h.Position == length {
			return h
		}
	}

	return emptyHeader
}
