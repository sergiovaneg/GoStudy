package p1289_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p1289"
)

func TestMinFallingPathSum(t *testing.T) {
	var res int

	res = p1289.MinFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	if res != 13 {
		t.Fatalf("Expected 13; got %v", res)
	}

	res = p1289.MinFallingPathSum([][]int{{7}})
	if res != 7 {
		t.Fatalf("Expected 7; got %v", res)
	}
}
