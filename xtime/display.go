package xtime

import (
	"fmt"
	"time"
)

func Display(t time.Duration) string {
	return makeDisplay(t).std()
}

func DisplayDays(t time.Duration) string {
	return makeDisplay(t).days()
}

func DisplayMs(t time.Duration) string {
	return makeDisplay(t).mls()
}

func DisplayAuto(t time.Duration) string {
	return makeDisplay(t).auto()
}

type display struct {
	h  int
	m  int
	s  int
	ms int
}

func makeDisplay(t time.Duration) display {
	t = t.Round(time.Millisecond)
	h := t / time.Hour
	t -= h * time.Hour
	m := t / time.Minute
	t -= m * time.Minute
	s := t / time.Second
	t -= s * time.Second
	ms := t / time.Millisecond
	return display{int(h), int(m), int(s), int(ms)}
}

func (o display) std() string {
	return fmt.Sprintf("%02d:%02d:%02d", o.h, o.m, o.s)
}

func (o display) days() string {
	days := o.h / 24
	o.h = o.h % 24
	s := o.std()
	if days > 0 {
		plural := ""
		if days > 1 {
			plural = "s"
		}
		s = fmt.Sprintf("%d day%s %s", days, plural, s)
	}
	return s
}

func (o display) mls() string {
	return fmt.Sprintf("%s.%03d", o.std(), o.ms)
}

func (o display) auto() string {
	if o.h == 0 && o.m == 0 && o.s == 0 {
		return fmt.Sprintf("%d ms", o.ms)
	}
	return o.days()
}
