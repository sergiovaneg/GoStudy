package p0011_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0011"
)

func TestMaxArea(t *testing.T) {
	var res int

	res = p0011.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if res != 49 {
		t.Fatalf("Expected 49; got %v", res)
	}

	res = p0011.MaxArea([]int{1, 1})
	if res != 1 {
		t.Fatalf("Expected 1; got %v", res)
	}
}
