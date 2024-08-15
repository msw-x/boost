package xstring

import "testing"

func TestPair(t *testing.T) {
	t0 := func(v, s, e1, e2 string) {
		r1, r2 := Pair(v, s)
		if r1 != e1 || r2 != e2 {
			t.Errorf("Pair(%q, %q) = (%q, %q); expected: (%q, %q)", v, s, r1, r2, e1, e2)
		}
	}
	t0("", ":", "", "")
	t0("a", ":", "a", "")
	t0("a:", ":", "a", "")
	t0("a:b", ":", "a", "b")
	t0("a:b:c", ":", "a", "b:c")
}
