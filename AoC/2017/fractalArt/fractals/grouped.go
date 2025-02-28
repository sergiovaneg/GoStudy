package fractals

import (
	"strings"
)

type GroupedSolver struct{}

func (s GroupedSolver) String() string { return "Grouped Solver" }

type state map[string]uint

// From a 4x4 seed to the following (isolated) 3 steps of seeds
// 4x4 -> 6x6 -> 9x9 -> 12x12 (9 * 4x4)
type DPMemory map[string]state
type DynamicProgram struct {
	rules  pointerRuleset
	memory DPMemory
}

func (dp *DynamicProgram) updateDPMemory(serial string) {
	f0 := deserializeNaive(serial)

	var f1 naiveFractal

	// 4x4 -> 6x6
	f1 = makeEmptyNaive(6)
	for i := range 2 {
		for j := range 2 {
			subFractal := f0.getSubfractal(i, j, 2)
			subFractal = dp.transform(subFractal)
			f1.setSubfractal(i, j, 3, subFractal)
		}
	}
	f0 = f1

	// 6x6 -> 9x9
	f1 = makeEmptyNaive(9)
	for i := range 3 {
		for j := range 3 {
			subFractal := f0.getSubfractal(i, j, 2)
			subFractal = dp.transform(subFractal)
			f1.setSubfractal(i, j, 3, subFractal)
		}
	}
	f0 = f1

	// 9x9 -> 12x12
	f1 = makeEmptyNaive(12)
	for i := range 3 {
		for j := range 3 {
			subFractal := f0.getSubfractal(i, j, 3)
			subFractal = dp.transform(subFractal)
			f1.setSubfractal(i, j, 4, subFractal)
		}
	}
	f0 = f1

	// 12x12 -> 9 * 4x4
	s := make(state)
	for i := range 3 {
		for j := range 3 {
			subfractal := f0.getSubfractal(i, j, 4)
			s[subfractal.serializeNaive()]++
		}
	}

	dp.memory[serial] = s
}

func initDP(lines []string) DynamicProgram {
	var dp DynamicProgram

	dp.rules = initPointerRuleset(lines)
	dp.memory = make(DPMemory)

	return dp
}

func (dp DynamicProgram) transform(f naiveFractal) naiveFractal {
	return dp.rules.transform(f)
}

func (dp DynamicProgram) grow(f naiveFractal) naiveFractal {
	return dp.rules.grow(f)
}

func (dp DynamicProgram) getInitialState(seed string) state {
	if len(seed) != 11 {
		panic("Invalid seed; initial matrix must be 3x3.")
	}

	f := deserializeNaive(seed)

	return state{
		dp.transform(f).serializeNaive(): 1,
	}
}

func (dp *DynamicProgram) growThrice(serial string) state {
	if len(serial) != 19 {
		panic("Invalid entry requested; key must be a 4x4 matrix serial.")
	}

	if _, ok := dp.memory[serial]; !ok {
		dp.updateDPMemory(serial)
	}

	return dp.memory[serial]
}

func (GroupedSolver) Solve(seed string, nIters int, lines []string) uint {
	if nIters == 0 {
		return uint(strings.Count(seed, "#"))
	}

	dp := initDP(lines)

	state0 := dp.getInitialState(seed)
	nIters-- // We get the initial state by running a single iteration

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
			f = dp.grow(f)
		}

		finalCount += cnt0 * f.count()
	}

	return finalCount
}
