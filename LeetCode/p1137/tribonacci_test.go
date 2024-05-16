package p1137_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p1137"
)

func TestTribonacci(t *testing.T) {
	var res int

	res = p1137.Tribonacci(4)
	if res != 4 {
		t.Fatalf("Expected 4; got %v.", res)
	}

	res = p1137.Tribonacci(25)
	if res != 1389537 {
		t.Fatalf("Expected 1389537; got %v.", res)
	}
}
