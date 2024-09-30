package xmath

import (
	"math"
)

// Order of magnitude
func Scale[T Number](v T) int {
	if v == 0 {
		return 1 // undefined
	}
	f := math.Abs(float64(v))
	n := int(math.Floor(math.Log10(f)))
	if f >= 1 {
		n++
	}
	return n
}
