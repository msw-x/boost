package xtime

import "time"

type Stopwatch struct {
	ts time.Time
}

func NewStopwatch() *Stopwatch {
	o := new(Stopwatch)
	o.Reset()
	return o
}

func (o *Stopwatch) Reset() {
	o.ts = time.Now()
}

func (o Stopwatch) Time() time.Duration {
	return time.Since(o.ts)
}

func (o Stopwatch) String() string {
	return Pretty(o.Time())
}
