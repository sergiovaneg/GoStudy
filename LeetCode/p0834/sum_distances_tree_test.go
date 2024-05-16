package p0834_test

import (
	"reflect"
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p0834"
)

func TestSumOfDistancesInTree(t *testing.T) {
	var res []int

	res = p0834.SumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}})
	if !reflect.DeepEqual(res, []int{8, 12, 6, 10, 10, 10}) {
		t.Fatal("Wrong answer")
	}

	res = p0834.SumOfDistancesInTree(1, [][]int{{}})
	if !reflect.DeepEqual(res, []int{0}) {
		t.Fatal("Wrong answer")
	}

	res = p0834.SumOfDistancesInTree(2, [][]int{{1, 0}})
	if !reflect.DeepEqual(res, []int{1, 1}) {
		t.Fatal("Wrong answer")
	}
}
