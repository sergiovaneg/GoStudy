package p0950_test

import (
	"reflect"
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0950"
)

func TestDeckRevealedIncreasing(t *testing.T) {
	var res, target []int

	target = []int{2, 13, 3, 11, 5, 17, 7}
	res = p0950.DeckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7})
	if !reflect.DeepEqual(res, target) {
		t.Fatalf("Expected %v; got %v", target, res)
	}

	target = []int{1, 1000}
	res = p0950.DeckRevealedIncreasing([]int{1, 1000})
	if !reflect.DeepEqual(res, target) {
		t.Fatalf("Expected %v; got %v", target, res)
	}
}
