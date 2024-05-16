package p0881_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0881"
)

func TestNumRescueBoats(t *testing.T) {
	var res int

	res = p0881.NumRescueBoats([]int{1, 2}, 3)
	if res != 1 {
		t.Fatalf("Expected 1; got %v", res)
	}

	res = p0881.NumRescueBoats([]int{3, 2, 2, 1}, 3)
	if res != 3 {
		t.Fatalf("Expected 3; got %v", res)
	}

	res = p0881.NumRescueBoats([]int{3, 5, 3, 4}, 5)
	if res != 4 {
		t.Fatalf("Expected 4; got %v", res)
	}
}
