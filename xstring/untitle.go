package xstring

import (
	"unicode"
	"unicode/utf8"
)

func UnTitle(s string) string {
	if s != "" {
		r, i := utf8.DecodeRuneInString(s)
		if i > 0 {
			return string(unicode.ToLower(r)) + s[i:]
		}
	}
	return s
}
