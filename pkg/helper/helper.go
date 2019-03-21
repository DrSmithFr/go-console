package helper

import (
	"github.com/DrSmithFr/go-console/pkg/formatter"
	"regexp"
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
