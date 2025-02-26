//go:build improved

package main

import (
	"slices"
	"strings"
)

type fractal [][]byte
type ruleset map[string]*fractal

// From a 4x4 seed to the following (isolated) 3 steps of seeds
// 4x4 -> 6x6 -> 9x9 -> 12x12 (9 * 4x4)
type State map[string]int
type DPMemory map[string][3]State
type DynamicProgram struct {
	memory DPMemory
	rules  ruleset
}

func (f fractal) serialize() string {
	rows := make([]string, len(f))

	for i, row := range f {
		rows[i] = string(row)
	}

	return strings.Join(rows, "/")
}

func deserialize(serial string) fractal {
	f := make(fractal, 0)

	for _, row := range strings.Split(serial, "/") {
		f = append(f, []byte(row))
	}

	return f
}

func initRuleset(lines []string) ruleset {
	r := make(ruleset)
	for _, line := range lines {
		kv := strings.Split(line, " => ")
		r[kv[0]] = new(fractal)
		*r[kv[0]] = deserialize(kv[1])
	}
	return r
}

func makeEmpty(n int) fractal {
	f := make(fractal, n)
	for i := range n {
		f[i] = make([]byte, n)
	}
	return f
}

func (f fractal) mirror() fractal {
	fNew := make(fractal, len(f))

	for i, row := range f {
		fNew[i] = slices.Clone(row)
		slices.Reverse(fNew[i])
	}

	return fNew
}

func (f fractal) rotate() fractal {
	n := len(f)
	fNew := makeEmpty(n)

	for i := range n {
		for j := range n {
			fNew[i][j] = f[j][n-i-1]
		}
	}

	return fNew
}

func (r *ruleset) transform(f fractal) fractal {
	if out, ok := (*r)[f.serialize()]; ok {
		return *out
	}

	f0, f1 := f, f.mirror()
	var fOut *fractal
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

func (r ruleset) getInitialState(seed string) State {
	if len(seed) != 11 {
		panic("Invalid seed; initial matrix must be 3x3.")
	}

	return State{
		r.transform(deserialize(seed)).serialize(): 1,
	}
}

func (f fractal) getSubfractal(i, j, n int) fractal {
	subFractal := makeEmpty(n)
	for k := range n {
		copy(subFractal[k], f[n*i+k][n*j:])
	}
	return subFractal
}

func (f *fractal) setSubfractal(i, j, n int, sf fractal) {
	for k := range n {
		copy((*f)[n*i+k][n*j:], sf[k])
	}
}

func (dp *DynamicProgram) calculateEntry(serial string) {
	f0 := deserialize(serial)
	var f1 fractal
	var aux [3]State

	// 4x4 -> 6x6
	f1 = makeEmpty(6)
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
	f1 = makeEmpty(9)
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
	f1 = makeEmpty(12)
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

func (dp *DynamicProgram) evolve(state0 State, nIters int) State {
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
		return state0
	}

	state1 := make(State)
	for serial0, cnt0 := range state0 {
		for serial1, cnt1 := range dp.getNthState(serial0, nIters%3-1) {
			state1[serial1] += cnt0 * cnt1
		}
	}

	return state1
}

func (s State) count() int {
	var res int
	for serial, cnt := range s {
		res += strings.Count(serial, "#") * cnt
	}
	return res
}

func init() {
	Run = func(seed string, nIters int, lines []string) int {
		if nIters == 0 {
			return State{seed: 1}.count()
		}

		r := initRuleset(lines)
		s := r.getInitialState(seed)
		dp := DynamicProgram{
			memory: make(DPMemory),
			rules:  r,
		}

		return dp.evolve(s, nIters-1).count()
	}
}
