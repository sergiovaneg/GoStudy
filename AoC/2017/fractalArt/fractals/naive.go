package fractals

import (
	"slices"
	"strings"
)

type NaiveSolver struct{}

func (s NaiveSolver) String() string { return "Naïve Solver" }

type naiveFractal [][]byte
type naiveRuleset map[string]naiveFractal

func (f naiveFractal) serializeNaive() string {
	rows := make([]string, len(f))

	for i, row := range f {
		rows[i] = string(row)
	}

	return strings.Join(rows, "/")
}

func deserializeNaive(serial string) naiveFractal {
	f := make(naiveFractal, 0)

	for row := range strings.SplitSeq(serial, "/") {
		f = append(f, []byte(row))
	}

	return f
}

func initNaiveRuleset(lines []string) naiveRuleset {
	r := make(naiveRuleset)

	for _, line := range lines {
		kv := strings.Split(line, " => ")
		r[kv[0]] = deserializeNaive(kv[1])
	}

	return r
}

func makeEmptyNaive(n int) naiveFractal {
	f := make(naiveFractal, n)

	for i := range n {
		f[i] = make([]byte, n)
	}

	return f
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

	fNew := makeEmptyNaive(n)

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
		if aux, ok := r[f.serializeNaive()]; ok {
			return aux
		}

		if aux, ok := r[fm.serializeNaive()]; ok {
			return aux
		}

		f = f.rotate()
		fm = fm.rotate()
	}

	panic("Unregistered source pattern: " + f.serializeNaive())
}

func (f naiveFractal) getSubfractal(i, j, n int) naiveFractal {
	subFractal := makeEmptyNaive(n)

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

	nSubfrac := n / s0
	fNext := makeEmptyNaive(nSubfrac * s1)

	for i := range nSubfrac {
		for j := range nSubfrac {
			fNext.setSubfractal(
				i, j, s1,
				r.transform(f.getSubfractal(i, j, s0)))
		}
	}

	return fNext
}

func (f naiveFractal) count() uint {
	var res uint

	for _, row := range f {
		for _, v := range row {
			if v == '#' {
				res++
			}
		}
	}

	return res
}

func (NaiveSolver) Solve(seed string, nIters int, lines []string) uint {
	if nIters == 0 {
		return uint(strings.Count(seed, "#"))
	}

	f := deserializeNaive(seed)
	r := initNaiveRuleset(lines)

	for range nIters {
		f = r.grow(f)
	}

	return f.count()
}
