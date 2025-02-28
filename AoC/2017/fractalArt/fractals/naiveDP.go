package fractals

import (
	"strings"
	"sync"
)

type NaiveDPSolver struct{}

type pointerRuleset map[string]*naiveFractal

func (NaiveDPSolver) String() string {
	return "Naïve DP Solver"
}

func initPointerRuleset(lines []string) pointerRuleset {
	r := make(pointerRuleset)

	for _, line := range lines {
		kv := strings.Split(line, " => ")

		f := deserializeNaive(kv[0])
		fm := f.mirror()
		out := deserializeNaive(kv[1])

		for range 4 {
			s, sm := f.serializeNaive(), fm.serializeNaive()

			r[s], r[sm] = new(naiveFractal), new(naiveFractal)
			*r[s], *r[sm] = out, out

			f, fm = f.rotate(), fm.rotate()
		}
	}

	return r
}

func (r pointerRuleset) transform(f naiveFractal) naiveFractal {
	if out, ok := r[f.serializeNaive()]; ok {
		return *out
	}

	panic("Unregistered source pattern: " + f.serializeNaive())

}

func (r pointerRuleset) grow(f naiveFractal) naiveFractal {
	n := len(f)

	var s0, s1 int
	if n%2 == 0 {
		s0, s1 = 2, 3
	} else {
		s0, s1 = 3, 4
	}

	nSubfrac := n / s0
	fNext := makeEmptyNaive(nSubfrac * s1)

	var wg sync.WaitGroup
	wg.Add(nSubfrac)
	for i := range nSubfrac {
		go func() {
			for j := range nSubfrac {
				fNext.setSubfractal(
					i, j, s1,
					r.transform(f.getSubfractal(i, j, s0)))
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return fNext
}

func (NaiveDPSolver) Solve(seed string, nIters int, lines []string) uint {
	if nIters == 0 {
		return uint(strings.Count(seed, "#"))
	}

	f := deserializeNaive(seed)
	r := initPointerRuleset(lines)

	for range nIters {
		f = r.grow(f)
	}

	return f.count()
}
