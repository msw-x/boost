package xstring

import (
	"crypto/md5"
	"crypto/sha1"

	"github.com/msw-x/boost/xfmt"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return xfmt.Hex(h.Sum(nil))
}

func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return xfmt.Hex(h.Sum(nil))
}
