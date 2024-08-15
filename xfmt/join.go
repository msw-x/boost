package xfmt

import (
	"fmt"
	"strings"
)

func Join(v ...any) string {
	return JoinWith(" ", v...)
}

func JoinWith(s string, v ...any) string {
	return JoinSliceWith(s, v[:])
}

func JoinSlice[T any](v []T) string {
	return JoinSliceWith(" ", v)
}

func JoinSliceWith[T any](s string, v []T) string {
	l := make([]string, len(v))
	for n, a := range v {
		l[n] = fmt.Sprint(a)
	}
	return strings.Join(l, s)
}

func FactJoin(v ...any) string {
	return FactJoinWith(" ", v...)
}

func FactJoinWith(s string, v ...any) string {
	return FactJoinSliceWith(s, v[:])
}

func FactJoinSlice[T any](v []T) string {
	return FactJoinSliceWith(" ", v)
}

func FactJoinSliceWith[T any](s string, v []T) string {
	var l []string
	for _, a := range v {
		e := fmt.Sprint(a)
		if e != "" {
			l = append(l, e)
		}
	}
	return strings.Join(l, s)
}
