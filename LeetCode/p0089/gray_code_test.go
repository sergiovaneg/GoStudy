package p0089_test

import (
	"math/bits"
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0089"
)

func checkBitDiff(a, b int) int {
	return bits.OnesCount(uint(a ^ b))
}

func TestGrayCode(t *testing.T) {
	res := p0089.GrayCode(16)

	n := len(res)
	found := make([]bool, n)

	found[res[0]] = true
	if checkBitDiff(res[0], res[n-1]) != 1 {
		t.Fatal("Difference is not 1 bit between first and last number")
	}

	for idx, num := range res[1:] {
		if checkBitDiff(num, res[idx]) != 1 {
			t.Fatal("Difference is not 1 bit")
		}
		found[num] = true
	}

	for idx, f := range found {
		if !f {
			t.Fatalf("%v not found", idx)
		}
	}
}
