package p1219_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p1219"
)

func TestGetMaximumGold(t *testing.T) {
	var res int

	res = p1219.GetMaximumGold([][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	})
	if res != 24 {
		t.Fatalf("Expected 24; got %v", res)
	}

	res = p1219.GetMaximumGold([][]int{
		{1, 0, 7},
		{2, 0, 6},
		{3, 4, 5},
		{0, 3, 0},
		{9, 0, 20},
	})
	if res != 28 {
		t.Fatalf("Expected 28; got %v", res)
	}
}
