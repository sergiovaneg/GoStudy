package main

import (
	"fmt"
	"testing"

	"github.com/sergiovaneg/GoStudy/AoC/2017/fractalArt/fractals"
)

const maxBenchmark = 22

func benchSingle(b *testing.B, s fractals.Solver, nIters int, rules []string) {
	for b.Loop() {
		s.Solve(SEED, nIters, rules)
	}
}

func benchLoop(b *testing.B, s fractals.Solver, rules []string) {
	for i := range maxBenchmark + 1 {
		b.Run(fmt.Sprint(i), func(b *testing.B) {
			benchSingle(b, s, i, rules)
		})
	}
}

func BenchmarkSolvers(b *testing.B) {
	solvers := []fractals.Solver{
		fractals.NaiveSolver{},
		// fractals.NaiveConcurrentSolver{},
		// fractals.NaiveDPSolver{},
		// fractals.GroupedSolver{},
	}

	rules := getRules()

	for _, s := range solvers {
		b.Run(s.String(), func(b *testing.B) {
			benchLoop(b, s, rules)
		})
	}
}
