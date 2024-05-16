package p2091_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p2091"
)

func TestSortEvenOdd(t *testing.T) {
	var wants int = 5
	result := p2091.MinimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6})
	if result != wants {
		t.Fatalf("Received %v; expected %v", result, wants)
	}
}
