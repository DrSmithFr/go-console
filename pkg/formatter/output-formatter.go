package formatter

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/color"
	"regexp"
	"strings"
)

// output formatter constructor
func NewOutputFormatter() *OutputFormatter {
	formatter := & OutputFormatter{
		stylesCache: make(map[string]OutputFormatterStyle),
		styleStack: NewOutputFormatterStyleStack(nil),
	}

	formatter.SetStyle("error", *NewOutputFormatterStyle(color.WHITE, color.RED, nil))
	formatter.SetStyle("info", *NewOutputFormatterStyle(color.GREEN, color.NULL, nil))
	formatter.SetStyle("comment", *NewOutputFormatterStyle(color.YELLOW, color.NULL, nil))
	formatter.SetStyle("question", *NewOutputFormatterStyle(color.BLACK, color.CYAN, nil))
	formatter.SetStyle("b", *NewOutputFormatterStyle(color.NULL, color.NULL, []string{color.BOLD}))
	formatter.SetStyle("u", *NewOutputFormatterStyle(color.NULL, color.NULL, []string{color.UNDERSCORE}))

	return formatter
}

// Escapes "<" special char in given text.
func Escape(message string) string {
	regex := regexp.MustCompile("([^\\\\]?)<")
	escaped := regex.ReplaceAllString(message, "$1\\<")
	final := EscapeTrailingBackslash(escaped)
	return final
}

// Escapes trailing "\" in given text.
func EscapeTrailingBackslash(message string) string {
	lastChar := message[len(message)-1:]

	if "\\" == lastChar {
		totalLenght := len(message)
		message = strings.TrimSuffix(message, "\\")
		message = fmt.Sprintf(
			"%s%s",
			message,
			strings.Repeat("\x00", totalLenght-len(message)),
		)
	}

	return message
}

// Formatter class for console output.
type OutputFormatter struct {
	decorated   bool
	styleStack  *OutputFormatterStyleStack
	stylesCache map[string]OutputFormatterStyle
}

// Sets the decorated flag.
func (o *OutputFormatter) SetDecorated(decorated bool) {
	o.decorated = decorated
}

// Gets the decorated flag.
func (o *OutputFormatter) IsDecorated() bool {
	return o.decorated
}

// Sets a new style to cache.
func (o *OutputFormatter) SetStyle(name string, style OutputFormatterStyle) {
	o.stylesCache[name] = style
}

// Gets style from cache with specified name.
func (o *OutputFormatter) GetStyle(name string) *OutputFormatterStyle {
	if style, ok := o.stylesCache[name]; ok {
		return &style
	}

	return nil
}

// Gets style stack
func (o *OutputFormatter) GetStyleStack() *OutputFormatterStyleStack {
	return o.styleStack
}

// Checks if output formatter has style in cache with specified name.
func (o *OutputFormatter) HasStyle(name string) bool {
	style := o.GetStyle(name)
	return nil != style
}

// Formats a message according to the given styles.
func (o *OutputFormatter) Format(message string) string {
	offset := 0
	output := ""

	tags := o.findTagsInString(message)

	for _, tag := range tags {
		if 0 != tag.Start && '\\' == message[tag.Start-1] {
			continue
		}

		text := message[offset:][0 : tag.Start-offset]

		// add the text up to the next tag
		output = fmt.Sprintf(
			"%s%s",
			output,
			o.applyCurrentStyle(text),
		)

		offset = tag.Start + len(tag.Text)

		if !tag.Opening && "" == tag.Style {
			// </>
			o.GetStyleStack().Pop(nil)
		} else {
			style := o.createStyleFromString(tag.Style)

			if nil == style {
				output = fmt.Sprintf(
					"%s%s",
					output,
					o.applyCurrentStyle(text),
				)
			} else if tag.Opening {
				o.GetStyleStack().Push(style)
			} else {
				o.GetStyleStack().Pop(style)
			}
		}
	}

	output = fmt.Sprintf("%s%s", output, message[offset:])

	if strings.Contains(output, "\x00") {
		output = strings.Replace(output, "\x00", "\\", -1)
		output = strings.Replace(output, "\\<", "<", -1)
	}

	result := strings.Replace(output, "\\<", "<", -1)

	return result
}

// struct to describe a color tag
type tagPos struct {
	Text    string
	Tag     string
	Style   string
	Start   int
	End     int
	Opening bool
}

// Make a tagMap from a message
func (o *OutputFormatter) findTagsInString(text string) []tagPos {
	tagNameRegex := "[a-z][a-z0-9,_=;-]*"
	tagRegex := fmt.Sprintf("<((%s)|/(%s)?)>", tagNameRegex, tagNameRegex)
	regex := regexp.MustCompile(tagRegex)

	tags := regex.FindAllString(text, -1)
	indexes := regex.FindAllStringIndex(text, -1)

	var positions []tagPos

	for i := 0; i < len(tags); i++ {
		// TODO find a clever way to remove <>
		tagName := strings.Replace(tags[i][1:], ">", "", -1)
		opening := true

		style := tagName

		if '/' == tagName[0] {
			opening = false
			style = tagName[1:]
		}

		positions = append(
			positions,
			tagPos{
				Text:    tags[i],
				Tag:     tagName,
				Style:   style,
				Start:   indexes[i][0],
				End:     indexes[i][1],
				Opening: opening,
			},
		)
	}

	return positions
}

// Applies current style from stack to text, if must be applied.
func (o *OutputFormatter) applyCurrentStyle(text string) string {
	if "" == text {
		return ""
	}

	if o.IsDecorated() {
		return o.GetStyleStack().GetCurrent().Apply(text)
	}

	return text
}

// create a style from a tag string
func (o *OutputFormatter) createStyleFromString(text string) *OutputFormatterStyle {
	text = strings.ToLower(text)

	if style, ok := o.stylesCache[text]; ok {
		return &style
	}

	regex := regexp.MustCompile("([^=]+)=([^;]+)(;|$)")
	matches := regex.FindAllStringSubmatch(text, -1)

	if nil == matches {
		return nil
	}

	style := NewOutputFormatterStyle(color.DEFAULT, color.DEFAULT, nil)

	for _, match := range matches {
		match = match[1:]

		if "fg" == match[0] {
			style.SetForeground(match[1])
		} else if "bg" == match[0] {
			style.SetBackground(match[1])
		} else if "options" == match[0] {
			optionRegex := regexp.MustCompile("([^,;]+)")
			options := optionRegex.FindAllString(match[1], -1)
			style.SetOptions(options)
		} else {
			return nil
		}
	}

	return style
}
