package xfmt

import (
	"testing"
)

func TestWideInt(t *testing.T) {
	t0 := func(v int, e string) {
		r := WideInt(v)
		if r != e {
			t.Errorf("WideInt(%d) = %q; expected: %q", v, r, e)
		}
	}
	t1 := func(v int, e string) {
		t0(v, e)
		t0(-v, "-"+e)
	}
	t1(2, "2")
	t1(475, "475")
	t1(4852, "4 852")
	t1(36775, "36 775")
	t1(938588, "938 588")
	t1(2383484, "2 383 484")
	t1(28475828, "28 475 828")
	t1(378485015, "378 485 015")
	t1(3049201294, "3 049 201 294")
}
