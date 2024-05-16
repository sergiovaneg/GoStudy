package p0930_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0930"
)

func TestTrap(t *testing.T) {
	if p0930.NumSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2) != 4 {
		t.Fatal()
	}
	if p0930.NumSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0) != 15 {
		t.Fatal()
	}
}
