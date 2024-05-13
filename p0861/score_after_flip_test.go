package p0861_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0861"
)

func TestMatrixScore(t *testing.T) {
	var res int

	res = p0861.MatrixScore([][]int{
		{0, 0, 1, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 0},
	})
	if res != 39 {
		t.Fatalf("Expected 39; got %v", res)
	}

	res = p0861.MatrixScore([][]int{
		{0},
	})
	if res != 1 {
		t.Fatalf("Expected 1; got %v", res)
	}
}
