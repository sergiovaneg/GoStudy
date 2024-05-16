package p0713_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0713"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	var res int

	res = p0713.NumSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)
	if res != 8 {
		t.Fatalf("Expected 8; got %v", res)
	}

	res = p0713.NumSubarrayProductLessThanK([]int{1, 2, 3}, 0)
	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}

	res = p0713.NumSubarrayProductLessThanK([]int{1, 1, 1}, 2)
	if res != 6 {
		t.Fatalf("Expected 6; got %v", res)
	}
}
