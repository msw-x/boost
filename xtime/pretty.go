package xtime

import (
	"strings"
	"time"
)

func Pretty(t time.Duration) string {
	t = PrettyTruncate(t)
	s := t.String()
	if strings.HasSuffix(s, "m0s") {
		s = strings.TrimSuffix(s, "0s")
	}
	return s
}
