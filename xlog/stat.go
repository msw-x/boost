package xlog

type Stat struct {
	Size     uint
	Trace    uint
	Debug    uint
	Info     uint
	Warning  uint
	Error    uint
	Critical uint
}

func (o *Stat) Push(level Level, size int) {
	o.Size += uint(size)
	switch level {
	case LevelTrace:
		o.Trace++
	case LevelDebug:
		o.Debug++
	case LevelInfo:
		o.Info++
	case LevelWarning:
		o.Warning++
	case LevelError:
		o.Error++
	case LevelCritical:
		o.Critical++
	}
}

func (o Stat) String() {
	/*
		ufmt.Int(count, ufmt.IntCtx{Precision: 0, Dense: true})
		xfmt.Int{Precision: 0, Dense: true}.Fmt(count)
		s := ""
		f := func(level Level, count uint) {
			if count > 0 {
				s = fmt.Sprintf("%s %v[%s]", s, level.Laconic())
			}
		}
		f(LevelTrace, o.stat.Trace)
		f(LevelDebug, o.stat.Debug)
		f(LevelInfo, o.stat.Info)
		f(LevelWarning, o.stat.Warning)
		f(LevelError, o.stat.Error)
		f(LevelCritical, o.stat.Critical)
	*/
}
