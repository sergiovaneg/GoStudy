package fractals

import (
	"strings"
)

type improvedRuleset map[string]*naiveFractal

// From a 4x4 seed to the following (isolated) 3 steps of seeds
// 4x4 -> 6x6 -> 9x9 -> 12x12 (9 * 4x4)
type State map[string]int
type DPMemory map[string][3]State
type DynamicProgram struct {
	memory DPMemory
	rules  improvedRuleset
}

func (r *improvedRuleset) initRuleset(lines []string) {
	*r = make(improvedRuleset)
	for _, line := range lines {
		kv := strings.Split(line, " => ")
		var f naiveFractal
		f.deserialize(kv[1])

		(*r)[kv[0]] = &f
	}
}

func (r *improvedRuleset) transform(f naiveFractal) naiveFractal {
	if out, ok := (*r)[f.serialize()]; ok {
		return *out
	}

	f0, f1 := f, f.mirror()
	var fOut *naiveFractal
	validSerials := []string{}

	for range 4 {
		validSerials = append(validSerials, f0.serialize(), f1.serialize())
		if aux, ok := (*r)[f0.serialize()]; ok {
			fOut = aux
		}

		if aux, ok := (*r)[f1.serialize()]; ok {
			fOut = aux
		}

		f0 = f0.rotate()
		f1 = f1.rotate()
	}

	for _, serial := range validSerials {
		(*r)[serial] = fOut
	}

	return *fOut
}

func (r *improvedRuleset) getInitialState(seed string) State {
	if len(seed) != 11 {
		panic("Invalid seed; initial matrix must be 3x3.")
	}

	var f naiveFractal
	f.deserialize(seed)

	return State{
		r.transform(f).serialize(): 1,
	}
}

func (dp *DynamicProgram) calculateEntry(serial string) {
	var f0 naiveFractal
	f0.deserialize(serial)

	var f1 naiveFractal
	var aux [3]State

	// 4x4 -> 6x6
	f1.makeEmpty(6)
	for i := range 2 {
		for j := range 2 {
			subFractal := f0.getSubfractal(i, j, 2)
			subFractal = dp.rules.transform(subFractal)
			f1.setSubfractal(i, j, 3, subFractal)
		}
	}
	aux[0] = State{f1.serialize(): 1}
	f0 = f1

	// 6x6 -> 9x9
	f1.makeEmpty(9)
	for i := range 3 {
		for j := range 3 {
			subFractal := f0.getSubfractal(i, j, 2)
			subFractal = dp.rules.transform(subFractal)
			f1.setSubfractal(i, j, 3, subFractal)
		}
	}
	aux[1] = State{f1.serialize(): 1}
	f0 = f1

	// 9x9 -> 12x12
	f1.makeEmpty(12)
	for i := range 3 {
		for j := range 3 {
			subFractal := f0.getSubfractal(i, j, 3)
			subFractal = dp.rules.transform(subFractal)
			f1.setSubfractal(i, j, 4, subFractal)
		}
	}
	f0 = f1

	// 12x12 -> 9 * 4x4
	aux[2] = make(State)
	for i := range 3 {
		for j := range 3 {
			subfractal := f0.getSubfractal(i, j, 4)
			aux[2][subfractal.serialize()] += 1
		}
	}

	dp.memory[serial] = aux
}

func (dp *DynamicProgram) getNthState(serial string, n int) State {
	if len(serial) != 19 {
		panic("Invalid entry requested; key must be a 4x4 matrix serial.")
	}

	if _, ok := dp.memory[serial]; !ok {
		dp.calculateEntry(serial)
	}

	return dp.memory[serial][n]
}

func (s State) count() int {
	var res int
	for serial, cnt := range s {
		res += strings.Count(serial, "#") * cnt
	}
	return res
}

type ImprovedSolver struct{}

func (ImprovedSolver) Solve(seed string, nIters int, lines []string) int {
	if nIters == 0 {
		return State{seed: 1}.count()
	}

	var r improvedRuleset
	r.initRuleset(lines)

	state0 := r.getInitialState(seed)
	nIters-- // We get the initial state by running an iteration

	dp := DynamicProgram{
		memory: make(DPMemory),
		rules:  r,
	}

	for range nIters / 3 {
		state1 := make(State)

		for serial0, cnt0 := range state0 {
			for serial1, cnt1 := range dp.getNthState(serial0, 2) {
				state1[serial1] += cnt0 * cnt1
			}
		}

		state0 = state1
	}

	if nIters%3 == 0 {
		return state0.count()
	}

	state1 := make(State)
	for serial0, cnt0 := range state0 {
		for serial1, cnt1 := range dp.getNthState(serial0, nIters%3-1) {
			state1[serial1] += cnt0 * cnt1
		}
	}

	return state1.count()
}
