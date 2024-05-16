package p1971_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p1971"
)

func TestValidPath(t *testing.T) {
	var edges [][]int

	edges = [][]int{{0, 1}, {1, 2}, {2, 0}}
	if !p1971.ValidPath(3, edges, 0, 2) {
		t.Fatal("Valid path exists.")
	}

	edges = [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	if p1971.ValidPath(6, edges, 0, 5) {
		t.Fatal("No valid path exists.")
	}
}
