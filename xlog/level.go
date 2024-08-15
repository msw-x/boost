package xlog

import (
	"fmt"
	"strings"
)

type Level int

const (
	LevelNone Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

const DefaultLevel = LevelInfo

func (l Level) Laconic() string {
	switch l {
	case LevelTrace:
		return "trc"
	case LevelDebug:
		return "dbg"
	case LevelInfo:
		return "inf"
	case LevelWarning:
		return "wrn"
	case LevelError:
		return "err"
	case LevelCritical:
		return "crt"
	}
	return ""
}

func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarning:
		return "warning"
	case LevelError:
		return "error"
	case LevelCritical:
		return "critical"
	}
	return ""
}

func ParseLevel(s string) (l Level, err error) {
	switch strings.ToLower(s) {
	case "trace", "trc":
		l = LevelTrace
	case "debug", "dbg":
		l = LevelDebug
	case "info", "inf":
		l = LevelInfo
	case "warning", "wrn":
		l = LevelWarning
	case "error", "err":
		l = LevelError
	case "critical", "crt":
		l = LevelCritical
	default:
		err = fmt.Errorf("invalid log level: %q", s)
	}
	return
}
