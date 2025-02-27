package fractals

import "strings"

type ImprovedSolver struct{}

func (s ImprovedSolver) String() string { return "Improved Solver" }

type improvedRuleset map[string]*naiveFractal
type state map[string]int

// From a 4x4 seed to the following (isolated) 3 steps of seeds
// 4x4 -> 6x6 -> 9x9 -> 12x12 (9 * 4x4)
type DPMemory map[string]state
type DynamicProgram struct {
	memory DPMemory
	rules  improvedRuleset
}

func initDP(lines []string) DynamicProgram {
	var dp DynamicProgram

	dp.memory = make(DPMemory)
	dp.rules = make(improvedRuleset)

	for _, line := range lines {
		kv := strings.Split(line, " => ")
		dp.rules[kv[0]] = new(naiveFractal)
		*dp.rules[kv[0]] = deserializeNaive(kv[1])
	}

	return dp
}

func (dp *DynamicProgram) transform(f naiveFractal) naiveFractal {
	if out, ok := dp.rules[f.serializeNaive()]; ok {
		return *out
	}

	f0, f1 := f, f.mirror()
	var fOut *naiveFractal
	validSerials := []string{}

	for range 4 {
		validSerials = append(validSerials, f0.serializeNaive(), f1.serializeNaive())
		if aux, ok := dp.rules[f0.serializeNaive()]; ok {
			fOut = aux
		}

		if aux, ok := dp.rules[f1.serializeNaive()]; ok {
			fOut = aux
		}

		f0 = f0.rotate()
		f1 = f1.rotate()
	}

	if fOut == nil {
		panic("Unregistered source pattern:" + f.serializeNaive())
	}

	for _, serial := range validSerials {
		dp.rules[serial] = fOut
	}

	return *fOut
}

func (dp *DynamicProgram) grow(f naiveFractal) naiveFractal {
	n := len(f)

	var s0, s1 int
	if n%2 == 0 {
		s0, s1 = 2, 3
	} else {
		s0, s1 = 3, 4
	}

	fNext := makeEmptyNaive(n * s1 / s0)

	for i := range n / s0 {
		for j := range n / s0 {
			fNext.setSubfractal(
				i, j, s1,
				dp.transform(f.getSubfractal(i, j, s0)))
		}
	}

	return fNext
}

func (dp *DynamicProgram) getInitialState(seed string) state {
	if len(seed) != 11 {
		panic("Invalid seed; initial matrix must be 3x3.")
	}

	f := deserializeNaive(seed)

	return state{
		dp.transform(f).serializeNaive(): 1,
	}
}

func (dp *DynamicProgram) calculateEntry(serial string) {
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
	aux := make(state)
	for i := range 3 {
		for j := range 3 {
			subfractal := f0.getSubfractal(i, j, 4)
			aux[subfractal.serializeNaive()] += 1
		}
	}

	dp.memory[serial] = aux
}

func (dp *DynamicProgram) growThrice(serial string) state {
	if len(serial) != 19 {
		panic("Invalid entry requested; key must be a 4x4 matrix serial.")
	}

	if _, ok := dp.memory[serial]; !ok {
		dp.calculateEntry(serial)
	}

	return dp.memory[serial]
}

func (s state) count() int {
	var res int

	for serial, cnt := range s {
		res += strings.Count(serial, "#") * cnt
	}

	return res
}

func (ImprovedSolver) Solve(seed string, nIters int, lines []string) int {
	if nIters == 0 {
		return state{seed: 1}.count()
	}

	dp := initDP(lines)

	state0 := dp.getInitialState(seed)
	nIters-- // We get the initial state by running an iteration

	for range nIters / 3 {
		state1 := make(state)

		for serial0, cnt0 := range state0 {
			for serial1, cnt1 := range dp.growThrice(serial0) {
				state1[serial1] += cnt0 * cnt1
			}
		}

		state0 = state1
	}

	state1 := make(state, len(state0))
	for serial0, cnt0 := range state0 {
		f := deserializeNaive(serial0)

		for range nIters % 3 {
			f = dp.grow(f)
		}

		state1[f.serializeNaive()] = cnt0
	}

	return state1.count()
}
