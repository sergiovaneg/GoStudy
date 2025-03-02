package fractals

import (
	"strings"
	"sync"
)

type NaiveDPSolver struct{}

type normalizer map[string]*string
type normalizedRuleset struct {
	r naiveRuleset
	n normalizer
}

func (NaiveDPSolver) String() string {
	return "Naïve DP Solver"
}

func (nr normalizedRuleset) get(serial string) (naiveFractal, bool) {
	f, ok := nr.r[*nr.n[serial]]
	return f, ok
}

func initNormalizedRuleset(lines []string) normalizedRuleset {
	var nr normalizedRuleset
	nr.r = make(naiveRuleset)
	nr.n = make(normalizer)

	for _, line := range lines {
		kv := strings.Split(line, " => ")
		nr.r[kv[0]] = deserializeNaive(kv[1])

		f := deserializeNaive(kv[0])
		fm := f.mirror()

		for range 4 {
			s, sm := f.serializeNaive(), fm.serializeNaive()
			nr.n[s], nr.n[sm] = &kv[0], &kv[0]
			f, fm = f.rotate(), fm.rotate()
		}
	}

	return nr
}

func (nr normalizedRuleset) transform(f naiveFractal) naiveFractal {
	if out, ok := nr.get(f.serializeNaive()); ok {
		return out
	}

	panic("Unregistered source pattern: " + f.serializeNaive())

}

func (nr normalizedRuleset) grow(f naiveFractal) naiveFractal {
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
					nr.transform(f.getSubfractal(i, j, s0)))
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
	r := initNormalizedRuleset(lines)

	for range nIters {
		f = r.grow(f)
	}

	return f.count()
}
