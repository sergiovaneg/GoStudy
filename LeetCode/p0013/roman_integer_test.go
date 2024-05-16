package p0013_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p0013"
)

func TestRomanToInt(t *testing.T) {
	if p0013.RomanToInt("MCMXCIV") != 1994 {
		t.Fatal()
	}
}
