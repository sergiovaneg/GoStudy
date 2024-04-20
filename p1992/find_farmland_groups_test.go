package p1992_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p1992"
)

func TestFindFarmland(t *testing.T) {
	var land [][]int

	land = [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}
	t.Log(p1992.FindFarmland(land))

	land = [][]int{{1, 1}, {1, 1}}
	t.Log(p1992.FindFarmland(land))

	land = [][]int{{0}}
	t.Log(p1992.FindFarmland(land))
}
