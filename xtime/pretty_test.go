package xtime

import (
	"testing"
	"time"
)

func TestPretty(t *testing.T) {
	t0 := func(v time.Duration, e string) {
		r := Pretty(v)
		if r != e {
			t.Errorf("Pretty(%v) = %q; expected: %q", v, r, e)
		}
	}
	t0(time.Second, "1s")
	t0(72*time.Second, "1m12s")
	t0(84*time.Minute, "1h24m")
	t0(84*time.Minute+16*time.Second, "1h24m16s")
	t0(84*time.Minute+16*time.Second+97*time.Millisecond, "1h24m16s")
	t0(16*time.Second+97*time.Millisecond, "16s")
	t0(97*time.Millisecond, "97ms")
}
