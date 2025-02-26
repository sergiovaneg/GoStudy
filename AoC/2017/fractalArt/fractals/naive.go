package fractals

import (
	"slices"
	"strings"
)

type naiveFractal [][]byte
type naiveRuleset map[string]naiveFractal

func (f naiveFractal) serialize() string {
	rows := make([]string, len(f))

	for i, row := range f {
		rows[i] = string(row)
	}

	return strings.Join(rows, "/")
}

func (f *naiveFractal) deserialize(serial string) {
	*f = make(naiveFractal, 0)

	for _, row := range strings.Split(serial, "/") {
		*f = append(*f, []byte(row))
	}
}

func (r *naiveRuleset) initRuleset(lines []string) {
	*r = make(naiveRuleset)
	for _, line := range lines {
		kv := strings.Split(line, " => ")
		var f naiveFractal
		f.deserialize(kv[1])
		(*r)[kv[0]] = f
	}
}

func (f *naiveFractal) makeEmpty(n int) {
	*f = make(naiveFractal, n)
	for i := range n {
		(*f)[i] = make([]byte, n)
	}
}

func (f naiveFractal) mirror() naiveFractal {
	fNew := make(naiveFractal, len(f))

	for i, row := range f {
		fNew[i] = slices.Clone(row)
		slices.Reverse(fNew[i])
	}

	return fNew
}

func (f naiveFractal) rotate() naiveFractal {
	n := len(f)

	var fNew naiveFractal
	fNew.makeEmpty(n)

	for i := range n {
		for j := range n {
			fNew[i][j] = f[j][n-i-1]
		}
	}

	return fNew
}

func (r naiveRuleset) transform(f naiveFractal) naiveFractal {
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

func (f naiveFractal) getSubfractal(i, j, n int) naiveFractal {
	var subFractal naiveFractal
	subFractal.makeEmpty(n)

	for k := range n {
		copy(subFractal[k], f[n*i+k][n*j:])
	}

	return subFractal
}

func (f *naiveFractal) setSubfractal(i, j, n int, sf naiveFractal) {
	for k := range n {
		copy((*f)[n*i+k][n*j:], sf[k])
	}
}

func (r naiveRuleset) grow(f naiveFractal) naiveFractal {
	n := len(f)

	var s0, s1 int
	if n%2 == 0 {
		s0, s1 = 2, 3
	} else {
		s0, s1 = 3, 4
	}

	var fNext naiveFractal
	fNext.makeEmpty(n * s1 / s0)

	for i := range n / s0 {
		for j := range n / s0 {
			fNext.setSubfractal(
				i, j, s1,
				r.transform(f.getSubfractal(i, j, s0)))
		}
	}

	return fNext
}

func (f naiveFractal) count() int {
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

type NaiveSolver struct{}

func (NaiveSolver) Solve(seed string, nIters int, lines []string) int {
	var r naiveRuleset
	r.initRuleset(lines)

	var f naiveFractal
	f.deserialize(seed)

	for range nIters {
		f = r.grow(f)
	}

	return f.count()
}
