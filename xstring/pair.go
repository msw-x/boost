package xstring

import "strings"

func Pair(s string, sep string) (a string, b string) {
	l := strings.SplitN(s, sep, 2)
	n := len(l)
	if n > 0 {
		a = l[0]
	}
	if n > 1 {
		b = l[1]
	}
	return
}
