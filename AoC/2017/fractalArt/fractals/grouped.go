package fractals

import (
	"strings"
)

type GroupedSolver struct{}

func (s GroupedSolver) String() string { return "Grouped Solver" }

type state map[string]uint
type stateRuleset map[string]state
type groupedRuleset struct {
	rules      pointerRuleset
	stateRules stateRuleset
}

func initGR(lines []string) groupedRuleset {
	var dp groupedRuleset

	dp.rules = initPointerRuleset(lines)
	dp.stateRules = make(stateRuleset)

	return dp
}

func (gr *groupedRuleset) updateSR(serial string) {
	f0 := deserializeNaive(serial)

	var f1 naiveFractal

	// 3x3 -> 4x4
	f0 = gr.rules.transform(f0)

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
	s := make(state)
	for i := range 3 {
		for j := range 3 {
			subfractal := f0.getSubfractal(i, j, 3)
			s[subfractal.serializeNaive()]++
		}
	}

	gr.stateRules[serial] = s
}

func (gr *groupedRuleset) growThrice(serial string) state {
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

	dp := initGR(lines)
	state0 := state{seed: 1}

	for range nIters / 3 {
		state1 := make(state)

		for serial0, cnt0 := range state0 {
			for serial1, cnt1 := range dp.growThrice(serial0) {
				state1[serial1] += cnt0 * cnt1
			}
		}

		state0 = state1
	}

	var finalCount uint
	for serial0, cnt0 := range state0 {
		f := deserializeNaive(serial0)

		for range nIters % 3 {
			f = dp.rules.grow(f)
		}

		finalCount += cnt0 * f.count()
	}

	return finalCount
}
