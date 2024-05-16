package p0506_test

import (
	"reflect"
	"testing"

	"github.com/sergiovaneg/GoStudy/LeetCode/p0506"
)

func TestFindRelativeRanks(t *testing.T) {
	var res, want []string

	res = p0506.FindRelativeRanks([]int{5, 4, 3, 2, 1})
	want = []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}
	if !reflect.DeepEqual(res, want) {
		t.Fatalf("Expected %v; got %v", want, res)
	}

	res = p0506.FindRelativeRanks([]int{10, 3, 8, 9, 4})
	want = []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}
	if !reflect.DeepEqual(res, want) {
		t.Fatalf("Expected %v; got %v", want, res)
	}
}
