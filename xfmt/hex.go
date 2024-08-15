package ufmt

import (
	"encoding/hex"
)

func Hex(buf []byte) string {
	return hex.EncodeToString(buf)
}
