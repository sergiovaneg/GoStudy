package p2997_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p2997"
)

func TestMinOperations(t *testing.T) {
	var res int

	res = p2997.MinOperations([]int{2, 1, 3, 4}, 1)
	if res != 2 {
		t.Fatalf("Expected 2; got %v", res)
	}

	res = p2997.MinOperations([]int{2, 0, 2, 0}, 0)

	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}

	res = p2997.MinOperations([]int{0}, 1023)
	if res != 10 {
		t.Fatalf("Expected 10; got %v", res)
	}
}
