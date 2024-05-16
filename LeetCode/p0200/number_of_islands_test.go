package p0200_test

import (
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0200"
)

func TestNumIslands(t *testing.T) {
	var res int

	res = p0200.NumIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}})
	if res != 1 {
		t.Fatalf("Expected 1; got %v", res)
	}

	res = p0200.NumIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}})
	if res != 3 {
		t.Fatalf("Expected 3; got %v", res)
	}

	res = p0200.NumIslands([][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}})
	if res != 1 {
		t.Fatalf("Expected 1; got %v", res)
	}
}
