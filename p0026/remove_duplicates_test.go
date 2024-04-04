package p0026_test

import (
	"reflect"
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0026"
)

func TestRemoveDuplicates(t *testing.T) {
	var nums, expected []int
	var res int

	nums = []int{1, 1, 2}
	expected = []int{1, 2}
	res = p0026.RemoveDuplicates(nums)
	if res != len(expected) {
		t.Fatalf("Expected length %v; got %v", len(expected), res)
	} else if !reflect.DeepEqual(nums[:res], expected) {
		t.Fatalf("Expected vector %v; got %v", expected, nums[:res])
	}

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4}
	expected = []int{0, 1, 2, 3, 4}
	res = p0026.RemoveDuplicates(nums)
	if res != len(expected) {
		t.Fatalf("Expected length %v; got %v", len(expected), res)
	} else if !reflect.DeepEqual(nums[:res], expected) {
		t.Fatalf("Expected vector %v; got %v", expected, nums[:res])
	}

	nums = []int{0}
	expected = []int{0}
	res = p0026.RemoveDuplicates(nums)
	if res != len(expected) {
		t.Fatalf("Expected length %v; got %v", len(expected), res)
	} else if !reflect.DeepEqual(nums[:res], expected) {
		t.Fatalf("Expected vector %v; got %v", expected, nums[:res])
	}
}
