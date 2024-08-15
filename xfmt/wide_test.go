package xfmt

import (
	"testing"
)

func TestWideInt(t *testing.T) {
	test := func(v int, e string) {
		r := WideInt(v)
		if r != e {
			t.Errorf("WideInt(%d) = %q; expected: %q", v, r, e)
		}
	}
	test(2, "2")
	test(475, "475")
	test(4852, "4 852")
	test(36775, "36 775")
	test(938588, "938 588")
	test(2383484, "2 383 484")
	test(28475828, "28 475 828")
	test(378485015, "378 485 015")
	test(3049201294, "3 049 201 294")
}
