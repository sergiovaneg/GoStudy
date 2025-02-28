package fractals

import (
	"strings"
	"sync"
)

type NaiveConcurrentSolver struct{}

func (NaiveConcurrentSolver) String() string {
	return "Naïve Concurrent Solver"
}

func (r naiveRuleset) growParallel(f naiveFractal) naiveFractal {
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

func (NaiveConcurrentSolver) Solve(
	seed string, nIters int, lines []string) uint {
	if nIters == 0 {
		return uint(strings.Count(seed, "#"))
	}

	f := deserializeNaive(seed)
	r := initNaiveRuleset(lines)

	for range nIters {
		f = r.growParallel(f)
	}

	return f.count()
}
