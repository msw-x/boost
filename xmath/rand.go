package xmath

import "math/rand"

// Rand returns, as an int, a non-negative pseudo-random number in the half-open interval [min,max]
func Rand(min, max int) int {
	return rand.Intn(max-min+1) + min
}
