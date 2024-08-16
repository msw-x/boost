package xmath

import "math/rand"

func Rand(min, max int) int {
	return rand.Intn(max-min) + min
}
