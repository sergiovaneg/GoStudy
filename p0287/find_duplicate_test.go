package p0287_test

import (
	"testing"

	"github.com/sergiovaneg/GO_leetcode/p0287"
)

func TestFindDuplicate(t *testing.T) {
	if p0287.FindDuplicate([]int{1, 3, 4, 2, 2}) != 2 {
		t.Fatal()
	}
	if p0287.FindDuplicate([]int{3, 1, 3, 4, 2}) != 3 {
		t.Fatal()
	}
	if p0287.FindDuplicate([]int{3, 3, 3, 3, 3}) != 3 {
		t.Fatal()
	}
}
