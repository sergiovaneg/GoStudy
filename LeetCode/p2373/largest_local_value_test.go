package p2373_test

import (
	"reflect"
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p2373"
)

func TestLargestLocal(t *testing.T) {
	var res, wnt [][]int

	res = p2373.LargestLocal([][]int{
		{9, 9, 8, 1},
		{5, 6, 2, 6},
		{8, 2, 6, 4},
		{6, 2, 2, 2},
	})
	wnt = [][]int{
		{9, 9},
		{8, 6},
	}
	if !reflect.DeepEqual(res, wnt) {
		t.Fatalf("Expected:\n%v\nGot:\n%v", wnt, res)
	}

	res = p2373.LargestLocal([][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	})
	wnt = [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}
	if !reflect.DeepEqual(res, wnt) {
		t.Fatalf("Expected:\n%v\nGot:\n%v", wnt, res)
	}
}
