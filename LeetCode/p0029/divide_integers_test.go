package p0029_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0029"
)

func TestDivide(t *testing.T) {
	if p0029.Divide(-2147483648, -1) != 2147483648 {
		t.Fatal()
	}

	if p0029.Divide(2147483647, 2) != 1073741823 {
		t.Fatal()
	}
}
