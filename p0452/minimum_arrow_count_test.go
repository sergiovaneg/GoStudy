package p0452_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0452"
)

func TestMinArrowShots(t *testing.T) {
	points := [][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}
	if p0452.FindMinArrowShots(points) != 2 {
		t.Fatal()
	}
	if p0452.FindMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}) != 2 {
		t.Fatal()
	}
	if p0452.FindMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}) != 4 {
		t.Fatal()
	}
}
