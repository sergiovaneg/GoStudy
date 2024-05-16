package p0041_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0041"
)

func TestFirstMissingPositive(t *testing.T) {
	var res int

	res = p0041.FirstMissingPositive([]int{1, 2, 0})
	if res != 3 {
		t.Fatalf("Expected 3; got %v", res)
	}

	res = p0041.FirstMissingPositive([]int{3, 4, -1, 1})
	if res != 2 {
		t.Fatalf("Expected 2; got %v", res)
	}

	res = p0041.FirstMissingPositive([]int{7, 8, 9, 11, 12})
	if res != 1 {
		t.Fatalf("Expected 1; got %v", res)
	}
}
