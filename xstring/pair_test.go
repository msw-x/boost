package xstring

import "testing"

func TestPair(t *testing.T) {
	test := func(v, s, e1, e2 string) {
		r1, r2 := Pair(v, s)
		if r1 != e1 || r2 != e2 {
			t.Errorf("Pair(%q, %q) = (%q, %q); expected: (%q, %q)", v, s, r1, r2, e1, e2)
		}
	}
	test("", ":", "", "")
	test("a", ":", "a", "")
	test("a:", ":", "a", "")
	test("a:b", ":", "a", "b")
	test("a:b:c", ":", "a", "b:c")
}
