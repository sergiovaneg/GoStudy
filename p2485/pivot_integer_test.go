package p2485_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p2485"
)

func TestPivotInteger(t *testing.T) {
	if p2485.PivotInteger(8) != 6 {
		t.Fatal()
	}
	if p2485.PivotInteger(1) != 1 {
		t.Fatal()
	}
	if p2485.PivotInteger(4) != -1 {
		t.Fatal()
	}
}
