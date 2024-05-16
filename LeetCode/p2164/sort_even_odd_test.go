package p2164_test

import (
	"reflect"
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p2164"
)

func TestSortEvenOdd(t *testing.T) {
	nums := []int{4, 1, 2, 3}
	nums = p2164.SortEvenOdd(nums)
	if !reflect.DeepEqual(nums, []int{2, 3, 4, 1}) {
		t.Fatal()
	}
}
