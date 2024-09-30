package rt

import (
	"reflect"
	"runtime"
)

func FuncName(f any) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
