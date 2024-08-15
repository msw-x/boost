package xtime

import (
	"testing"
	"time"
)

func TestPretty(t *testing.T) {
	test := func(v time.Duration, e string) {
		r := Pretty(v)
		if r != e {
			t.Errorf("Pretty(%v) = %q; expected: %q", v, r, e)
		}
	}
	test(time.Second, "1s")
	test(72*time.Second, "1m12s")
	test(84*time.Minute, "1h24m")
	test(84*time.Minute+16*time.Second, "1h24m16s")
	test(84*time.Minute+16*time.Second+97*time.Millisecond, "1h24m16s")
	test(16*time.Second+97*time.Millisecond, "16s")
	test(97*time.Millisecond, "97ms")
}
