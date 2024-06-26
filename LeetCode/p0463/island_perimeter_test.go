package p0463_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0463"
)

func TestIslandPerimeter(t *testing.T) {
	var grid [][]int
	var res int

	IslandPerimeter := p0463.IslandPerimeterParallelIterative

	grid = [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	res = IslandPerimeter(grid)
	if res != 16 {
		t.Fatalf("Expected 16; got %v", res)
	}

	grid = [][]int{{1}}
	res = IslandPerimeter(grid)
	if res != 4 {
		t.Fatalf("Expected 4; got %v", res)
	}

	grid = [][]int{{1, 0}}
	res = IslandPerimeter(grid)
	if res != 4 {
		t.Fatalf("Expected 4; got %v", res)
	}
}
