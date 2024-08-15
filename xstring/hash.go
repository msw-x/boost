package xstring

import (
	"crypto/md5"
	"crypto/sha1"
)

func Md5(s string) []byte {
	h := md5.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func Sha1(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
