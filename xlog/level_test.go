package xlog

import (
	"fmt"
	"testing"
)

func TestLevelPrint(t *testing.T) {
	t0 := func(v Level, e string) {
		r := fmt.Sprint(v)
		if r != e {
			t.Errorf("fmt.Sprint(%q) = %q; expected: %q", v, r, e)
		}
	}
	t0(LevelNone, "")
	t0(LevelTrace, "trace")
	t0(LevelDebug, "debug")
	t0(LevelInfo, "info")
	t0(LevelWarning, "warning")
	t0(LevelError, "error")
	t0(LevelCritical, "critical")
}

func TestLevelLaconic(t *testing.T) {
	t0 := func(v Level, e string) {
		r := v.Laconic()
		if r != e {
			t.Errorf("Laconic(%q) = %q; expected: %q", v, r, e)
		}
	}
	t0(LevelNone, "")
	t0(LevelTrace, "trc")
	t0(LevelDebug, "dbg")
	t0(LevelInfo, "inf")
	t0(LevelWarning, "wrn")
	t0(LevelError, "err")
	t0(LevelCritical, "crt")
}

func TestParseLevel(t *testing.T) {
	t0 := func(v string, e1 Level, e2 string) {
		r1, err := ParseLevel(v)
		r2 := "nil"
		if err != nil {
			r2 = err.Error()
		}
		if r1 != e1 || r2 != e2 {
			t.Errorf("ParseLevel(%q) = (%q, %s); expected: (%q, %s)", v, r1, r2, e1, e2)
		}
	}
	t0("", LevelNone, `invalid log level: ""`)
	t0("iinfo", LevelNone, `invalid log level: "iinfo"`)
	t0("trc", LevelTrace, "nil")
	t0("TRACE", LevelTrace, "nil")
	t0("dbg", LevelDebug, "nil")
	t0("debug", LevelDebug, "nil")
	t0("INF", LevelInfo, "nil")
	t0("info", LevelInfo, "nil")
	t0("wrn", LevelWarning, "nil")
	t0("Warning", LevelWarning, "nil")
	t0("Err", LevelError, "nil")
	t0("error", LevelError, "nil")
	t0("crt", LevelCritical, "nil")
	t0("critical", LevelCritical, "nil")
}
