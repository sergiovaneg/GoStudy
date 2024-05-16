package p0048_test

import (
	"reflect"
	"testing"

	"github.com/sergiovaneg/GO_leetcode/LeetCode/p0048"
)

func TestRotate(t *testing.T) {
	var image, expected [][]int

	image = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	p0048.Rotate(image)
	expected = [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	if !reflect.DeepEqual(image, expected) {
		t.Fatalf("Expected:\n%v\nGot:\n%v", expected, image)
	}

	image = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	p0048.Rotate(image)
	expected = [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	if !reflect.DeepEqual(image, expected) {
		t.Fatalf("Expected:\n%v\nGot:\n%v", expected, image)
	}
}
