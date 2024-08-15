package xlog

import (
	"time"
)

type Options struct {
	Level        Level
	AppName      string
	TimeLocation *time.Location
	StickGoId    bool
	SplitArgs    bool

	Console          bool
	CriticalToStdErr bool

	File       string
	FileAppend bool

	Dir string

	FileSizeLimit  uint64
	FileTimeLimit  time.Duration
	DaysCountLimit int
	TotalSizeLimit uint64
}

func DefaultOptions() (o Options) {
	o.Level = LevelInfo
	o.SplitArgs = true
	o.Console = true
	o.CriticalToStdErr = true
	return o
}
