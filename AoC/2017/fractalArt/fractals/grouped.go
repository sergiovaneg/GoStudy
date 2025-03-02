package fractals

import (
	"strings"
)

type GroupedSolver struct{}

func (s GroupedSolver) String() string { return "Grouped Solver" }

type stateSequence struct {
	s1, s2 string
	s3     map[string]uint
}
type stateRuleset map[string]stateSequence
type groupedRuleset struct {
	rules      normalizedRuleset
	stateRules stateRuleset
}

func initGR(lines []string) groupedRuleset {
	var dp groupedRuleset

	dp.rules = initNormalizedRuleset(lines)
	dp.stateRules = make(stateRuleset)

	return dp
}

func (gr *groupedRuleset) updateSR(serial string) {
	f0 := deserializeNaive(serial)

	var f1 naiveFractal
	s := stateSequence{
		s3: make(map[string]uint),
	}

	// 3x3 -> 4x4
	f0 = gr.rules.transform(f0)
	s.s1 = f0.serializeNaive()

	// 4x4 -> 6x6
	f1 = makeEmptyNaive(6)
	for i := range 2 {
		for j := range 2 {
			subFractal := f0.getSubfractal(i, j, 2)
			subFractal = gr.rules.transform(subFractal)
			f1.setSubfractal(i, j, 3, subFractal)
		}
	}
	f0 = f1
	s.s2 = f0.serializeNaive()

	// 6x6 -> 9x9
	f1 = makeEmptyNaive(9)
	for i := range 3 {
		for j := range 3 {
			subFractal := f0.getSubfractal(i, j, 2)
			subFractal = gr.rules.transform(subFractal)
			f1.setSubfractal(i, j, 3, subFractal)
		}
	}
	f0 = f1

	// 9x9 -> 9 * 3x3
	for i := range 3 {
		for j := range 3 {
			subfractal := f0.getSubfractal(i, j, 3)
			normSerial := *gr.rules.n[subfractal.serializeNaive()]
			s.s3[normSerial]++
		}
	}

	gr.stateRules[serial] = s
}

func (gr *groupedRuleset) growThrice(serial string) stateSequence {

	if _, ok := gr.stateRules[serial]; !ok {
		gr.updateSR(serial)
	}

	return gr.stateRules[serial]
}

func (GroupedSolver) Solve(seed string, nIters int, lines []string) uint {
	if len(seed) != 11 {
		panic("Invalid seed: serial should match a 3x3 fractal.")
	}

	if nIters == 0 {
		return uint(strings.Count(seed, "#"))
	}

	gr := initGR(lines)

	// Normalize from the beginning
	state0 := map[string]uint{*gr.rules.n[seed]: 1}

	for range nIters / 3 {
		state1 := make(map[string]uint)

		for serial0, cnt0 := range state0 {
			for serial1, cnt1 := range gr.growThrice(serial0).s3 {
				state1[serial1] += cnt0 * cnt1
			}
		}

		state0 = state1
	}

	var finalCount uint

	switch nIters % 3 {
	case 0:
		for serial, cnt := range state0 {
			finalCount += cnt * uint(strings.Count(serial, "#"))
		}
	case 1:
		for serial, cnt := range state0 {
			finalCount += cnt * uint(strings.Count(
				gr.growThrice(serial).s1, "#"))
		}
	case 2:
		for serial, cnt := range state0 {
			finalCount += cnt * uint(strings.Count(
				gr.growThrice(serial).s2, "#"))
		}
	}

	return finalCount
}
