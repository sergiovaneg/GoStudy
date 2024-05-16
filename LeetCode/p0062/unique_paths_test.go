package p0062_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0062"
)

func TestUniquePaths(t *testing.T) {
	var res int

	res = p0062.UniquePaths(3, 7)
	if res != 28 {
		t.Fatalf("Expected 28; got %v", res)
	}

	res = p0062.UniquePaths(3, 2)
	if res != 3 {
		t.Fatalf("Expected 3; got %v", res)
	}
}
