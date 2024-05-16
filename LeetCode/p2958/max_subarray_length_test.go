package p2958_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p2958"
)

func TestMaxSubarrayLength(t *testing.T) {
	var res int

	res = p2958.MaxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2)
	if res != 6 {
		t.Fatalf("Expected 6; got %v", res)
	}

	res = p2958.MaxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1)
	if res != 2 {
		t.Fatalf("Expected 2; got %v", res)
	}

	res = p2958.MaxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4)
	if res != 4 {
		t.Fatalf("Expected 4; got %v", res)
	}
}
