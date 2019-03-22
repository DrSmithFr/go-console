package helper

import (
	"bytes"
	"github.com/DrSmithFr/go-console/pkg/formatter"
	"regexp"
	"unicode"
	"unicode/utf8"
)

func StrlenWithoutDecoration(outputFormatter formatter.OutputFormatterInterface, message string) int {
	return utf8.RuneCountInString(RemoveDecoration(outputFormatter, message))
}

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

func ArrayUnshift(s []string, elements ...string) []string {
	return append(elements, s...)
}

// TODO this func as be stolen for a random repo. Extremely disgusting, Need refactoring
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
