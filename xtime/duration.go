package xtime

import (
	"strings"
	"time"
)

func Duration(t time.Time) time.Duration {
	return time.Hour*time.Duration(t.Hour()) + time.Minute*time.Duration(t.Minute()) + time.Second*time.Duration(t.Second())
}

func ParseDuration(s string) (r time.Duration, err error) {
	if strings.Contains(s, ":") {
		var t time.Time
		t, err = time.Parse("15:04:05", s)
		if err == nil {
			r = Duration(t)
		}
		return
	}
	return time.ParseDuration(s)
}
