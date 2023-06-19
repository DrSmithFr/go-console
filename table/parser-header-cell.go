package table

func makeHeadersCellsFromStructHeaders(hs []StructHeader) (headers [][]string) {
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
