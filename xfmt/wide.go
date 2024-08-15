package xfmt

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func wideInt[V constraints.Integer](v V, space string) string {
	s := fmt.Sprint(v)
	var parts []string
	for len(s) > 3 {
		parts = append(parts, s[len(s)-3:])
		s = s[:len(s)-3]
	}
	if len(s) > 0 {
		parts = append(parts, s)
	}
	s = ""
	n := len(parts)
	for i := n - 1; i >= 0; i-- {
		if len(s) != 0 {
			s += space
		}
		s += parts[i]
	}
	return s
}
