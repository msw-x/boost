package xfmt

import (
	"fmt"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

const UnitPrefixInc = "KMGTPE"
const UnitPrefixDec = "mµnpfa" // TODO

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
	o.base = 1000
	o.precision = 1
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
				o.maxRank = len(UnitPrefixInc) + 1
			case ':':
				rank = &o.maxRank
			default:
				if '0' <= b && b <= '9' {
					ps += string(b)
				} else {
					i := slices.Index([]byte(UnitPrefixInc), b) + 1
					if i == 0 {
						i = -(slices.Index([]byte(UnitPrefixDec), b) + 1)
					}
					if i == 0 {
						break
					} else {
						*rank = i
						if rank == &o.minRank {
							o.maxRank = i
						}
					}
				}
			}
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
	negative := v < 0
	if negative {
		v = -v
	}
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
	tail := ""
	if o.precision > 0 {
		div := V(math.Pow(float64(base), float64(rank)))
		f := float64(v)/float64(div) - float64(i)
		g := fmt.Sprintf(fmt.Sprintf("%%.%df", o.precision), f)
		fract := g[2:]
		if g[0] == '1' {
			i++
		}
		if !o.fixed {
			ok := true
			for ok {
				fract, ok = strings.CutSuffix(fract, "0")
			}
		}
		if len(fract) > 0 {
			tail = "." + fract
		}
	}
	if len(o.space) == 0 {
		s = fmt.Sprintf("%d", i)
	} else {
		s = wideInt(i, o.space)
	}
	s += tail
	if rank != 0 {
		var r byte
		if rank > 0 {
			r = UnitPrefixInc[rank-1]
		} else {
			r = UnitPrefixDec[-rank+1]
		}
		s += o.space + string(r)
	}
	if negative {
		s = "-" + s
	}
	return
}
