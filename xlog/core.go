package xlog

import (
	"os"
	"sync"
	"time"
)

type core struct {
	mutex    sync.Mutex
	opts     Options
	stat     Stat
	file     *os.File
	fileSize uint64
	fileTime time.Time
	fname    string
	timeLoc  *time.Location
	maxid    int
	mapid    map[int]bool
	//hook     func(Message)
	initedAt time.Time
}

var c core

func (o *core) init(opts Options) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	openFile := o.opts.File != opts.File || o.opts.Dir != opts.Dir
	o.opts = opts
	if openFile {
		//o.openFile(false)
	}
	o.maxid = 2
	o.mapid = make(map[int]bool)
	o.initedAt = time.Now()
}

func (o *core) close() {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if o.file != nil {
		o.file.Close()
		o.file = nil
	}
}

func (o *core) stat2() string {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	/*
		s := fmt.Sprintf("%s | %s", xtime.DisplayAuto(time.Since(o.initedAt)), ufmt.ByteSize(o.stat.Size))
		if o.opts.GoID {
			s = fmt.Sprintf("%s go[%s]", s, ufmt.WideInt(len(o.mapid)))
		}
		if o.stat.Size != 0 {
			s = fmt.Sprintf("%s %s", s, o.stat.String())
		}
		return s
	*/
	return ""
}

/*
func printText(ctx *context, level Level, v ...any) {
	if level >= ctx.opts.level {
		m := NewMessage(ctx, level, v...)
		printMessage(ctx, level, m)
		if ctx.hook != nil {
			ctx.hook(m)
		}
	}
}

func printMessage(ctx *context, level Level, m Message) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	ctx.stat.Push(level, m.Size())
	if ctx.opts.Console || level == LevelCritical {
		if level >= LevelError {
			if ctx.opts.Console {
				printStdErr(m.Format())
			} else {
				if ctx.opts.CrtStdErr {
					printStdErr(m.Text)
				}
			}
		} else {
			fmt.Print(m.Format())
		}
	}
	if ctx.file != nil {
		ctx.trim(m.Size())
		ctx.fileSize += uint64(m.Size())
		ctx.file.WriteString(m.Format())
	}
}

func printStdErr(text string) {
	fmt.Fprint(os.Stderr, text)
}
*/
