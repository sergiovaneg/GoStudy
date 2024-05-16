package p0042_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p0042"
)

func TestTrap(t *testing.T) {
	if p0042.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}) != 6 {
		t.Fatal()
	}
	if p0042.Trap([]int{4, 2, 0, 3, 2, 5}) != 9 {
		t.Fatal()
	}
}
