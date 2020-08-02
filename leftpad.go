package leftpad

import (
	"strings"
	"unicode/utf8"
)

var defaultChar = ' '

/*Format ..*/
func Format(s string, size int) string {
	return FormatRune(s, size, defaultChar)
}

/*FormatRune ..*/
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}

	out := strings.Repeat(string(r), size-1) + s
	return out
}
