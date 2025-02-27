package fractals

import "sync"

type NaiveParallelSolver struct{}

func (NaiveParallelSolver) String() string { return "Naïve Parallel Solver" }

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

func (NaiveParallelSolver) Solve(seed string, nIters int, lines []string) int {
	r := initNaiveRuleset(lines)
	f := deserializeNaive(seed)

	for range nIters {
		f = r.growParallel(f)
	}

	return f.count()
}
