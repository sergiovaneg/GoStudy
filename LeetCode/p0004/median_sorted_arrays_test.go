package p0004_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0004"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var res float64

	res = p0004.FindMedianSortedArrays([]int{1, 3}, []int{2})
	if res != 2. {
		t.Fatalf("Expected 2.0, got %v", res)
	}

	res = p0004.FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	if res != 2.5 {
		t.Fatalf("Expected 2.5, got %v", res)
	}
}
