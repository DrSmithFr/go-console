package table

func makeHeadersCellsFromStructHeaders(hs []StructHeader) (headers []TableRowInterface) {
	length, depth := headersLengthAndDepth(hs)

	for d := 0; d <= depth; d++ {
		headers = append(headers, NewTableRow())

		for l := 0; l <= length; {
			h, _ := getHeaderAt(hs, d, l)

			padding := PadDefault

			if h.ColSpan > 1 {
				padding = PadToCenter
			}

			headers[d].SetColumn(l, &TableColumn{
				Cell: &TableCell{
					Value:   h.Name,
					Colspan: h.ColSpan,
					PadType: padding,
				},
			})

			if h.ColSpan == 0 {
				l++
			} else {
				l += h.ColSpan
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

func getHeaderAt(headers []StructHeader, depth int, length int) (header StructHeader, found bool) {
	for _, h := range headers {
		if h.Depth == depth && h.Position == length {
			return h, true
		}
	}

	return emptyHeader, false
}
