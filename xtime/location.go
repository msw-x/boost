package xtime

import (
	"fmt"
	"time"
)

func FixedZone(offset time.Duration) *time.Location {
	h := int(offset.Hours())
	s := int(offset.Seconds())
	z := fmt.Sprintf("UTC%+d", h)
	return time.FixedZone(z, s)
}

func MoveLocation(t time.Time, loc *time.Location) time.Time {
	_, src := t.Zone()
	_, dst := time.Now().In(loc).Zone()
	return t.In(loc).Add(time.Second * time.Duration(src-dst))
}

func Timezone(s string) (*time.Location, error) {
	switch s {
	case "MSK":
		s = "Europe/Moscow"
	}
	return time.LoadLocation(s)
}
