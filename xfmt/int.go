package xfmt

import (
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/msw-x/boost/xstring"
	"golang.org/x/exp/constraints"
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

func (o *fmtInt[V]) convert(v V) (s string) {
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
	if len(o.space) == 0 {
		s = fmt.Sprintf("%d", i)
	} else {
		s = wideInt(i, o.space)
	}
	if o.precision > 0 {
		div := V(math.Pow(float64(base), float64(rank)))
		f := float64(v) / float64(div)
		_, fract := xstring.Pair(fmt.Sprintf(fmt.Sprintf("%%.%df", o.precision), f), ".")
		if !o.fixed {
			ok := true
			for ok {
				fract, ok = strings.CutSuffix(fract, "0")
			}
		}
		if len(fract) > 0 {
			s = fmt.Sprintf("%s.%s", s, fract)
		}
	}
	prefix := ""
	if rank > 0 {
		prefix = string(UnitPrefix[rank-1])
	}
	s = fmt.Sprintf("%s%s%s", s, o.space, prefix)
	return
}
