package helper

import (
	"bytes"
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/formatter"
	"regexp"
	"unicode"
	"unicode/utf8"
)

func Strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// length of an undecorated string
func StrlenWithoutDecoration(outputFormatter formatter.OutputFormatterInterface, message string) int {
	return utf8.RuneCountInString(RemoveDecoration(outputFormatter, message))
}

// remove all string decoration (tags)
func RemoveDecoration(outputFormatter formatter.OutputFormatterInterface, message string) string {
	wasDecorated := outputFormatter.IsDecorated()
	outputFormatter.SetDecorated(false)

	// remove <...> formatting
	noTag := outputFormatter.Format(message)

	// remove already formatted characters
	regex := regexp.MustCompile("\\033\\[[^m]*m/")
	noDecoration := regex.ReplaceAllString(noTag, "")

	outputFormatter.SetDecorated(wasDecorated)

	return noDecoration
}

// unshift a string from array
func ArrayUnshift(s []string, elements ...string) []string {
	return append(elements, s...)
}

// TODO this func as be stolen for a random repo (yep im an horrible person). Extremely disgusting, Need refactoring
func Wordwrap(message string, width int, breaker rune) string {
	// Initialize a buffer with a slightly larger size to account for breaks
	init := make([]byte, 0, len(message))
	buf := bytes.NewBuffer(init)

	var current int
	var wordBuf, spaceBuf bytes.Buffer

	for _, char := range message {
		if char == '\n' {
			if wordBuf.Len() == 0 {
				if current+spaceBuf.Len() > width {
					current = 0
				} else {
					current += spaceBuf.Len()
					_, _ = spaceBuf.WriteTo(buf)
				}
				spaceBuf.Reset()
			} else {
				current += spaceBuf.Len() + wordBuf.Len()
				_, _ = spaceBuf.WriteTo(buf)
				spaceBuf.Reset()

				_, _ = wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}
			buf.WriteRune(char)
			current = 0
		} else if unicode.IsSpace(char) {
			if spaceBuf.Len() == 0 || wordBuf.Len() > 0 {
				current += spaceBuf.Len() + wordBuf.Len()
				_, _ = spaceBuf.WriteTo(buf)
				spaceBuf.Reset()

				_, _ = wordBuf.WriteTo(buf)
				wordBuf.Reset()
			}

			spaceBuf.WriteRune(char)
		} else {

			wordBuf.WriteRune(char)

			if current+spaceBuf.Len()+wordBuf.Len() > width && wordBuf.Len() < width {
				buf.WriteRune(breaker)
				current = 0
				spaceBuf.Reset()
			}
		}
	}

	if wordBuf.Len() == 0 {
		if current+spaceBuf.Len() <= width {
			_, _ = spaceBuf.WriteTo(buf)
		}
	} else {
		_, _ = spaceBuf.WriteTo(buf)
		_, _ = wordBuf.WriteTo(buf)
	}

	return buf.String()
}

// compare byte by byte a string to another
func IsStringSliceEqual(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// equivalent of php implode()
func Implode(glue string, values []string) string {
	result := ""

	for _, value := range values {
		result = fmt.Sprintf("%s%s%s", result, value, glue)
	}

	return result
}

func ArrayDiffInt(array1 []int, arrayOthers ...[]int) []int {
	c := make(map[int]bool)
	for i := 0; i < len(array1); i++ {
		if _, hasKey := c[array1[i]]; hasKey {
			c[array1[i]] = true
		} else {
			c[array1[i]] = false
		}
	}
	for i := 0; i < len(arrayOthers); i++ {
		for j := 0; j < len(arrayOthers[i]); j++ {
			if _, hasKey := c[arrayOthers[i][j]]; hasKey {
				c[arrayOthers[i][j]] = true
			} else {
				c[arrayOthers[i][j]] = false
			}
		}
	}
	result := make([]int, 0)
	for k, v := range c {
		if !v {
			result = append(result, k)
		}
	}
	return result
}

func RangeInt(start int, end int) []int {
	vals := []int{}

	for i := start; i <= end; i++ {
		vals = append(vals, i)
	}

	return vals
}

func MaxInt(list []int) int {
	max := list[0]

	for _, val := range list {
		if val > max {
			max = val
		}
	}

	return max
}

func StrSplit(data string, length int) []string {
	if length < 0 {
		panic("length must be positive")
	} else if length == 0 {
		length = 1
	}

	result := []string{}

	for i := 0; ; i++ {
		if (i+1)*length > len(data) {
			last := data[i*length:]

			if len(last) > 0 {
				result = append(result, last)
			}

			break
		}

		result = append(result, data[i*length:(i+1)*length])
	}

	return result
}

func InsertNth(s string, n int, insert rune) string {
	var buffer bytes.Buffer

	var precedent = n - 1
	var last = len(s) - 1

	for i, r := range s {
		buffer.WriteRune(r)
		if i%n == precedent && i != last {
			buffer.WriteRune(insert)
		}
	}

	return buffer.String()
}
