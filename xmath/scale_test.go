package xmath

import (
	"testing"
)

func TestScale(t *testing.T) {
	t0 := func(v xmath.Number, e int) {
		r := Scale(v)
		if r != e {
			t.Errorf("Scale(%f) = %d; expected: %d", v, r, e)
		}
	}
	t1 := func(v xmath.Number, e int) {
		t0(v, e)
		t0(-v, e)
	}
	t1(1, 0)
	t1(0.1, -1)
	t1(0.01, -2)
	t1(0.001, -3)
	t1(0.0001, -4)
	t1(0.00001, -5)
	t1(0.000001, -6)
	t1(0.0000001, -7)
	t1(0.00000001, -8)
	t1(0.000000001, -9)
	t1(0.0000000001, -10)
	t1(0.00000000001, -11)
	t1(0.000000000001, -12)
}
