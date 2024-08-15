package xlog

import (
	"fmt"
	"testing"
)

func TestLevelPrint(t *testing.T) {
	test := func(v Level, e string) {
		r := fmt.Sprint(v)
		if r != e {
			t.Errorf("fmt.Sprint(%q) = %q; expected: %q", v, r, e)
		}
	}
	test(LevelNone, "")
	test(LevelTrace, "trace")
	test(LevelDebug, "debug")
	test(LevelInfo, "info")
	test(LevelWarning, "warning")
	test(LevelError, "error")
	test(LevelCritical, "critical")
}

func TestLevelLaconic(t *testing.T) {
	test := func(v Level, e string) {
		r := v.Laconic()
		if r != e {
			t.Errorf("Laconic(%q) = %q; expected: %q", v, r, e)
		}
	}
	test(LevelNone, "")
	test(LevelTrace, "trc")
	test(LevelDebug, "dbg")
	test(LevelInfo, "inf")
	test(LevelWarning, "wrn")
	test(LevelError, "err")
	test(LevelCritical, "crt")
}

func TestParseLevel(t *testing.T) {
	test := func(v string, e1 Level, e2 string) {
		r1, err := ParseLevel(v)
		r2 := "nil"
		if err != nil {
			r2 = err.Error()
		}
		if r1 != e1 || r2 != e2 {
			t.Errorf("ParseLevel(%q) = (%q, %s); expected: (%q, %s)", v, r1, r2, e1, e2)
		}
	}
	test("", LevelNone, `invalid log level: ""`)
	test("iinfo", LevelNone, `invalid log level: "iinfo"`)
	test("trc", LevelTrace, "nil")
	test("TRACE", LevelTrace, "nil")
	test("dbg", LevelDebug, "nil")
	test("debug", LevelDebug, "nil")
	test("INF", LevelInfo, "nil")
	test("info", LevelInfo, "nil")
	test("wrn", LevelWarning, "nil")
	test("Warning", LevelWarning, "nil")
	test("Err", LevelError, "nil")
	test("error", LevelError, "nil")
	test("crt", LevelCritical, "nil")
	test("critical", LevelCritical, "nil")
}
