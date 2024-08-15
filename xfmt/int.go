package xfmt

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"constraints"
)

const UnitPrefix = "KMGTPE"

func Int[V constraints.Integer](f string, v V) string {
	var m fmtInt[V]
	err := m.parse(f)
	if err == nil {
		return m.convert(v)
	}
	return fmt.Sprintf("%%!(%d)", v)
}

type fmtInt[V constraints.Integer] struct {
	base      int
	space     string
	fixed     bool
	precision int
	minRank   int
	maxRank   int
}

func (o *fmtInt[V]) parse(f string) (err error) {
	r := strings.NewReader(f)
	ps := ""
	rank := &o.minRank
	for {
		var b byte
		b, err = r.ReadByte()
		if err == nil {
			switch b {
			case 'i':
				o.base = 1024
			case ' ':
				o.space = " "
			case '.':
			case ':':
				rank = &o.maxRank
			default:
				if '0' <= b && b <= '9' {
					ps += string(b)
				} else {
					i := slices.Index([]byte(UnitPrefix), b)
					if i == -1 {
						break
					} else {
						*rank = i
					}
				}
			}
			fmt.Println(string(b))
		} else {
			if err == io.EOF {
				err = nil
				ps, o.fixed = strings.CutPrefix(ps, "0")
				if ps == "" {
					if o.fixed {
						o.precision = 0
					}
				} else {
					o.precision, err = strconv.Atoi(ps)
				}
			}
			break
		}
	}
	return
}

func (o *fmtInt[V]) convert(v V) string {
	i := v
	rank := 0
	base := V(o.base)
	for rank < o.minRank {
		i /= base
		rank++
	}
	for i >= base && rank < o.maxRank {
		i /= base
		rank++
	}
	return ""
}
