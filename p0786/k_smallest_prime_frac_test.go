package p0786_test

import (
	"slices"
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0786"
)

func TestKSmallestFrac(t *testing.T) {
	k_max := 28
	fracs := make([]float64, k_max)

	for k := 1; k <= k_max; k++ {
		res := p0786.KthSmallestPrimeFraction([]int{1, 2, 3, 5, 7, 11, 13, 17}, k)
		fracs[k-1] = float64(res[0]) / float64(res[1])
	}

	if !slices.IsSorted(fracs) {
		t.Fatal("Algorithm not working")
	}
}
