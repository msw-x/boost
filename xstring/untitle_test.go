package xstring

import "testing"

func TestUnTitle(t *testing.T) {
	test := func(v string, e string) {
		r := UnTitle(v)
		if r != e {
			t.Errorf("UnTitle(%q) = %q; expected: %q", v, r, e)
		}
	}
	test("title", "title")
	test("Title", "title")
	test("Title Two", "title Two")
}
