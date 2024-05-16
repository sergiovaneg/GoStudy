package p2241_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p2241"
)

func TestFindMaxK(t *testing.T) {
	var res int

	res = p2241.FindMaxK([]int{-1, 2, -3, 3})
	if res != 3 {
		t.Fatalf("Expected 3; got %v", res)
	}

	res = p2241.FindMaxK([]int{-1, 10, 6, 7, -7, 1})
	if res != 7 {
		t.Fatalf("Expected 7; got %v", res)
	}

	res = p2241.FindMaxK([]int{-10, 8, 6, 7, -2, -3})
	if res != -1 {
		t.Fatalf("Expected -1; got %v", res)
	}
}
