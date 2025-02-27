package fractals

import (
	"strings"
	"sync"
)

type NaiveDPSolver struct{}

func (NaiveDPSolver) String() string {
	return "Naïve DP Solver"
}

func initImprovedRuleset(lines []string) improvedRuleset {
	r := make(improvedRuleset)

	for _, line := range lines {
		kv := strings.Split(line, " => ")

		f := deserializeNaive(kv[0])
		fm := f.mirror()
		fOut := deserializeNaive(kv[1])

		for range 4 {
			s, sm := f.serializeNaive(), fm.serializeNaive()

			r[s], r[sm] = new(naiveFractal), new(naiveFractal)
			*r[s], *r[sm] = fOut, fOut

			f, fm = f.rotate(), fm.rotate()
		}
	}

	return r
}

func (r improvedRuleset) transform(f naiveFractal) naiveFractal {
	if out, ok := r[f.serializeNaive()]; ok {
		return *out
	} else {
		panic("Unregistered source pattern: " + f.serializeNaive())
	}
}

func (r improvedRuleset) grow(f naiveFractal) naiveFractal {
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

func (NaiveDPSolver) Solve(seed string, nIters int, lines []string) int {
	r := initImprovedRuleset(lines)

	f := deserializeNaive(seed)

	for range nIters {
		f = r.grow(f)
	}

	return f.count()
}
