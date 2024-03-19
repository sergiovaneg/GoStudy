package p0621_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0621"
)

func TestLeastInterval(t *testing.T) {
	if p0621.LeastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, 2) != 12 {
		t.Fatal()
	}
}
