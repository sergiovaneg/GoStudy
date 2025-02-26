//go:build naive

package main

import (
	"slices"
	"strings"
)

type fractal [][]byte
type ruleset map[string]fractal

func (f fractal) serialize() string {
	rows := make([]string, len(f))

	for i, row := range f {
		rows[i] = string(row)
	}

	return strings.Join(rows, "/")
}

func deserialize(serial string) fractal {
	f := make(fractal, 0)

	for _, row := range strings.Split(serial, "/") {
		f = append(f, []byte(row))
	}

	return f
}

func initRuleset(lines []string) ruleset {
	r := make(ruleset)
	for _, line := range lines {
		kv := strings.Split(line, " => ")
		r[kv[0]] = deserialize(kv[1])
	}
	return r
}

func makeEmpty(n int) fractal {
	f := make(fractal, n)
	for i := range n {
		f[i] = make([]byte, n)
	}
	return f
}

func (f fractal) mirror() fractal {
	fNew := make(fractal, len(f))

	for i, row := range f {
		fNew[i] = slices.Clone(row)
		slices.Reverse(fNew[i])
	}

	return fNew
}

func (f fractal) rotate() fractal {
	n := len(f)
	fNew := makeEmpty(n)

	for i := range n {
		for j := range n {
			fNew[i][j] = f[j][n-i-1]
		}
	}

	return fNew
}

func (r ruleset) transform(f fractal) fractal {
	fm := f.mirror()

	for range 4 {
		if aux, ok := r[f.serialize()]; ok {
			return aux
		}

		if aux, ok := r[fm.serialize()]; ok {
			return aux
		}

		f = f.rotate()
		fm = fm.rotate()
	}

	return nil
}

func (f fractal) getSubfractal(i, j, n int) fractal {
	subFractal := makeEmpty(n)
	for k := range n {
		copy(subFractal[k], f[n*i+k][n*j:])
	}
	return subFractal
}

func (f *fractal) setSubfractal(i, j, n int, sf fractal) {
	for k := range n {
		copy((*f)[n*i+k][n*j:], sf[k])
	}
}

func (r ruleset) grow(f fractal) fractal {
	n := len(f)

	var s0, s1 int
	if n%2 == 0 {
		s0, s1 = 2, 3
	} else {
		s0, s1 = 3, 4
	}
	fNext := makeEmpty(n * s1 / s0)

	for i := range n / s0 {
		for j := range n / s0 {
			fNext.setSubfractal(
				i, j, s1,
				r.transform(f.getSubfractal(i, j, s0)))
		}
	}

	return fNext
}

func (f fractal) count() int {
	var res int
	for _, row := range f {
		for _, v := range row {
			if v == '#' {
				res++
			}
		}
	}
	return res
}

func init() {
	Run = func(seed string, nIters int, lines []string) int {
		r := initRuleset(lines)
		f := deserialize(seed)

		for range nIters {
			f = r.grow(f)
		}

		return f.count()
	}
}
