package p0442_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p0442"
)

func TestFindDuplicates(t *testing.T) {
	res := p0442.FindDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})
	slices.Sort(res)
	if !reflect.DeepEqual(res, []int{2, 3}) {
		t.Fatal()
	}
}
