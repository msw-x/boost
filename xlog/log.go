package xlog

type Log struct {
	name string
}

func New(name string) *Log {
	o := new(Log)
	o.name = name
	return o
}
