package p0514_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0514"
)

func TestFindRotateSteps(t *testing.T) {
	var res int

	res = p0514.FindRotateSteps("godding", "gd")
	if res != 4 {
		t.Fatalf("Expected 4; got %v", res)
	}

	res = p0514.FindRotateSteps("godding", "godding")
	if res != 13 {
		t.Fatalf("Expected 13; got %v", res)
	}
}
