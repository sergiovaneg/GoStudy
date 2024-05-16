package p1700_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p1700"
)

func TestCountStudents(t *testing.T) {
	var res int

	res = p1700.CountStudents(
		[]int{1, 1, 0, 0},
		[]int{0, 1, 0, 1},
	)
	if res != 0 {
		t.Fatalf("Expected 0; got %v", res)
	}

	res = p1700.CountStudents(
		[]int{1, 1, 1, 0, 0, 1},
		[]int{1, 0, 0, 0, 1, 1},
	)
	if res != 3 {
		t.Fatalf("Expected 3; got %v", res)
	}
}
